package types

const (
	// ModuleName defines the module name
	ModuleName = "verifier"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_verifier"

	// total number of verified contracts till now
	ContractCountKey                = "ContractCountKey"
	ContractCompilationBlockTimeKey = "ContractCompilationBlockTime"
	PrevoteBlockTimeKey             = "PrevoteBlockTimeKey"
	VoteBlockTimeKey                = "VoteBlockTimeKey"
	OngoingValidationIdKey          = "OngoingBlockBlockId"
)

var (
	KeyPrefixAggregateCodeHashPrevote = []byte{1} // prefix for each key to a aggregate prevote
	KeyPrefixAggregateCodeHashVote    = []byte{2} // prefix for each key to a aggregate vote
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	// PendingContractsKey      = "PendingContracts/value/"
	ContractInfoKey          = "ContractInfo/value/"
	PendingContractsCountKey = "PendingContracts/count/"
)
