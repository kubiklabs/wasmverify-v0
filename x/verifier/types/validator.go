package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// Claim is an interface that directs its rewards to an attached bank account.
type ValidatorData struct {
	Power     int64
	Validator sdk.ValAddress
}

// NewClaim generates a Claim instance.
func NewValidatorData(power int64, v sdk.ValAddress) ValidatorData {
	return ValidatorData{
		Power:     power,
		Validator: v,
	}
}
