package config

import (
	"github.com/deissh/rf-cli/pkg/command/config/edit"
	"github.com/spf13/cobra"
)

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config <command>",
		Short: "Manage configuration",
	}

	cmd.AddCommand(edit.NewCmdConfigEdit())

	return cmd
}
