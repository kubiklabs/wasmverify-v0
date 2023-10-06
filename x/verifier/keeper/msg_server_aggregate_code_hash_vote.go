package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AggregateCodeHashVote(goCtx context.Context, msg *types.MsgAggregateCodeHashVote) (*types.MsgAggregateCodeHashVoteResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// derive validator address from tx sender address
	valAddr, err := sdk.ValAddressFromBech32(msg.Creator)
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

	if uint64(ctx.BlockHeight()) > contractInfo.AssignedVerificationBlockHeight {
		return nil, types.ErrVoteTimePassed
	}

	aggregatePrevote, err := k.GetAggregateCodeHashPrevote(ctx, valAddr)
	if err != nil {
		return nil, types.ErrNoAggregatePrevote.Wrap((msg.Creator + " for appId: " + string(msg.ApplicationId)))
	}

	// exchangeRateTuples, err := types.ParseExchangeRateTuples(msg.ExchangeRates)
	// if err != nil {
	// 	return nil, err
	// }

	// Verify that the vote hash and prevote hash match
	hash := types.GetAggregateVoteHash(msg.Salt, msg.CodeHash, msg.Creator)
	if aggregatePrevote.Hash != hash.String() {
		return nil, types.ErrVerificationFailed.Wrapf("must be given %s not %s", aggregatePrevote.Hash, hash)
	}

	// // Filter out rates which aren't included in the AcceptList
	// // This is also needed for slashing; in the end blocker we are checking if validator voted
	// // for all required currencies. If they missed some, then we increase their slashing counter.
	// filteredTuples := types.ExchangeRateTuples{}
	// for _, tuple := range exchangeRateTuples {
	// 	if params.AcceptList.Contains(tuple.Denom) {
	// 		filteredTuples = append(filteredTuples, tuple)
	// 	}
	// }

	// Move aggregate prevote to aggregate vote with given CodeHash
	aggregateCodeHashVote := types.CodeHashVote{}
	ms.SetAggregateExchangeRateVote(ctx, valAddr, types.NewAggregateCodeHashVote(filteredTuples, valAddr))
	ms.DeleteAggregateExchangeRatePrevote(ctx, valAddr)

	return &types.MsgAggregateCodeHashVoteResponse{}, nil
}
