package config

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
	return &Config{}
}
