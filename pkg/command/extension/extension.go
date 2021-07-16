package extension

import (
	"github.com/deissh/rf-cli/pkg/command/extension/init"
	"github.com/spf13/cobra"
)

func NewCmdExt() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "extension <command>",
		Short:   "Manage extensions",
		Aliases: []string{"ext"},
	}

	cmd.AddCommand(init.NewCmdExtInit())

	return cmd
}
