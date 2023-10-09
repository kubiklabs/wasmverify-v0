package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAggregateCodeHashPrevote = "aggregate_code_hash_prevote"

var _ sdk.Msg = &MsgAggregateCodeHashPrevote{}

func NewMsgAggregateCodeHashPrevote(applicationId uint64, validator sdk.ValAddress, hash string) *MsgAggregateCodeHashPrevote {
	return &MsgAggregateCodeHashPrevote{
		ApplicationId: applicationId,
		Validator:     validator.String(),
		Hash:          hash,
	}
}

func (msg *MsgAggregateCodeHashPrevote) Route() string {
	return RouterKey
}

func (msg *MsgAggregateCodeHashPrevote) Type() string {
	return TypeMsgAggregateCodeHashPrevote
}

func (msg *MsgAggregateCodeHashPrevote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAggregateCodeHashPrevote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAggregateCodeHashPrevote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
