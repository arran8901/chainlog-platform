package contract_test

import (
	"testing"

	keepertest "github.com/arran8901/chainlog-platform/testutil/keeper"
	"github.com/arran8901/chainlog-platform/testutil/nullify"
	"github.com/arran8901/chainlog-platform/x/contract"
	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SmartContractList: []types.SmartContract{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ContractKeeper(t)
	contract.InitGenesis(ctx, *k, genesisState)
	got := contract.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SmartContractList, got.SmartContractList)
	// this line is used by starport scaffolding # genesis/test/assert
}
