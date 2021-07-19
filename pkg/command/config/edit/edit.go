package edit

import (
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func NewCmdConfigEdit(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Use $EDITOR open config file",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	editor := getEnvOr("EDITOR", "vim")

	command := exec.Command(editor, config.Path)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	_ = command.Run()
}

func getEnvOr(key string, value string) string {
	env := os.Getenv(key)
	if env == "" {
		env = value
	}

	return env
}
