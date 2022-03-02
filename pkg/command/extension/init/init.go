package init

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deissh/rf-cli/internal/factory"
	"github.com/deissh/rf-cli/pkg/extension"
	"github.com/deissh/rf-cli/pkg/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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

	file, err := os.OpenFile(opt.File, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var ext *extension.Extension
	if opt.Interactive {
		ext = interactiveExtInfo(cfg)
	} else {
		fmt.Println("todo")
		os.Exit(1)
		return
	}

	data, err := yaml.Marshal(ext)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, _ = file.Write(data)
}

func interactiveExtInfo() *extension.Extension {
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

	err = showCommandList(extInfo)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return extInfo
}

const AddNewCommandSelect = "Add new command"
const CloseNewCommandSelect = "Exit"

func showCommandList(ext *extension.Extension) error {
	//var command extension.Command

	var options []string
	for _, c := range ext.Commands {
		options = append(options, c.Name)
	}

	options = append(options, AddNewCommandSelect, CloseNewCommandSelect)

	var selected string
	prompt := &survey.Select{
		Message: "Extension commands",
		Options: options,
		Default: CloseNewCommandSelect,
	}
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return err
	}

	switch selected {
	case AddNewCommandSelect:
		ext.Commands = append(ext.Commands, *promptCommand())
	case CloseNewCommandSelect:
		return errors.New("closed")
	default:
		fmt.Println("print current command")
	}

	return showCommandList(ext)
}

func promptCommand() *extension.Command {
	var cmd extension.Command

	q := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Command name:",
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
	}

	// todo: action
	// todo: rules

	err := survey.Ask(q, &cmd)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &cmd
}
