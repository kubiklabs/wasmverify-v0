package keeper

import (
	"verifier/x/verifier/types"
)

var _ types.QueryServer = Keeper{}
