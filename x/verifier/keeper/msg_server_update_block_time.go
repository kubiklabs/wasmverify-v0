package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Anyone can change the blocktime
// Add check for it
func (k msgServer) UpdateBlockTime(goCtx context.Context, msg *types.MsgUpdateBlockTime) (*types.MsgUpdateBlockTimeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.SetTotalVerificationBlocks(ctx, msg.CompilationBlocks, msg.PrevoteBlocks, msg.VoteBlocks)

	return &types.MsgUpdateBlockTimeResponse{}, nil
}
