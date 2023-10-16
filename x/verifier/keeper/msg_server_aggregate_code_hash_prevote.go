package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AggregateCodeHashPrevote(goCtx context.Context, msg *types.MsgAggregateCodeHashPrevote) (*types.MsgAggregateCodeHashPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}

	// Verify the voter is validator
	if err := k.ValidateValidator(ctx, valAddr); err != nil {
		return nil, err
	}
	// Check Already voted..?
	if k.HasAggregateCodeHashPrevote(ctx, valAddr) {
		return nil, types.ErrExistingPrevote
	}

	// Check if the prevote made by validator is in the vote period for the penidg Contarct
	finalPrevoteTime, err := k.GetContractPrevoteTime(ctx, msg.ApplicationId)

	if err != nil {
		return nil, err
	}

	if uint64(ctx.BlockHeight()) > finalPrevoteTime {
		return nil, types.ErrPreVoteTimePassed
	}

	aggregatePrevote := types.CodeHashPreVote{
		ApplicationId: msg.ApplicationId,
		Hash:          msg.Hash,
		Voter:         msg.Validator,
		SubmitBlock:   uint64(ctx.BlockHeight()),
	}
	// (ctxbyte val address prevote struct)
	k.SetAggregateCodeHashPrevote(ctx, valAddr, aggregatePrevote)

	return &types.MsgAggregateCodeHashPrevoteResponse{}, nil
}
