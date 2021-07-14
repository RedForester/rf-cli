package config

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

var defaultValue = []byte(`[client]
username = "not_set"
password_hash = "not_set"
`)

// InitConfig restore config structure in path
// configPath is optional
func InitConfig() {
	if !fileExists(Path) {
		if err := os.MkdirAll(Dir, 0755); err != nil {
			log.Fatalln(err)
		}

		if err := ioutil.WriteFile(Path, defaultValue, 0644); err != nil {
			log.Fatalln(err)
		}
	}

	buf, err := os.OpenFile(Path, os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	viper.SetConfigType("toml")
	viper.AddConfigPath(Path)
	err = viper.ReadConfig(buf)
}
