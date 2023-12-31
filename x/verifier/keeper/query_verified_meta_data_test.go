package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "verifier/testutil/keeper"
	"verifier/testutil/nullify"
	"verifier/x/verifier/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestVerifiedMetaDataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VerifierKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVerifiedMetaData(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetVerifiedMetaDataRequest
		response *types.QueryGetVerifiedMetaDataResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVerifiedMetaDataRequest{
				CodeId: msgs[0].CodeId,
			},
			response: &types.QueryGetVerifiedMetaDataResponse{VerifiedMetaData: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVerifiedMetaDataRequest{
				CodeId: msgs[1].CodeId,
			},
			response: &types.QueryGetVerifiedMetaDataResponse{VerifiedMetaData: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVerifiedMetaDataRequest{
				CodeId: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VerifiedMetaData(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestVerifiedMetaDataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VerifierKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVerifiedMetaData(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVerifiedMetaDataRequest {
		return &types.QueryAllVerifiedMetaDataRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VerifiedMetaDataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VerifiedMetaData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VerifiedMetaData),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VerifiedMetaDataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VerifiedMetaData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VerifiedMetaData),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VerifiedMetaDataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.VerifiedMetaData),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VerifiedMetaDataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
