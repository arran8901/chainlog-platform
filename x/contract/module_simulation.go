package contract

import (
	"math/rand"

	"github.com/arran8901/chainlog-platform/testutil/sample"
	contractsimulation "github.com/arran8901/chainlog-platform/x/contract/simulation"
	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = contractsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateContract = "op_weight_msg_create_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContract int = 100

	opWeightMsgCallContract = "op_weight_msg_call_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCallContract int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	contractGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&contractGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateContract, &weightMsgCreateContract, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContract = defaultWeightMsgCreateContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContract,
		contractsimulation.SimulateMsgCreateContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCallContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCallContract, &weightMsgCallContract, nil,
		func(_ *rand.Rand) {
			weightMsgCallContract = defaultWeightMsgCallContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCallContract,
		contractsimulation.SimulateMsgCallContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
