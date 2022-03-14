package rf_api

import (
	"net/http"
)

type Options struct {
	BaseURL string
}

type Client struct {
	Ext  *ExtensionsApi
	User *UserApi
}

func New(rfBaseUrl string, client *http.Client) *Client {
	svc := Service{
		Client: client,
		Options: &Options{
			BaseURL: rfBaseUrl,
		},
	}

	ext := &ExtensionsApi{svc}
	user := &UserApi{svc}

	return &Client{ext, user}
}
