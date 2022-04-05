package dex_test

import (
	"testing"

	keepertest "github.com/lostak/interchange/testutil/keeper"
	"github.com/lostak/interchange/testutil/nullify"
	"github.com/lostak/interchange/x/dex"
	"github.com/lostak/interchange/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		SellOrderBookList: []types.SellOrderBook{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		BuyOrderBookList: []types.BuyOrderBook{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.SellOrderBookList, got.SellOrderBookList)
	require.ElementsMatch(t, genesisState.BuyOrderBookList, got.BuyOrderBookList)
	// this line is used by starport scaffolding # genesis/test/assert
}
