package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var defaultValue = []byte(`[client]
username = "not_set"
password_hash = "not_set"

[rf]
base_url = "https://app.redforester.com"
`)

// InitConfig restore config structure in path
// configPath is optional
func InitConfig() error {
	if !fileExists(Path) {
		fmt.Println("Creating config")
		if err := os.MkdirAll(Dir, 0755); err != nil {
			return err
		}

		if err := ioutil.WriteFile(Path, defaultValue, 0644); err != nil {
			return err
		}
	}

	buf, err := os.OpenFile(Path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	viper.SetConfigType("toml")
	viper.AddConfigPath(Path)
	err = viper.ReadConfig(buf)

	return nil
}
