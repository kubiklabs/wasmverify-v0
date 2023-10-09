package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAggregateCodeHashPrevote{}, "verifier/AggregateCodeHashPrevote", nil)
	cdc.RegisterConcrete(&MsgApplyVerifyApplication{}, "verifier/ApplyVerifyApplication", nil)
	cdc.RegisterConcrete(&MsgUpdateBlockTime{}, "verifier/UpdateBlockTime", nil)
	cdc.RegisterConcrete(&MsgAggregateCodeHashVote{}, "verifier/AggregateCodeHashVote", nil)
	cdc.RegisterConcrete(&MsgFinalVerification{}, "verifier/FinalVerification", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAggregateCodeHashPrevote{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApplyVerifyApplication{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateBlockTime{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAggregateCodeHashVote{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinalVerification{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
