package main

import (
	"github.com/redforester/rf-cli/internal/config"
	"github.com/redforester/rf-cli/pkg/command/root"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	os.Exit(Execute())
}

// Execute and return status code
func Execute() int {
	cobra.OnInitialize(config.InitConfig)

	rootCmd := root.NewCmdRoot()

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}