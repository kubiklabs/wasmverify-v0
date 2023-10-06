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

	// (total contract - total pending contract) = id of next pending contract
	totalContract := k.GetContractCount(ctx)
	totalPendingContract := k.GetPendingContractsCount(ctx)

	var currentPendingContractId uint64 = 0
	if totalPendingContract == 0 {
		currentPendingContractId = 0
	}

	currentPendingContractId = ((totalContract) - totalPendingContract + 1)

	_, found := k.GetContractInfo(ctx, currentPendingContractId)
	if !found {
		return &types.QueryCurrentPendingContractResponse{
			Id: 0,
		}, nil
	}

	return &types.QueryCurrentPendingContractResponse{
		Id: currentPendingContractId,
	}, nil
}
