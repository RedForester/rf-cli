package init

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func NewCmdInit() *cobra.Command {
	return &cobra.Command{
		Use:     "init",
		Short:   "Init initializes config",
		Long:    "Init initializes configuration required for the tool to work properly.",
		Aliases: []string{"initialize", "configure", "config", "setup"},
		Run:     run,
	}
}

func run(*cobra.Command, []string) {
	c := config.New()

	if err := c.Generate(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Configuration generated: %s\n", viper.ConfigFileUsed())
}
