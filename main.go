package main

import (
	"github.com/deissh/rf-cli/internal/cmd"
	"github.com/deissh/rf-cli/pkg/log"
)

func main() {
	rootCmd := cmd.NewCmdRoot()
	if _, err := rootCmd.ExecuteC(); err != nil {
		log.Error(err)
	}
}
