package main

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/pkg/command/root"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	os.Exit(Execute())
}

// Execute and return status code
func Execute() int {
	cobra.OnInitialize(OnInitialize)

	rootCmd := root.NewCmdRoot()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func OnInitialize() {
	err := config.InitConfig()
	if err != nil {
		fmt.Println("[error]", err)
		os.Exit(1)
	}
}
