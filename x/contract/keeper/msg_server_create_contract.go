package keeper

import (
	"context"

	"github.com/arran8901/chainlog-platform/x/contract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateContract(goCtx context.Context, msg *types.MsgCreateContract) (*types.MsgCreateContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse creator and value
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	value, _ := sdk.ParseCoinNormalized(msg.Value)

	// Check sender has enough funds to pay value
	if !k.bankKeeper.HasBalance(ctx, creator, value) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "account %s has insufficient funds to pay value %s", msg.Creator, msg.Value)
	}

	// Generate contract address
	creatorSequence, _ := k.accountKeeper.GetSequence(ctx, creator)
	contractAddress := types.GenerateContractAddress(creator, creatorSequence)

	// Create new account for contract and set in account store
	contractAccount := k.accountKeeper.NewAccountWithAddress(ctx, contractAddress)
	k.accountKeeper.SetAccount(ctx, contractAccount)

	// Transfer value from creator to contract
	err := k.bankKeeper.SendCoins(ctx, creator, contractAddress, sdk.NewCoins(value))
	if err != nil {
		return nil, err
	}

	// Create contract struct with code and initially empty dynamic KB
	contract := types.SmartContract{
		Code:      msg.Code,
		DynamicKb: "",
	}
	// Set contract in contract store
	k.SetSmartContract(ctx, contractAddress, contract)

	// Emit event with contract address
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute("address", contractAddress.String()),
		),
	)

	// Return contract address
	return &types.MsgCreateContractResponse{ContractAddress: contractAddress.String()}, nil
}
