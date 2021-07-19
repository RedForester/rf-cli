package http_client

import "net/http"

type ClientOption = func(http.RoundTripper) http.RoundTripper

func AddHeader(name, value string) ClientOption {
	return func(tr http.RoundTripper) http.RoundTripper {
		return &funcTripper{roundTrip: func(req *http.Request) (*http.Response, error) {
			if req.Header.Get(name) == "" {
				req.Header.Add(name, value)
			}
			return tr.RoundTrip(req)
		}}
	}
}

type funcTripper struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (tr funcTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return tr.roundTrip(req)
}
