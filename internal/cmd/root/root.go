package root

import (
	"fmt"
	initCmd "github.com/deissh/rf-cli/internal/cmd/init"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	configPath string
	debug      bool
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rf <command> <subcommand> [flags]",
		Short: "CLI include some shortcuts for RedForester",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			subCmd := cmd.Name()
			if !cmdRequireToken(subCmd) {
				return
			}

			if configPath != "" {
				viper.SetConfigFile(configPath)
			} else {
				home, err := config.GetConfigHome()
				if err != nil {
					fmt.Println("Missing configuration file.")
					fmt.Println("Run 'rf init' to configure the tool.")
					os.Exit(1)
					return
				}

				viper.AddConfigPath(fmt.Sprintf("%s/%s", home, config.Dir))
				viper.SetConfigName(config.FileName)
			}

			viper.AutomaticEnv()
			viper.SetEnvPrefix("rf")

			if err := viper.ReadInConfig(); err == nil && debug {
				fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
			}
		},
	}

	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Turn on debug output")

	addChildCommands(cmd)

	return cmd
}

func addChildCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		initCmd.NewCmdInit(),
	)
}

func cmdRequireToken(cmd string) bool {
	allowList := []string{
		"init",
		"help",
		"jira",
		"version",
		"completion",
		"man",
	}

	for _, item := range allowList {
		if item == cmd {
			return false
		}
	}

	return true
}
