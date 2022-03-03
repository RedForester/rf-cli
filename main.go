package main

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewCmdRoot()
	if _, err := rootCmd.ExecuteC(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
