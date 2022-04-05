package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	"github.com/lostak/interchange/x/dex/types"
)

func (k msgServer) SendSellOrder(goCtx context.Context, msg *types.MsgSendSellOrder) (*types.MsgSendSellOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// if an order book doesn't exist, throw an error
	_, found := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.AmountDenom, msg.PriceDenom)
	if !found {
		return &types.MsgSendSellOrderResponse{}, errors.New("the pair doesn't exist")
	}

	//Get sender's address
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgSendSellOrderResponse{}, err
	}

	//Use SafeBurn to ensure no new native tokens are minted
	if err := k.SafeBurn(ctx, msg.Port, msg.ChannelID, sender, msg.AmountDenom, msg.Amount); err != nil {
		return &types.MsgSendSellOrderResponse{}, err
	}

	// svae the voucher received on the other chanin, to have the ability to resolve it into the original denom
	k.SaveVoucherDenom(ctx, msg.Port, msg.Channel, msg.AmountDenom)
	var packet types.SellOrderPacketData
	packet.Seller = msg.Creator
	packet.AmountDenom = msg.AmountDenom
	packet.Amount = msg.Amount
	packet.PriceDenom = msg.PriceDenom
	packet.Price = msg.Price

	// Transmit the packet
	err := k.TransmitSellOrderPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendSellOrderResponse{}, nil
}
