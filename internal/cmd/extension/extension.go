package extension

import (
	"github.com/deissh/rf-cli/internal/config"
	"github.com/spf13/cobra"
)

func NewCmdExtension() *cobra.Command {
	return &cobra.Command{
		Use:     "extension",
		Short:   "Manage extension",
		Aliases: []string{"ext", "e"},
		Run:     run,
	}
}

func run(*cobra.Command, []string) {
	_ = config.New()
}
