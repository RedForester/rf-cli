package factory

import (
	"encoding/base64"
	"github.com/deissh/rf-cli/internal/config"
	"github.com/deissh/rf-cli/pkg/http_client"
	"net/http"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func newHTTPClient(cfg *config.Config) (*http.Client, error) {
	opts := []http_client.ClientOption{
		http_client.AddHeader("Authorization", basicAuth(cfg.Client.Username, cfg.Client.PasswordHash)),
		http_client.AddHeader("Content-Type", "application/json"),
	}

	return http_client.NewHTTPClient(opts...), nil
}
