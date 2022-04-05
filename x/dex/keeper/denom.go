// x/dex/keeper/denom.go
package keeper

import (

	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
)

func (k keeper) SaveVoucherDenom(ctx sdk.Context, port string, channel string, denom string) {
	voucher := VoucherDenom(port, channel, denom)

	// Store the origin denom
	_, saved := k.GetDenomTrace(ctx, voucher)
	if !saved {
		k.SetDenomTrace(ctx, types.DenomTrace {
			Index:		voucher,
			Port:		port,
			Channel:	channel,
			Origin:		denom,
		})
	}
}

func VoucherDenom(port string, channel string, denom string) string {
	// prefix
	sourcePrefix := ibctransfertypes.GetDenomPrefix(port, channel)
	// source prefix contains the trailing "/"
	prefixedDenom := sourcePrefix + denom
	// construct the denomination trace form the full raw denomination
	denomTrace := ibctransfertypes.ParseDenomTrace(prefixedDenom)
	voucher := denomTrace.IBCDenom()
	return voucher [:16]
}

func (k keeper) OriginalDenom(ctx sdk.Context, port string, channel string, voucher string) (string, bool) {
	trace, exist := k.GetDenomTrace(ctx, voucher)
	if exist {
		// Check if original port and channel
		if trace.Port == port && trace.Channel == channel {
			return trace.Origin, true
		}
	}
	// not the original chain
	return "", false
}

