package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPort = "port"

var _ sdk.Msg = &MsgPort{}

func NewMsgPort(creator string, channel string, amountDenom string, priceDenom string, orderID int32) *MsgPort {
	return &MsgPort{
		Creator:     creator,
		Channel:     channel,
		AmountDenom: amountDenom,
		PriceDenom:  priceDenom,
		OrderID:     orderID,
	}
}

func (msg *MsgPort) Route() string {
	return RouterKey
}

func (msg *MsgPort) Type() string {
	return TypeMsgPort
}

func (msg *MsgPort) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPort) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPort) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
