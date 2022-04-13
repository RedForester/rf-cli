package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deissh/rf-cli/internal/build"
	configCmd "github.com/deissh/rf-cli/internal/cmd/config"
	extensionCmd "github.com/deissh/rf-cli/internal/cmd/extension"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/log"
	"github.com/spf13/cobra"
)

var (
	configPath string
	debug      bool

	username    string
	askPassword bool
	apiUrl      string
)

func init() {
	cobra.OnInitialize(func() {
		if debug {
			log.Level = log.DebugLevel
		}

		log.Debug("RF CLI version %s (%s)", build.Version, build.Date)

		path := config.GetConfigFile()

		if configPath != "" {
			path = configPath
		}

		if err := config.Load(path); err != nil {
			log.Warn("Config not loaded, %s", err)
			log.Warn("Run command to configure the tool.")
			log.Warn(" $ rf-cli config init --help")
			return
		}

		log.Debug("Using config file: %s", config.GetConfigFile())

		setupClient()
	})
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rf <command> <subcommand> [flags]",
		Short:   "CLI include some shortcuts for RedForester",
		Version: build.Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVarP(
		&configPath, "config", "c", config.GetConfigFile(),
		"Config file",
	)
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Turn on debug output")
	cmd.PersistentFlags().StringVar(
		&apiUrl, "api", factory.BaseRFUrl,
		"Override api redforester url from config",
	)
	cmd.PersistentFlags().StringVarP(
		&username, "username", "U", "", "Username (for example \"username@mail.com\")",
	)
	cmd.PersistentFlags().BoolVarP(
		&askPassword, "password", "W", false, "Force password prompt (should happen automatically)",
	)

	addChildCommands(cmd)

	return cmd
}

func addChildCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		configCmd.NewCmdConfig(),
		extensionCmd.NewCmdExtension(),
	)
}

func md5FromString(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

func setupClient() {
	baseRFUrl := config.Config.Rf.BaseURL
	userUsername := config.Config.Client.Username
	pwdHash := config.Config.Client.PasswordHash

	if username != "" {
		userUsername = username
	}
	if pwdHash == "" || askPassword || username != "" {
		var password string
		prompt := &survey.Password{Message: "Password:"}
		if err := survey.AskOne(prompt, &password); err != nil {
			utils.Exit("invalid")
		}
		pwdHash = md5FromString(password)
	}

	if apiUrl != "" && apiUrl != factory.BaseRFUrl {
		baseRFUrl = apiUrl
	}

	factory.ClientInstance = factory.NewClient(baseRFUrl, userUsername, pwdHash)
}
