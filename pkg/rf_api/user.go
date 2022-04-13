package rf_api

import (
	"errors"
	"github.com/deissh/rf-cli/pkg/rf"
	"io"
	"net/http"
)

type UserApi struct {
	Service
}

func (u UserApi) GetMe() (*rf.User, error) {
	resp, err := u.Client.Get(u.Options.BaseURL + "/api/user")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data, err := rf.UnmarshalUser(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
