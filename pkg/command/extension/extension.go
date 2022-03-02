package extension

import (
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/pkg/command/extension/info"
	initCmd "github.com/deissh/rf-cli/pkg/command/extension/init"
	"github.com/deissh/rf-cli/pkg/command/extension/list"
	"github.com/deissh/rf-cli/pkg/command/extension/update"
	"github.com/spf13/cobra"
)

func NewCmdExt(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "extension <command>",
		Short:   "Manage extension",
		Aliases: []string{"extension"},
	}

	cmd.AddCommand(initCmd.NewCmdExtInit(f))
	cmd.AddCommand(list.NewCmdExtList(f))
	cmd.AddCommand(info.NewCmdExtInfo(f))
	cmd.AddCommand(update.NewCmdExtUpdate(f))

	return cmd
}
