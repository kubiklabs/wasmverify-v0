package verifier

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"verifier/testutil/sample"
	verifiersimulation "verifier/x/verifier/simulation"
	"verifier/x/verifier/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = verifiersimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgAggregateCodeHashPrevote = "op_weight_msg_aggregate_code_hash_prevote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAggregateCodeHashPrevote int = 100

	opWeightMsgApplyVerifyApplication = "op_weight_msg_apply_verify_application"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApplyVerifyApplication int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	verifierGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&verifierGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAggregateCodeHashPrevote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAggregateCodeHashPrevote, &weightMsgAggregateCodeHashPrevote, nil,
		func(_ *rand.Rand) {
			weightMsgAggregateCodeHashPrevote = defaultWeightMsgAggregateCodeHashPrevote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAggregateCodeHashPrevote,
		verifiersimulation.SimulateMsgAggregateCodeHashPrevote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApplyVerifyApplication int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApplyVerifyApplication, &weightMsgApplyVerifyApplication, nil,
		func(_ *rand.Rand) {
			weightMsgApplyVerifyApplication = defaultWeightMsgApplyVerifyApplication
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApplyVerifyApplication,
		verifiersimulation.SimulateMsgApplyVerifyApplication(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAggregateCodeHashPrevote,
			defaultWeightMsgAggregateCodeHashPrevote,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				verifiersimulation.SimulateMsgAggregateCodeHashPrevote(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApplyVerifyApplication,
			defaultWeightMsgApplyVerifyApplication,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				verifiersimulation.SimulateMsgApplyVerifyApplication(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
