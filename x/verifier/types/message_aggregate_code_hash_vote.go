package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAggregateCodeHashVote = "aggregate_code_hash_vote"

var _ sdk.Msg = &MsgAggregateCodeHashVote{}

func NewMsgAggregateCodeHashVote(applicationId uint64, creator string, salt string, codeHash string) *MsgAggregateCodeHashVote {
	return &MsgAggregateCodeHashVote{
		ApplicationId: applicationId,
		Creator:       creator,
		Salt:          salt,
		CodeHash:      codeHash,
	}
}

func (msg *MsgAggregateCodeHashVote) Route() string {
	return RouterKey
}

func (msg *MsgAggregateCodeHashVote) Type() string {
	return TypeMsgAggregateCodeHashVote
}

func (msg *MsgAggregateCodeHashVote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAggregateCodeHashVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAggregateCodeHashVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
