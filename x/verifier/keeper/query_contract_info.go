package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContractInfo(goCtx context.Context, req *types.QueryContractInfoRequest) (*types.QueryContractInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	contractinfo, found := k.GetContractInfo(ctx, req.Id)

	if !found {
		return nil, status.Error(codes.NotFound, "Contract Not Found")
	}

	return &types.QueryContractInfoResponse{
		Contractinfo: &contractinfo,
	}, nil
}
