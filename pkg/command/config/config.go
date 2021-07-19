package config

import (
	"github.com/deissh/rf-cli/pkg/command/config/edit"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/spf13/cobra"
)

func NewCmdConfig(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config <command>",
		Short: "Manage configuration",
	}

	cmd.AddCommand(edit.NewCmdConfigEdit(f))

	return cmd
}
