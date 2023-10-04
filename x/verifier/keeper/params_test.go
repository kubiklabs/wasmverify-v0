package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "verifier/testutil/keeper"
	"verifier/x/verifier/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VerifierKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
