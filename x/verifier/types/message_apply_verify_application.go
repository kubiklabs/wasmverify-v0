package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApplyVerifyApplication = "apply_verify_application"

var _ sdk.Msg = &MsgApplyVerifyApplication{}

func NewMsgApplyVerifyApplication(creator string, offchainCodeUrl string) *MsgApplyVerifyApplication {
	return &MsgApplyVerifyApplication{
		Creator:         creator,
		OffchainCodeUrl: offchainCodeUrl,
	}
}

func (msg *MsgApplyVerifyApplication) Route() string {
	return RouterKey
}

func (msg *MsgApplyVerifyApplication) Type() string {
	return TypeMsgApplyVerifyApplication
}

func (msg *MsgApplyVerifyApplication) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApplyVerifyApplication) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApplyVerifyApplication) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
