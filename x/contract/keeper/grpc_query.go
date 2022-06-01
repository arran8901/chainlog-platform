package keeper

import (
	"github.com/arran8901/chainlog-platform/x/contract/types"
)

var _ types.QueryServer = Keeper{}
