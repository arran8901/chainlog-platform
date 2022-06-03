package simulation

import (
	"math/rand"

	"github.com/arran8901/chainlog-platform/x/contract/keeper"
	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCallContract(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCallContract{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CallContract simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CallContract simulation not implemented"), nil, nil
	}
}
