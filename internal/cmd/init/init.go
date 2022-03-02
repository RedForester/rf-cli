package init

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/rf"
	"github.com/spf13/cobra"
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

func run(cmd *cobra.Command, args []string) {
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
	fmt.Println("Configuration generated")
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
	if err != nil {
		return err
	}

	useCurrentUser := true
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Do you like use user %s?", user.Username),
	}
	_ = survey.AskOne(prompt, &useCurrentUser)
	if !useCurrentUser {
		return setUserCreds(creds)
	}

	return nil
}

func tryLogin(creds *UserCreds) (*rf.User, error) {
	cl := factory.NewClient(creds.Email, md5FromString(creds.Password))

	return cl.User.GetMe()
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
