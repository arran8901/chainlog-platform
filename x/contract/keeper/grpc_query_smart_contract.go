package keeper

import (
	"context"

	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SmartContractAll(c context.Context, req *types.QueryAllSmartContractRequest) (*types.QueryAllSmartContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var smartContracts []types.SmartContract
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	smartContractStore := prefix.NewStore(store, types.KeyPrefix(types.SmartContractKeyPrefix))

	pageRes, err := query.Paginate(smartContractStore, req.Pagination, func(key []byte, value []byte) error {
		var smartContract types.SmartContract
		if err := k.cdc.Unmarshal(value, &smartContract); err != nil {
			return err
		}

		smartContracts = append(smartContracts, smartContract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSmartContractResponse{SmartContract: smartContracts, Pagination: pageRes}, nil
}

func (k Keeper) SmartContract(c context.Context, req *types.QueryGetSmartContractRequest) (*types.QueryGetSmartContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSmartContract(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSmartContractResponse{SmartContract: val}, nil
}
