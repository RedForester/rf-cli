package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Rf     RfConfig
	Client ClientConfig
}

type RfConfig struct {
	BaseURL string `toml:"base_url"`
}

type ClientConfig struct {
	Username     string `toml:"username"`
	PasswordHash string `toml:"password_hash"`
}

func New() *Config {
	c := Config{}

	err := viper.Unmarshal(c)
	if err != nil {
		fmt.Println(err)
	}

	baseURL := c.Rf.BaseURL
	if len(baseURL) == 0 {
		fmt.Println("please edit base url first; `rf config edit`")
	}
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "https://" + baseURL
	}
	if strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL[:len(baseURL)-1]
	}

	return &c
}
