package init

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/manifest"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
	"os"
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
	forceYes, err := cmd.Flags().GetBool("yes")

	fmt.Print("Press ^C at any time to quit.\n\n")
	info, err := askBaseExtInfo()
	utils.ExitIfError(err)

	err = view.NewManifest(info).Render()
	utils.ExitIfError(err)

	if ok := utils.Confirm(forceYes); !ok {
		utils.Exit("aborted")
	}

	path, err := cmd.Flags().GetString("file")
	utils.ExitIfError(err)

	err = createManifest(path, info)
	utils.ExitIfError(err)

	fmt.Println()
	fmt.Printf("The extension \"%s\" is initialized, to register the extension in the registry, run the command:\n", info.Name)
	fmt.Println(" $ rf-cli extension register --help")
}

func createManifest(path string, info *manifest.Manifest) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return manifest.Write(f, info)
}

func askBaseExtInfo() (*manifest.Manifest, error) {
	result := manifest.Manifest{
		Email: config.Config.Client.Username,
	}

	q := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Extensions name:",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "Description",
			Prompt: &survey.Input{
				Message: "Description (optional):",
			},
		},
		{
			Name: "email",
			Prompt: &survey.Input{
				Message: "Author email:",
				Default: result.Email,
			},
			Validate: survey.Required,
		},
	}

	if err := survey.Ask(q, &result); err != nil {
		return nil, err
	}

	if err := result.Validate(); err != nil {
		return nil, err
	}

	return &result, nil
}
