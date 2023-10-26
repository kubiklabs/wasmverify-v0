package verifier_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "verifier/testutil/keeper"
	"verifier/testutil/nullify"
	"verifier/x/verifier"
	"verifier/x/verifier/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		VerifiedMetaDataList: []types.VerifiedMetaData{
			{
				CodeId: 0,
			},
			{
				CodeId: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VerifierKeeper(t)
	verifier.InitGenesis(ctx, *k, genesisState)
	got := verifier.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VerifiedMetaDataList, got.VerifiedMetaDataList)
	// this line is used by starport scaffolding # genesis/test/assert
}
