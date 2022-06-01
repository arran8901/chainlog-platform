package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arran8901/chainlog-platform/testutil/keeper"
	"github.com/arran8901/chainlog-platform/x/contract/keeper"
	"github.com/arran8901/chainlog-platform/x/contract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ContractKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
