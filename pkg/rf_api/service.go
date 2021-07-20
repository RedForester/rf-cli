package rf_api

import "net/http"

type Service struct {
	Client  *http.Client
	Options *Options
}
