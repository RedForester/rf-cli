package utils

import (
	"fmt"
	"os"
)

func Exit(err string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func ExitIfError(err error) {
	if err == nil {
		return
	}

	msg := fmt.Sprintf("Error: %s", err.Error())

	fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(1)
}
