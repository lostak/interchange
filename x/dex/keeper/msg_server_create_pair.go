package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	"github.com/lostak/interchange/x/dex/types"
)

func (k msgServer) SendCreatePair(goCtx context.Context, msg *types.MsgSendCreatePair) (*types.MsgSendCreatePairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//Get an order book index
	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.SourceDenom, msg.TargetDenom)
	//If order book is found return error
	_, found := k.GetSellOrderBook(ctx, pairIndex)
	if found {
		return &types.MsgSendCreatePairResponse{}, errors.New("the pair already exists")
	}

	// Construct the packet
	var packet types.CreatePairPacketData

	packet.SourceDenom = msg.SourceDenom
	packet.TargetDenom = msg.TargetDenom

	// Transmit the packet
	err := k.TransmitCreatePairPacket(
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

	return &types.MsgSendCreatePairResponse{}, nil
}
