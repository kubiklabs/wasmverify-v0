package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FinalVerification(goCtx context.Context, msg *types.MsgFinalVerification) (*types.MsgFinalVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	status, err := k.verifyHash(ctx, msg)

	if err != nil {
		return nil, err
	}

	return &types.MsgFinalVerificationResponse{
		Response: status,
	}, nil
}
