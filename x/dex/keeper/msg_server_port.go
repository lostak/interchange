package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lostak/interchange/x/dex/types"
)

func (k msgServer) Port(goCtx context.Context, msg *types.MsgPort) (*types.MsgPortResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPortResponse{}, nil
}
