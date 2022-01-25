package config

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/spf13/viper"
	"os"
)

func (c Config) Generate() error {
	func() bool {
		s := utils.PrintSpinner("Checking configuration...")
		defer s.Stop()

		return FileExists(viper.ConfigFileUsed())
	}()

	if err := func() error {
		s := utils.PrintSpinner("Creating new configuration...")
		defer s.Stop()

		home := GetConfigHome()

		return create(home, fmt.Sprintf("%s.%s", FileName, FileExt))
	}(); err != nil {
		return err
	}

	return c.write()
}

func create(path, name string) error {
	const perm = 0o700

	if _, err := os.Stat(path); err != nil {
		if err := os.MkdirAll(path, perm); err != nil {
			return err
		}
	}

	file := fmt.Sprintf("%s/%s", path, name)
	if FileExists(file) {
		if err := os.Rename(file, file+".bkp"); err != nil {
			return err
		}
	}
	_, err := os.Create(file)

	return err
}
