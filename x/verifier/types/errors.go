package types

// DONTCOVER

import (
	fmt "fmt"

	"cosmossdk.io/errors"
	"github.com/cometbft/cometbft/crypto/tmhash"
)

// x/verifier module sentinel errors
var (
	ErrNoPrevote               = errors.Register(ModuleName, 2, "no prevote")
	ErrNoVote                  = errors.Register(ModuleName, 3, "no vote")
	ErrNoVotingPermission      = errors.Register(ModuleName, 4, "unauthorized voter")
	ErrInvalidHash             = errors.Register(ModuleName, 5, "invalid hash")
	ErrInvalidHashLength       = errors.Register(ModuleName, 6, fmt.Sprintf("invalid hash length; should equal %d", tmhash.TruncatedSize)) //nolint: lll
	ErrVerificationFailed      = errors.Register(ModuleName, 7, "hash verification failed")
	ErrRevealPeriodMissMatch   = errors.Register(ModuleName, 8, "reveal period of submitted vote does not match with registered prevote") //nolint: lll
	ErrInvalidSaltLength       = errors.Register(ModuleName, 9, "invalid salt length; must be 64")
	ErrInvalidSaltFormat       = errors.Register(ModuleName, 10, "invalid salt format")
	ErrNoAggregatePrevote      = errors.Register(ModuleName, 11, "no aggregate prevote")
	ErrNoAggregateVote         = errors.Register(ModuleName, 12, "no aggregate vote")
	ErrUnknownDenom            = errors.Register(ModuleName, 13, "unknown denom")
	ErrNegativeOrZeroRate      = errors.Register(ModuleName, 14, "invalid exchange rate; should be positive")
	ErrExistingPrevote         = errors.Register(ModuleName, 15, "prevote already submitted for this voting period")
	ErrBallotNotSorted         = errors.Register(ModuleName, 16, "ballot must be sorted before this operation")
	ErrNoHistoricPrice         = errors.Register(ModuleName, 18, "no historic price for this denom at this block")
	ErrNoMedian                = errors.Register(ModuleName, 19, "no median for this denom at this block")
	ErrNoMedianDeviation       = errors.Register(ModuleName, 20, "no median deviation for this denom at this block")
	ErrMalformedLatestAvgPrice = errors.Register(ModuleName, 21, "malformed latest avg price, expecting one byte")
	ErrContractInfoNotFound    = errors.Register(ModuleName, 22, "ContractInfo not found")
	// This error we will get in GetPrevoteTime for findinng prevote time for any contract info
	ErrVoteBlockTimeNotFound  = errors.Register(ModuleName, 23, "Vote blocks not found")
	ErrPreVoteTimePassed      = errors.Register(ModuleName, 24, "Prevote time passed for this Application Id")
	ErrVoteTimePassed         = errors.Register(ModuleName, 25, "Vote time passed for this Application Id")
	ErrNoPendingContractfound = errors.Register(ModuleName, 26, "No pending contract to verify")
)
