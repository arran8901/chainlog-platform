package keeper

import (
	"context"

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

	_ = contract

	return &types.MsgCallContractResponse{}, nil
}
