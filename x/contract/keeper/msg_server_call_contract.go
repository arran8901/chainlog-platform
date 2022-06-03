package keeper

import (
	"context"

	chainlog "github.com/arran8901/chainlog-lang"
	"github.com/arran8901/chainlog-platform/x/contract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CallContract(goCtx context.Context, msg *types.MsgCallContract) (*types.MsgCallContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse creator, contract address and value
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	contractAddress, _ := sdk.AccAddressFromBech32(msg.ContractAddress)
	value, _ := sdk.ParseCoinNormalized(msg.Value)

	// Check sender has enough funds to pay value
	if !k.bankKeeper.HasBalance(ctx, creator, value) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "account %s has insufficient funds to pay value %s", msg.Creator, msg.Value)
	}

	// Attempt to get contract from store
	// Return an error if not found
	contract, found := k.GetSmartContract(ctx, contractAddress)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "no contract found at address %s", msg.ContractAddress)
	}

	// Transfer value from sender to contract
	err := k.bankKeeper.SendCoins(ctx, creator, contractAddress, sdk.NewCoins(value))
	if err != nil {
		return nil, err
	}

	// Get contract account's balance
	contractBalance := k.bankKeeper.GetBalance(ctx, contractAddress, types.SmartContractCoinDenom)

	// Initialise new Chainlog interpreter and load code and dynamic KB
	i := chainlog.NewInterpreter()
	i.ConsultWithDynamicKB(contract.Code, chainlog.ParseDynamicKBLogicProgram(contract.DynamicKb))

	// Construct message context
	msgCtx := chainlog.MessageContext{
		Sender:  chainlog.Address(msg.Creator),
		Value:   value.Amount.Uint64(),
		Time:    ctx.BlockTime().Unix(),
		Balance: contractBalance.Amount.Uint64(),
	}

	// Submit message to interpreter within this context and get resultant actions
	// If an exception occurs, return an error
	actions, err := i.Message(msg.MessageTerm, &msgCtx)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrContractMessageException, err.Error())
	}

	// Perform all actions
	dynamicKbUpdated := false
	for _, action := range actions {
		switch v := action.(type) {
		case chainlog.AssertAction:
			i.Assert(v.Term)
			dynamicKbUpdated = true
		case chainlog.RetractAction:
			i.Retract(v.Term)
			dynamicKbUpdated = true
		case chainlog.TransferAction:
			// Check to address is a valid address
			toAddress, err := sdk.AccAddressFromBech32(v.ToAddress)
			if err != nil {
				return nil, sdkerrors.Wrapf(types.ErrContractMessageException, "invalid to address: %s in transfer action %s", v.ToAddress, v.String())
			}
			// Check the contract has the required balance
			coin := sdk.NewCoin(types.SmartContractCoinDenom, sdk.NewIntFromUint64(v.Value))
			if !k.bankKeeper.HasBalance(ctx, contractAddress, coin) {
				return nil, sdkerrors.Wrapf(types.ErrContractMessageException, "contract has insufficient balance to transfer value of %d", v.Value)
			}
			// Perform the coin transfer
			k.bankKeeper.SendCoins(ctx, contractAddress, toAddress, sdk.NewCoins(coin))
		}
	}
	if dynamicKbUpdated {
		contract.DynamicKb = chainlog.DynamicKBAsLogicProgram(i.GetDynamicKB())
		k.SetSmartContract(ctx, contractAddress, contract)
	}

	return &types.MsgCallContractResponse{}, nil
}
