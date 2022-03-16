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

		return utils.FileExists(path)
	}()

	if err := func() error {
		s := utils.PrintSpinner("Creating new configuration...")
		defer s.Stop()

		return create(path)
	}(); err != nil {
		return err
	}

	return Write(path)
}

func create(path string) error {
	const perm = 0o700
	//if _, err := os.Stat(path); err != nil {
	//	if err := os.MkdirAll(path, perm); err != nil {
	//		return err
	//	}
	//}

	if utils.FileExists(path) {
		if err := os.Rename(path, path+".bkp"); err != nil {
			return err
		}
	}
	_, err := os.OpenFile(path, os.O_CREATE, perm)

	return err
}
