package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Rf RfConfig
	Client ClientConfig
}

type RfConfig struct {
	BaseURL string `toml:"base_url"`
}

type ClientConfig struct {
	Username string `toml:"username"`
	PasswordHash string `toml:"password_hash"`
}

func New() *Config {
	c := Config{}

	err := viper.Unmarshal(c)
	if err != nil {
		log.Fatalln(err)
	}

	baseURL := c.Rf.BaseURL
	if len(baseURL) == 0 {
		log.Fatalln("edit base url first, use `rf config init`")
	}
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "https://" + baseURL
	}
	if strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL[:len(baseURL)-1]
	}

	return &c
}