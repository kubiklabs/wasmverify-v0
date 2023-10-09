package simulation

import (
	"math/rand"

	"verifier/x/verifier/keeper"
	"verifier/x/verifier/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAggregateCodeHashVote(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAggregateCodeHashVote{
			Operator: simAccount.Address.String(),
		}

		// TODO: Handling the AggregateCodeHashVote simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AggregateCodeHashVote simulation not implemented"), nil, nil
	}
}
