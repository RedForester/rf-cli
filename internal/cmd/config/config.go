package config

import (
	"github.com/spf13/cobra"
)

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "configurations for RF CLI",
	}

	addChildCommands(cmd)

	return cmd
}

func addChildCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		NewCmdInit(),
	)
}
