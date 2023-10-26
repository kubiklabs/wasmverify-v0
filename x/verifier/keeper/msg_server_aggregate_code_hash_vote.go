package keeper

import (
	"context"
	"encoding/hex"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AggregateCodeHashVote(goCtx context.Context, msg *types.MsgAggregateCodeHashVote) (*types.MsgAggregateCodeHashVoteResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}

	if err := k.ValidateValidator(ctx, valAddr); err != nil {
		return nil, err
	}

	// check vote time for the validator
	contractInfo, found := k.GetContractInfo(ctx, msg.ApplicationId)

	if !found {
		return nil, types.ErrContractInfoNotFound
	}
	finalPrevoteTime, err := k.GetContractPrevoteTime(ctx, msg.ApplicationId)
	if err != nil {
		return nil, err
	}

	if (uint64(ctx.BlockHeight()) > contractInfo.AssignedVerificationBlockHeight) || (uint64(ctx.BlockHeight()) < finalPrevoteTime) {
		return nil, types.ErrVoteTiming
	}

	aggregatePrevote, err := k.GetAggregateCodeHashPrevote(ctx, valAddr)
	if err != nil {
		return nil, types.ErrNoAggregatePrevote.Wrap((msg.Operator + " for appId: " + string(msg.ApplicationId)))
	}

	// Verify that the vote hash and prevote hash match
	hash := types.GetAggregateVoteHash(msg.Salt, msg.CodeHash, msg.Operator)
	if aggregatePrevote.Hash != hex.EncodeToString(hash) {
		return nil, types.ErrVerificationFailed.Wrapf("must be given %s not %s", aggregatePrevote.Hash, hash)
	}

	// Move aggregate prevote to aggregate vote with given CodeHash
	aggregateCodeHashVote := types.CodeHashVote{
		ApplicationId: msg.ApplicationId,
		CodeHash:      msg.CodeHash,
		Voter:         valAddr.String(),
	}
	k.SetAggregateCodeHashVote(ctx, valAddr, aggregateCodeHashVote)
	k.DeleteAggregateCodeHashPrevote(ctx, valAddr)

	return &types.MsgAggregateCodeHashVoteResponse{}, nil
}
