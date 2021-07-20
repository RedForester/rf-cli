package rf_api

import "net/http"

const RfBaseurl = "https://app.redforester.com"

type Options struct {
	BaseURL string
}

type Client struct {
	Ext *ExtensionsApi
}

func NewOptions(baseUrl string) *Options {
	if baseUrl == "" {
		baseUrl = RfBaseurl
	}

	return &Options{
		BaseURL: baseUrl,
	}
}

func New(client *http.Client, opt *Options) *Client {
	svc := Service{
		Client:  client,
		Options: opt,
	}

	ext := &ExtensionsApi{svc}

	return &Client{ext}
}
