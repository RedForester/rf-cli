package utils

import "github.com/AlecAivazis/survey/v2"

// Confirm user action
func Confirm(forceYes bool) bool {
	if forceYes {
		return true
	}

	ok := true
	prompt := &survey.Confirm{Message: "Continue?", Default: true}
	if err := survey.AskOne(prompt, &ok); err != nil || !ok {
		return false
	}

	return ok
}
