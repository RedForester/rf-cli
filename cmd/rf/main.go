package main

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/cmd/root"
	"os"
)

func main() {
	rootCmd := root.NewCmdRoot()
	if _, err := rootCmd.ExecuteC(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
