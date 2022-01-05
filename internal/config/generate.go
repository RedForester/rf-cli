package config

import (
	"fmt"
	"github.com/deissh/rf-cli/internal/utils"
	"github.com/spf13/viper"
	"os"
	"time"
)

func (c Config) Generate() error {
	func() bool {
		s := utils.PrintSpinner("Checking configuration...")
		time.Sleep(time.Second * 10)
		defer s.Stop()

		return FileExists(viper.ConfigFileUsed())
	}()

	if err := func() error {
		s := utils.PrintSpinner("Creating new configuration...")
		defer s.Stop()

		home, err := GetConfigHome()
		if err != nil {
			return err
		}

		return create(fmt.Sprintf("%s/%s/", home, Dir), fmt.Sprintf("%s.%s", FileName, FileExt))
	}(); err != nil {
		return err
	}

	return nil
}

func create(path, name string) error {
	const perm = 0o700

	if _, err := os.Stat(path); err != nil {
		if err := os.MkdirAll(path, perm); err != nil {
			return err
		}
	}

	file := path + name
	if FileExists(file) {
		if err := os.Rename(file, file+".bkp"); err != nil {
			return err
		}
	}
	_, err := os.Create(file)

	return err
}
