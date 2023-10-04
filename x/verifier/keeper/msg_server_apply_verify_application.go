package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApplyVerifyApplication(goCtx context.Context, msg *types.MsgApplyVerifyApplication) (*types.MsgApplyVerifyApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validationId := k.handleResponseBlock(ctx, msg)

	return &types.MsgApplyVerifyApplicationResponse{
		ApplicationId: validationId,
	}, nil
}
