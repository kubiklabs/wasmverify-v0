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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VerifierKeeper(t)
	verifier.InitGenesis(ctx, *k, genesisState)
	got := verifier.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
