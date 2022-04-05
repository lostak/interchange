package keeper

import (
	"github.com/lostak/interchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
