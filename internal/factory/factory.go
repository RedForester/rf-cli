package factory

import (
	"fmt"
	"github.com/deissh/rf-cli/pkg/rf_api"
)

var BaseRFUrl = "https://app.redforester.com"
var ClientInstance = NewClient(BaseRFUrl, "", "")

func NewClient(rfBaseUrl, username, passwordHash string) *rf_api.Client {
	httpClient, err := newHTTPClient(username, passwordHash)
	if err != nil {
		fmt.Println("ERROR: http client not loaded:", err)
		return nil
	}

	return rf_api.New(rfBaseUrl, httpClient)
}
