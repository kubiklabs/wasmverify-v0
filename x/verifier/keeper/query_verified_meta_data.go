package keeper

import (
	"context"

	"verifier/x/verifier/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifiedMetaDataAll(goCtx context.Context, req *types.QueryAllVerifiedMetaDataRequest) (*types.QueryAllVerifiedMetaDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var verifiedMetaDatas []types.VerifiedMetaData
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	verifiedMetaDataStore := prefix.NewStore(store, types.KeyPrefix(types.VerifiedMetaDataKeyPrefix))

	pageRes, err := query.Paginate(verifiedMetaDataStore, req.Pagination, func(key []byte, value []byte) error {
		var verifiedMetaData types.VerifiedMetaData
		if err := k.cdc.Unmarshal(value, &verifiedMetaData); err != nil {
			return err
		}

		verifiedMetaDatas = append(verifiedMetaDatas, verifiedMetaData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVerifiedMetaDataResponse{VerifiedMetaData: verifiedMetaDatas, Pagination: pageRes}, nil
}

func (k Keeper) VerifiedMetaData(goCtx context.Context, req *types.QueryGetVerifiedMetaDataRequest) (*types.QueryGetVerifiedMetaDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetVerifiedMetaData(
		ctx,
		req.CodeId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVerifiedMetaDataResponse{VerifiedMetaData: val}, nil
}
