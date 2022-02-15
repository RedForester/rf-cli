package config

import (
	"fmt"
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
	config := Config{
		Rf: RfConfig{
			BaseURL: "https://beta.app.redforester.com",
		},
	}

	if err := viper.UnmarshalKey("config", &config); err != nil {
		fmt.Println("WARN: unmarshall config error")
	}
	return &config
}

func (c *Config) write() error {
	return viper.WriteConfig()
}
