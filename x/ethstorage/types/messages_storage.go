package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateStorage = "create_storage"
	TypeMsgUpdateStorage = "update_storage"
	TypeMsgDeleteStorage = "delete_storage"
)

var _ sdk.Msg = &MsgCreateStorage{}

func NewMsgCreateStorage(creator string) *MsgCreateStorage {
	return &MsgCreateStorage{
		Creator: creator,
	}
}

func (msg *MsgCreateStorage) Route() string {
	return RouterKey
}

func (msg *MsgCreateStorage) Type() string {
	return TypeMsgCreateStorage
}

func (msg *MsgCreateStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err) //nolint:staticcheck
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateStorage{}

func NewMsgUpdateStorage(creator string, id uint64) *MsgUpdateStorage {
	return &MsgUpdateStorage{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgUpdateStorage) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStorage) Type() string {
	return TypeMsgUpdateStorage
}

func (msg *MsgUpdateStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err) //nolint:staticcheck
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteStorage{}

func NewMsgDeleteStorage(creator string, id uint64) *MsgDeleteStorage {
	return &MsgDeleteStorage{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteStorage) Route() string {
	return RouterKey
}

func (msg *MsgDeleteStorage) Type() string {
	return TypeMsgDeleteStorage
}

func (msg *MsgDeleteStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err) //nolint:staticcheck
	}
	return nil
}
