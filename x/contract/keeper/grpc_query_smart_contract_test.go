package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/arran8901/chainlog-platform/testutil/keeper"
	"github.com/arran8901/chainlog-platform/testutil/nullify"
	"github.com/arran8901/chainlog-platform/x/contract/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestSmartContractQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSmartContract(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSmartContractRequest
		response *types.QueryGetSmartContractResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSmartContractRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetSmartContractResponse{SmartContract: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSmartContractRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetSmartContractResponse{SmartContract: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSmartContractRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SmartContract(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestSmartContractQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSmartContract(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSmartContractRequest {
		return &types.QueryAllSmartContractRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SmartContractAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SmartContract), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SmartContract),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SmartContractAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SmartContract), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SmartContract),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SmartContractAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SmartContract),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SmartContractAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
