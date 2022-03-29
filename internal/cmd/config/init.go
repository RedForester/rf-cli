package config

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/log"
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Init initializes config",
		Long:    "Init initializes configuration required for the tool to work properly.",
		Aliases: []string{"initialize", "configure", "config", "setup"},
		Run:     run,
	}
	cmd.Flags().StringP("username", "u", "", "RedForester email")
	cmd.Flags().StringP("password", "p", "", "RedForester password")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	email, _ := cmd.Flags().GetString("username")
	pwd, _ := cmd.Flags().GetString("password")
	creds := &UserCreds{email, pwd}

	if email == "" && pwd == "" {
		err := askUserCredentials(creds)
		utils.ExitIfError(err)
	}

	err := setUserCreds(creds)
	utils.ExitIfError(err)

	config.Config.Client.Username = creds.Email
	config.Config.Client.PasswordHash = md5FromString(creds.Password)

	err = config.Generate()
	utils.ExitIfError(err)
	log.Info("Configuration generated")
	log.Info("Path: %s", config.GetConfigFile())
}

type UserCreds struct {
	Email    string
	Password string
}

func md5FromString(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

func setUserCreds(creds *UserCreds) error {
	user, err := func() (*rf.User, error) {
		s := utils.PrintSpinner("Checking user credentials...")
		defer s.Stop()

		return tryLogin(creds)
	}()
	if err != nil || user.Username == "nobody" {
		return errors.New("wrong login or password")
	}

	r := view.CurrentUser{Data: user, Writer: os.Stdout}
	if err := r.Render(); err != nil {
		return err
	}

	if !utils.Confirm(false) {
		utils.Exit("aborted")
	}

	return nil
}

func tryLogin(creds *UserCreds) (*rf.User, error) {
	client := factory.NewClient(factory.BaseRFUrl, creds.Email, md5FromString(creds.Password))

	return client.User.GetMe()
}

func askUserCredentials(creds *UserCreds) error {
	q := []*survey.Question{
		{
			Name: "email",
			Prompt: &survey.Input{
				Message: "User email:",
				Default: creds.Email,
			},
			Validate: survey.Required,
		},
		{
			Name: "password",
			Prompt: &survey.Password{
				Message: "Password:",
			},
			Validate: survey.Required,
		},
	}

	return survey.Ask(q, creds)
}
