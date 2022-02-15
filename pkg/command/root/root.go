package root

import (
	"github.com/deissh/rf-cli/pkg/command/config"
	"github.com/deissh/rf-cli/pkg/command/extension"
	"github.com/deissh/rf-cli/pkg/command/version"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/deissh/rf-cli/pkg/utils"
	"github.com/spf13/cobra"
)

var example = `
$ rf config edit
$ rf extension init
`

func NewCmdRoot(f *factory.Factory) *cobra.Command {
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

	rootCmd.AddCommand(version.NewCmdVersion(f))
	rootCmd.AddCommand(config.NewCmdConfig(f))
	rootCmd.AddCommand(extension.NewCmdExt(f))

	utils.DisableAuthCheck(rootCmd)

	return rootCmd
}
