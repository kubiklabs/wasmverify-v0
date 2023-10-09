package types

import (
	"testing"

	"verifier/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAggregateCodeHashPrevote_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAggregateCodeHashPrevote
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAggregateCodeHashPrevote{
				Operator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAggregateCodeHashPrevote{
				Operator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
