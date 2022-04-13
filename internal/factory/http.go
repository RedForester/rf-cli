package factory

import (
	"encoding/base64"
	"github.com/deissh/rf-cli/pkg/http_client"
	"net/http"
)

func basicAuth(username, passwordHash string) string {
	auth := username + ":" + passwordHash
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func newHTTPClient(username, passwordHash string) (*http.Client, error) {
	opts := []http_client.ClientOption{
		http_client.AddHeader("Content-Type", "application/json"),
	}
	if username != "" && passwordHash != "" {
		opts = append(opts, http_client.AddHeader("Authorization", basicAuth(username, passwordHash)))
	}

	return http_client.NewHTTPClient(opts...), nil
}
