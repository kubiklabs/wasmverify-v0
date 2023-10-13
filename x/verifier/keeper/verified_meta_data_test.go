package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "verifier/testutil/keeper"
	"verifier/testutil/nullify"
	"verifier/x/verifier/keeper"
	"verifier/x/verifier/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVerifiedMetaData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VerifiedMetaData {
	items := make([]types.VerifiedMetaData, n)
	for i := range items {
		items[i].CodeId = uint64(i)

		keeper.SetVerifiedMetaData(ctx, items[i])
	}
	return items
}

func TestVerifiedMetaDataGet(t *testing.T) {
	keeper, ctx := keepertest.VerifierKeeper(t)
	items := createNVerifiedMetaData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVerifiedMetaData(ctx,
			item.CodeId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVerifiedMetaDataRemove(t *testing.T) {
	keeper, ctx := keepertest.VerifierKeeper(t)
	items := createNVerifiedMetaData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVerifiedMetaData(ctx,
			item.CodeId,
		)
		_, found := keeper.GetVerifiedMetaData(ctx,
			item.CodeId,
		)
		require.False(t, found)
	}
}

func TestVerifiedMetaDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.VerifierKeeper(t)
	items := createNVerifiedMetaData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVerifiedMetaData(ctx)),
	)
}
