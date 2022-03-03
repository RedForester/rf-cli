package extension

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/deissh/rf-cli/pkg/manifest"
	"github.com/deissh/rf-cli/pkg/view"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create new extension and save manifest.yaml",
		Aliases: []string{"c"},
		Run:     runCmdCreate,
	}

	cmd.Flags().StringP("file", "f", "manifest.yaml", "file <path>")
	cmd.Flags().BoolP("yes", "y", false, "Automatically answer \"yes\" to any prompts")

	return cmd
}

func runCmdCreate(cmd *cobra.Command, args []string) {
	forceYes, err := cmd.Flags().GetBool("yes")

	info, err := askBaseExtInfo()
	utils.ExitIfError(err)

	err = view.NewManifest(info).Render()
	utils.ExitIfError(err)

	if !forceYes {
		ok := true
		prompt := &survey.Confirm{Message: "Is this OK?", Default: true}
		if err = survey.AskOne(prompt, &ok); err != nil || !ok {
			utils.Exit("aborted")
		}
	}

	path, err := cmd.Flags().GetString("file")
	utils.ExitIfError(err)

	err = createManifest(path, info)
	utils.ExitIfError(err)
}

func createManifest(path string, info *manifest.Manifest) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := yaml.Marshal(info)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = f.Write(data)

	return err
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
