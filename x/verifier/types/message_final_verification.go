package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFinalVerification = "final_verification"

var _ sdk.Msg = &MsgFinalVerification{}

func NewMsgFinalVerification(creator string, applicationId uint64, codeId uint64) *MsgFinalVerification {
	return &MsgFinalVerification{
		Creator:       creator,
		ApplicationId: applicationId,
		CodeId:        codeId,
	}
}

func (msg *MsgFinalVerification) Route() string {
	return RouterKey
}

func (msg *MsgFinalVerification) Type() string {
	return TypeMsgFinalVerification
}

func (msg *MsgFinalVerification) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinalVerification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinalVerification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
