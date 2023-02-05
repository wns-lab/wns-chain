package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/wns-lab/wns-chain/x/wns/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	wnsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "wns transactions subcommands",
		Long:                       "Provides the most common wns logic for upper-level applications",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	wnsTxCmd.AddCommand(
		NewCmdSetMetaData(),
	)

	return wnsTxCmd
}

// NewCmdSetMetaData creates a CLI command for MsgSetMetaData.
func NewCmdSetMetaData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set --owner [owner] --name [domain-name] [owner] [resolver] [ttl]",
		Args:  cobra.ExactArgs(5),
		Short: "set metadata for domain name",
		Long: strings.TrimSpace(fmt.Sprintf(`
			$ %s tx %s set --owner <owner> --name <domain-name> <owner> <resolver> <ttl>`, version.AppName, types.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ttl, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return err
			}

			msg := types.MsgSetMetaData{
				Sender: args[0],
				Name:   args[1],
				Metadata: &types.MetaData{
					Owner:    args[2],
					Resolver: args[3],
					TTL:      ttl,
				},
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
