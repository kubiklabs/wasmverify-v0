package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBlockTime = "update_block_time"

var _ sdk.Msg = &MsgUpdateBlockTime{}

func NewMsgUpdateBlockTime(creator string, compilationBlocks uint64, prevoteBlocks uint64, voteBlocks uint64) *MsgUpdateBlockTime {
	return &MsgUpdateBlockTime{
		Creator:           creator,
		CompilationBlocks: compilationBlocks,
		PrevoteBlocks:     prevoteBlocks,
		VoteBlocks:        voteBlocks,
	}
}

func (msg *MsgUpdateBlockTime) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBlockTime) Type() string {
	return TypeMsgUpdateBlockTime
}

func (msg *MsgUpdateBlockTime) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBlockTime) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBlockTime) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
