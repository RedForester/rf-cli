package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Rf     RfConfig     `yaml:"rf"`
	Client ClientConfig `yaml:"client"`
}

type RfConfig struct {
	BaseURL string `yaml:"base_url"`
}

type ClientConfig struct {
	Username     string `yaml:"username"`
	PasswordHash string `yaml:"password_hash"`
}

func New() *Config {
	return &Config{
		Rf: RfConfig{BaseURL: "asd"},
	}
}

func (c *Config) write() error {
	viper.SetConfigName(FileName)
	viper.SetConfigType(FileExt)
	viper.SetConfigFile(GetConfigFile())

	viper.Set(".", c)

	return viper.WriteConfig()
}
