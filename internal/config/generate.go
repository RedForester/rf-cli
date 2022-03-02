package config

import (
	"github.com/deissh/rf-cli/internal/utils"
	"os"
)

func Generate() error {
	path := GetConfigFile()

	func() bool {
		s := utils.PrintSpinner("Checking configuration...")
		defer s.Stop()

		return FileExists(GetConfigFile())
	}()

	if err := func() error {
		s := utils.PrintSpinner("Creating new configuration...")
		defer s.Stop()

		return create()
	}(); err != nil {
		return err
	}

	return Write(path)
}

func create() error {
	const perm = 0o700
	path := GetConfigHome()

	if _, err := os.Stat(path); err != nil {
		if err := os.MkdirAll(path, perm); err != nil {
			return err
		}
	}

	file := GetConfigFile()
	if FileExists(file) {
		if err := os.Rename(file, file+".bkp"); err != nil {
			return err
		}
	}
	_, err := os.OpenFile(file, os.O_CREATE, perm)

	return err
}
