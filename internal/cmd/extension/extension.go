package extension

import (
	initCmd "github.com/deissh/rf-cli/internal/cmd/extension/init"
	"github.com/deissh/rf-cli/internal/cmd/extension/list"
	"github.com/deissh/rf-cli/internal/cmd/extension/register"
	"github.com/deissh/rf-cli/internal/cmd/extension/update"
	"github.com/deissh/rf-cli/internal/cmd/extension/view"
	"github.com/spf13/cobra"
)

func NewCmdExtension() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "extension",
		Short:   "Manage extension",
		Aliases: []string{"ext", "e"},
	}

	cmd.AddCommand(
		list.NewCmd(),
		view.NewCmd(),
		initCmd.NewCmd(),
		register.NewCmd(),
		update.NewCmd(),
	)

	return cmd
}
