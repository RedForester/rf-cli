package root

import (
	"fmt"
	extensionCmd "github.com/deissh/rf-cli/internal/cmd/extension"
	initCmd "github.com/deissh/rf-cli/internal/cmd/init"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configPath string
	debug      bool
)

func init() {
	cobra.OnInitialize(func() {
		path := config.GetConfigFile()

		if configPath != "" {
			path = configPath
		}

		if config.FileExists(path) != true {
			fmt.Println("Missing configuration file.")
			fmt.Println("Run 'rf init' to configure the tool.")
			fmt.Println()
		}

		viper.SetConfigFile(path)

		viper.AutomaticEnv()
		viper.SetEnvPrefix("RF")

		if err := viper.ReadInConfig(); err != nil && debug {
			fmt.Printf("Config not loaded: %s\n", err)
		}
		if debug {
			fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
		}
	})
}

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

		},
	}

	cmd.PersistentFlags().StringVarP(
		&configPath, "config", "c", "",
		fmt.Sprintf("Config file (default is %s/%s.%s)", config.GetConfigHome(), config.FileName, config.FileExt),
	)
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Turn on debug output")

	addChildCommands(cmd)

	return cmd
}

func addChildCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		initCmd.NewCmdInit(),
		extensionCmd.NewCmdExtension(),
	)
}

func cmdRequireToken(cmd string) bool {
	allowList := []string{
		"init",
		"help",
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
