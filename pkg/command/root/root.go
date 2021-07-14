package root

import (
	"github.com/redforester/rf-cli/pkg/command/config"
	"github.com/redforester/rf-cli/pkg/command/version"
	"github.com/spf13/cobra"
)

var example = `
$ rf config init
$ rf ext create
`

func NewCmdRoot() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "rf <command> <subcommand> [flags]",
		Short: "Include some shortcut for RedForester",
		SilenceErrors: true,
		SilenceUsage:  true,
		Example: example,
	}

	rootCmd.SetUsageFunc(rootUsageFunc)

	rootCmd.AddCommand(version.NewCmdVersion())
	rootCmd.AddCommand(config.NewCmdConfig())

	return rootCmd
}

func rootUsageFunc(command *cobra.Command) error {
	command.Printf("Usage:  %s", command.UseLine())

	subcommands := command.Commands()
	if len(subcommands) > 0 {
		command.Print("\n\nAvailable commands:\n")
		for _, c := range subcommands {
			if c.Hidden {
				continue
			}
			command.Printf("  %s\n", c.Name())
		}
		return nil
	}

	flagUsages := command.LocalFlags().FlagUsages()
	if flagUsages != "" {
		command.Println("\n\nFlags:")
		command.Print(flagUsages)
	}
	return nil
}