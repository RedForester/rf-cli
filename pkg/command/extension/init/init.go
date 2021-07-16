package init

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func NewCmdExtInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create new manifest",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {

}

func interactiveExtInfo() {
	q := []*survey.Question{
		{
			Name: "chooseGitIgnore",
			Prompt: &survey.Select{
				Message: "Choose a .gitignore template",
			},
		},
	}

	err := survey.Ask(q, nil)
	if err != nil {
		fmt.Errorf("err: %e", err)
		return
	}
}
