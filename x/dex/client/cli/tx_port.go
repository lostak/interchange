package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/lostak/interchange/x/dex/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPort() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "port [channel] [amount-denom] [price-denom] [order-id]",
		Short: "Cancel a sell order",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChannel := args[0]
			argAmountDenom := args[1]
			argPriceDenom := args[2]
			argOrderID, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPort(
				clientCtx.GetFromAddress().String(),
				argChannel,
				argAmountDenom,
				argPriceDenom,
				argOrderID,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
