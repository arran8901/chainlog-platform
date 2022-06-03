package keeper

import (
	"context"

	"github.com/arran8901/chainlog-platform/x/contract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContractCode(goCtx context.Context, req *types.QueryContractCodeRequest) (*types.QueryContractCodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse address
	address, err := sdk.AccAddressFromBech32(req.ContractAddress)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address: %s (%s)", req.ContractAddress, err.Error())
	}

	// Attempt to get contract from store
	// Return an error if not found
	contract, found := k.GetSmartContract(ctx, address)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "no contract found at address %s", req.ContractAddress)
	}

	// Return contract code and dynamic KB
	return &types.QueryContractCodeResponse{Code: contract.Code, DynamicKb: contract.DynamicKb}, nil
}
