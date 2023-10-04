package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"verifier/x/verifier/types"
)

func (k msgServer) AggregateCodeHashPrevote(goCtx context.Context, msg *types.MsgAggregateCodeHashPrevote) (*types.MsgAggregateCodeHashPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAggregateCodeHashPrevoteResponse{}, nil
}
