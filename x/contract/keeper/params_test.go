package keeper_test

import (
	"testing"

	testkeeper "github.com/arran8901/chainlog-platform/testutil/keeper"
	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ContractKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
