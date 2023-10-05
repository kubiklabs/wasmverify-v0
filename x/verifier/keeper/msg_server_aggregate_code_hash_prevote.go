package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AggregateCodeHashPrevote(goCtx context.Context, msg *types.MsgAggregateCodeHashPrevote) (*types.MsgAggregateCodeHashPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if err := k.ValidateValidator(ctx, valAddr); err != nil {
		return nil, err
	}
	// Till here sennder is validator
	// now save the prevote sent by him

	// Ensure prevote wasn't already submitted
	if ms.HasAggregateExchangeRatePrevote(ctx, valAddr) {
		return nil, types.ErrExistingPrevote
	}

	// Convert hex string to votehash
	voteHash, err := types.AggregateVoteHashFromHex(msg.Hash)
	if err != nil {
		return nil, types.ErrInvalidHash.Wrap(err.Error())
	}

	aggregatePrevote := types.NewAggregateExchangeRatePrevote(voteHash, valAddr, uint64(ctx.BlockHeight()))
	ms.SetAggregateExchangeRatePrevote(ctx, valAddr, aggregatePrevote)

	return &types.MsgAggregateCodeHashPrevoteResponse{}, nil
}
