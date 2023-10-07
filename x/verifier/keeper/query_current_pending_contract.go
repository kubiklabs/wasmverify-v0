package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CurrentPendingContract(goCtx context.Context, req *types.QueryCurrentPendingContractRequest) (*types.QueryCurrentPendingContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	currentPendingContractId := k.GetCurrentPendingContractId(ctx)

	if currentPendingContractId == 0 {
		return nil, types.ErrNoPendingContractfound
	}

	return &types.QueryCurrentPendingContractResponse{
		Id: currentPendingContractId,
	}, nil
}
