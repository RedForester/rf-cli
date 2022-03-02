package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Config = New()

type config struct {
	Rf     rfConfig     `yaml:"rf"`
	Client clientConfig `yaml:"client"`
}

type rfConfig struct {
	BaseURL string `yaml:"base_url"`
}

type clientConfig struct {
	Username     string `yaml:"username"`
	PasswordHash string `yaml:"password_hash"`
}

// New create config with default value
func New() *config {
	return &config{
		Rf: rfConfig{
			BaseURL: "https://beta.app.redforester.com",
		},
	}
}

func Load(configPath string) error {
	if FileExists(configPath) != true {
		fmt.Println("config file not exist, please sign in first")
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&Config); err != nil {
		return err
	}

	return nil
}

func Write(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	out, err := yaml.Marshal(Config)
	if err != nil {
		return err
	}

	_, err = file.Write(out)
	if err != nil {
		return err
	}

	return nil
}
