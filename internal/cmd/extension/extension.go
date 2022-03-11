package extension

import (
	"github.com/deissh/rf-cli/internal/cmd/extension/create"
	"github.com/deissh/rf-cli/internal/cmd/extension/list"
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
		create.NewCmd(),
	)

	return cmd
}
