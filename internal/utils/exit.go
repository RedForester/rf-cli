package utils

import (
	"errors"
	"github.com/deissh/rf-cli/pkg/log"
	"os"
)

func Exit(err string) {
	if err != "" {
		log.Error(errors.New(err))
	}
	os.Exit(1)
}

func ExitIfError(err error) {
	if err == nil {
		return
	}

	log.Error(err)
	os.Exit(1)
}
