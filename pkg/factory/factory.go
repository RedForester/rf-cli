package factory

import (
	"github.com/deissh/rf-cli/internal/config"
	"net/http"
)

type Factory struct {
	HttpClient func() (*http.Client, error)
	Config     func() (*config.Config, error)
}

func New() *Factory {
	f := &Factory{
		Config: configFunc(),
	}

	f.HttpClient = httpClientFunc(f)

	return f
}

func httpClientFunc(f *Factory) func() (*http.Client, error) {
	return func() (*http.Client, error) {
		cfg, err := f.Config()
		if err != nil {
			return nil, err
		}

		return newHTTPClient(cfg)
	}
}

func configFunc() func() (*config.Config, error) {
	var cachedConfig *config.Config
	var configError error

	return func() (*config.Config, error) {
		if cachedConfig != nil || configError != nil {
			return cachedConfig, configError
		}

		configError = config.InitConfig()
		cachedConfig = config.New()

		return cachedConfig, configError
	}
}
