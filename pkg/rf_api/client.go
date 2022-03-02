package rf_api

import (
	"net/http"
)

var RfBaseurl = "https://beta.app.redforester.com"

type Options struct {
	BaseURL string
}

type Client struct {
	Ext  *ExtensionsApi
	User *UserApi
}

func New(client *http.Client) *Client {
	svc := Service{
		Client: client,
		Options: &Options{
			BaseURL: RfBaseurl,
		},
	}

	ext := &ExtensionsApi{svc}
	user := &UserApi{svc}

	return &Client{ext, user}
}
