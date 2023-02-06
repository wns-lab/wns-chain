package experimental

import (
	gravitycmd "github.com/peggyjv/gravity-bridge/module/v2/cmd/gravity/cmd"
	"github.com/spf13/cobra"
	"github.com/wns-lab/wns-chain/app"
)

// add server commands
func AddCommands(rootCmd *cobra.Command) {
	experimentalCmd := &cobra.Command{
		Use:   "experimental",
		Short: "experimental subcommands (unsafe)",
	}

	experimentalCmd.AddCommand(
		gravitycmd.Commands(app.DefaultNodeHome),
	)

	rootCmd.AddCommand(
		experimentalCmd,
	)
}
