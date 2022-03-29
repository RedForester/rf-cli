package config

import (
	"github.com/deissh/rf-cli/internal/utils"
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

		return utils.CreateFileAndBackup(path)
	}(); err != nil {
		return err
	}

	return Write(path)
}
