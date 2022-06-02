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
	value, _ := sdk.ParseCoinsNormalized(msg.Value)

	// Check sender has enough funds to pay value
	if balances := k.bankKeeper.GetAllBalances(ctx, creator); value.IsAllGT(balances) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "account %s has insufficient funds to pay value %s", msg.Creator, msg.Value)
	}

	// Generate contract address
	creatorSequence, _ := k.accountKeeper.GetSequence(ctx, creator)
	contractAddress := types.GenerateContractAddress(creator, creatorSequence)

	return &types.MsgCreateContractResponse{ContractAddress: contractAddress.String()}, nil
}
