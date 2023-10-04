package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"verifier/x/verifier/types"
)

func (k msgServer) ApplyVerifyApplication(goCtx context.Context, msg *types.MsgApplyVerifyApplication) (*types.MsgApplyVerifyApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgApplyVerifyApplicationResponse{}, nil
}
