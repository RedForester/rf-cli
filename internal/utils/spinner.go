package utils

import (
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"time"
)

func PrintSpinner(msg string) *spinner.Spinner {
	const refreshRate = 100 * time.Millisecond

	s := spinner.New(
		spinner.CharSets[14],
		refreshRate,
		spinner.WithSuffix(" "+msg),
		spinner.WithHiddenCursor(true),
		spinner.WithWriter(color.Error),
	)
	s.Start()

	return s
}
