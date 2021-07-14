package root

import (
	"github.com/redforester/rf-cli/pkg/command/config"
	"github.com/redforester/rf-cli/pkg/command/version"
	"github.com/spf13/cobra"
)

var example = `
$ rf config edit
$ rf ext create
`

func NewCmdRoot() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:           "rf <command> <subcommand> [flags]",
		Short:         "CLI include some shortcut for RedForester",
		SilenceErrors: true,
		SilenceUsage:  true,
		Example:       example,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.AddCommand(version.NewCmdVersion())
	rootCmd.AddCommand(config.NewCmdConfig())

	return rootCmd
}
