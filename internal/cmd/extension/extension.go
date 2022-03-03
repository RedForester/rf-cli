package extension

import (
	"github.com/spf13/cobra"
)

func NewCmdExtension() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "extension",
		Short:   "Manage extension",
		Aliases: []string{"ext", "e"},
	}

	cmd.AddCommand(
		NewCmdList(),
		NewCmdCreate(),
	)

	return cmd
}
