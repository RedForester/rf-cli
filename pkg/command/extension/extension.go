package extension

import (
	initCmd "github.com/deissh/rf-cli/pkg/command/extension/init"
	"github.com/deissh/rf-cli/pkg/command/extension/list"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/spf13/cobra"
)

func NewCmdExt(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "extension <command>",
		Short:   "Manage extensions",
		Aliases: []string{"ext"},
	}

	cmd.AddCommand(initCmd.NewCmdExtInit(f))
	cmd.AddCommand(list.NewCmdExtList(f))

	return cmd
}
