package main

import (
	"fmt"
	"github.com/deissh/rf-cli/pkg/command/root"
	"github.com/deissh/rf-cli/pkg/factory"
	"os"
)

func main() {
	os.Exit(Execute())
}

// Execute and return status code
func Execute() int {
	cmdFactory := factory.New()

	rootCmd := root.NewCmdRoot(cmdFactory)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
