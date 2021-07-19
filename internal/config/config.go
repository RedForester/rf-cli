package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Config struct {
	Rf     RfConfig     `mapstructure:"rf"`
	Client ClientConfig `mapstructure:"client"`
}

type RfConfig struct {
	BaseURL string `mapstructure:"base_url"`
}

type ClientConfig struct {
	Username     string `mapstructure:"username"`
	PasswordHash string `mapstructure:"password_hash"`
}

func New() *Config {
	c := Config{}

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	baseURL := c.Rf.BaseURL
	if len(baseURL) == 0 {
		fmt.Println("please edit base url first; `rf config edit`")
		os.Exit(1)
	}
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "https://" + baseURL
	}
	if strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL[:len(baseURL)-1]
	}

	return &c
}
