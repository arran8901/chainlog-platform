package keeper

import (
	"context"

	"github.com/arran8901/chainlog-platform/x/contract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllContracts(goCtx context.Context, req *types.QueryAllContractsRequest) (*types.QueryAllContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var contracts []types.SmartContractWithAddress

	contractStore := smartContractStore(k, ctx)

	// Paginate contracts based on req.Pagination
	pageRes, err := query.Paginate(contractStore, req.Pagination, func(key []byte, value []byte) error {
		var contract types.SmartContract
		if err := k.cdc.Unmarshal(value, &contract); err != nil {
			return err
		}
		contractAddress := sdk.AccAddress(key[len(types.SmartContractKeyPrefix):])
		contracts = append(contracts, types.SmartContractWithAddress{
			ContractAddress: contractAddress.String(),
			Code:            contract.Code,
			DynamicKb:       contract.DynamicKb,
		})
		return nil
	})

	// Internal error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractsResponse{Contracts: contracts, Pagination: pageRes}, nil
}
