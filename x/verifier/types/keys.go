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

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PendingContractsKey      = "PendingContracts/value/"
	PendingContractsCountKey = "PendingContracts/count/"
)
