package init

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/log"
	"github.com/deissh/rf-cli/pkg/manifest"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create a manifest.yaml file",
		Long:  "Can be used to set up a new extension",
		Aliases: []string{
			"i",
		},
		Run: run,
	}

	cmd.Flags().StringP("file", "f", "manifest.yaml", "file <path>")
	cmd.Flags().BoolP("yes", "y", false, "Automatically answer \"yes\" to any prompts")

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	forceYes, _ := cmd.Flags().GetBool("yes")
	path, err := cmd.Flags().GetString("file")
	utils.ExitIfError(err)

	info, err := loadManifest(path)
	utils.ExitIfError(err)
	log.Debug("data: %d", *info)

	fmt.Print("Press ^C at any time to quit.\n\n")
	err = askBaseExtInfo(info)
	utils.ExitIfError(err)

	err = view.NewManifest(info).Render()
	utils.ExitIfError(err)

	if ok := utils.Confirm(forceYes); !ok {
		utils.Exit("aborted")
	}

	err = createManifest(path, info)
	utils.ExitIfError(err)

	fmt.Println()
	fmt.Printf("The extension \"%s\" is initialized, to register the extension in the registry, run the command:\n", info.Name)
	fmt.Println(" $ rf-cli extension register --help")
}

func loadManifest(path string) (*manifest.Manifest, error) {
	info := &manifest.Manifest{
		Email: config.Config.Client.Username,
		ExtensionUser: manifest.ExtUser{
			FirstName: "Test",
			LastName:  "Extension",
		},
	}

	if !utils.FileExists(path) {
		return info, nil
	}

	return manifest.ReadByPath(path)
}

func createManifest(path string, info *manifest.Manifest) error {
	err := utils.CreateFileAndBackup(path)
	if err != nil {
		return err
	}

	return manifest.WriteByPath(path, info)
}

func askBaseExtInfo(info *manifest.Manifest) error {
	answers := struct {
		Name        string
		Description string
		Email       string
		Username    string
		FirstName   string
		LastName    string
	}{
		Name:        info.Name,
		Description: info.Description,
		Email:       info.Email,
		Username:    info.ExtensionUser.Username,
		FirstName:   info.ExtensionUser.FirstName,
		LastName:    info.ExtensionUser.LastName,
	}

	q := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Default: answers.Name,
				Message: "Extensions name:",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "description",
			Prompt: &survey.Input{
				Default: answers.Description,
				Message: "Description (optional):",
			},
		},
		{
			Name: "email",
			Prompt: &survey.Input{
				Message: "Support email:",
				Default: answers.Email,
			},
			Validate: survey.Required,
		},
		{
			Name: "username",
			Prompt: &survey.Input{
				Default: answers.Username,
				Message: "Extension username (will be used as a unique extension identifier)",
			},
			Validate: survey.Required,
		},
		{
			Name: "firstName",
			Prompt: &survey.Input{
				Message: "Extension user firstname (optional)",
				Default: answers.FirstName,
			},
		},
		{
			Name: "lastName",
			Prompt: &survey.Input{
				Message: "Extension user lastname (optional)",
				Default: answers.LastName,
			},
		},
	}

	if err := survey.Ask(q, &answers); err != nil {
		return err
	}

	info.Name = answers.Name
	info.Description = answers.Description
	info.Email = answers.Email
	info.ExtensionUser.Username = answers.Username
	info.ExtensionUser.FirstName = answers.FirstName
	info.ExtensionUser.LastName = answers.LastName
	if info.ExtensionUser.AvatarUrl == "" {
		info.ExtensionUser.AvatarUrl = "https://avatars.redforester.com/?name=" + answers.Name
	}

	if err := info.Validate(); err != nil {
		return err
	}

	return nil
}
