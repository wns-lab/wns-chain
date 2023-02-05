package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/wns-lab/wns-chain/x/wns/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	wnsQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the wns module",
		Long:                       "",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	wnsQueryCmd.AddCommand(
		GetCmdQueryMetaData(),
	)
	return wnsQueryCmd
}

// GetCmdQueryMetaData implements the query class command.
func GetCmdQueryMetaData() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "metadata [domain-name]",
		Args:    cobra.ExactArgs(1),
		Short:   "query the metadata of a domain name",
		Example: fmt.Sprintf(`$ %s query %s metadata <domain-name>`, version.AppName, types.ModuleName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Metadata(cmd.Context(), &types.QueryMetaDataRequest{
				Name: args[0],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
