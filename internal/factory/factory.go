package factory

import (
	"github.com/deissh/rf-cli/pkg/log"
	"github.com/deissh/rf-cli/pkg/rf_api"
)

var BaseRFUrl = "https://app.redforester.com"
var ClientInstance = NewClient(BaseRFUrl, "", "")

func NewClient(rfBaseUrl, username, passwordHash string) *rf_api.Client {
	httpClient, err := newHTTPClient(username, passwordHash)
	if err != nil {
		log.Warn("http client not loaded: %e", err)
		return nil
	}

	return rf_api.New(rfBaseUrl, httpClient)
}
