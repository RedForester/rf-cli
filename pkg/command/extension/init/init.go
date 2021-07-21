package init

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/pkg/extension"
	"github.com/deissh/rf-cli/pkg/factory"
	"github.com/deissh/rf-cli/pkg/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
)

type Options struct {
	Interactive bool
	File        string
}

func NewCmdExtInit(f *factory.Factory) *cobra.Command {
	opt := Options{
		Interactive: true,
	}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create new manifest",
		Run: func(cmd *cobra.Command, args []string) {
			run(f, cmd, args, opt)
		},
	}

	cmd.Flags().StringVarP(&opt.File, "file", "f", "manifest.yaml", "manifest file path")

	utils.DisableAuthCheck(cmd)

	return cmd
}

func run(f *factory.Factory, cmd *cobra.Command, args []string, opt Options) {
	cfg, err := f.Config()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ext *extension.Extension
	if opt.Interactive {
		ext = interactiveExtInfo(cfg)
	} else {
		// todo
	}

	data, err := yaml.Marshal(ext)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(opt.File, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	file.Write(data)
}

func interactiveExtInfo(cfg *config.Config) *extension.Extension {
	extInfo := &extension.Extension{
		ID:    "NOT_CREATED",
		Email: cfg.Client.Username,
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
				Message: "Description:",
			},
		},
		{
			Name: "email",
			Prompt: &survey.Input{
				Message: "Author email:",
				Default: extInfo.Email,
			},
			Validate: survey.Required,
		},
		{
			Name: "baseURL",
			Prompt: &survey.Input{
				Message: "Extension base url:",
				Help:    "for example: https://localhost:2300",
			},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(q, extInfo)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return extInfo
}
