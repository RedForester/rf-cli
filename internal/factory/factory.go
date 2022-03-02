package factory

import (
	"fmt"
	"github.com/deissh/rf-cli/pkg/rf_api"
)

func NewClient(username, passwordHash string) *rf_api.Client {
	httpClient, err := newHTTPClient(username, passwordHash)
	if err != nil {
		fmt.Println("ERROR: http client not loaded:", err)
		return nil
	}

	return rf_api.New(httpClient)
}
