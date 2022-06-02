package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/arran8901/chainlog-platform/testutil/keeper"
	"github.com/arran8901/chainlog-platform/testutil/nullify"
	"github.com/arran8901/chainlog-platform/x/contract/keeper"
	"github.com/arran8901/chainlog-platform/x/contract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSmartContract(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SmartContract {
	items := make([]types.SmartContract, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetSmartContract(ctx, items[i])
	}
	return items
}

func TestSmartContractGet(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSmartContract(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSmartContract(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSmartContractRemove(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSmartContract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSmartContract(ctx,
			item.Address,
		)
		_, found := keeper.GetSmartContract(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestSmartContractGetAll(t *testing.T) {
	keeper, ctx := keepertest.ContractKeeper(t)
	items := createNSmartContract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSmartContract(ctx)),
	)
}
