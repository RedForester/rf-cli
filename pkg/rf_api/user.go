package rf_api

import (
	"encoding/json"
	"errors"
	"github.com/deissh/rf-cli/pkg/rf"
	"io/ioutil"
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

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := &rf.User{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
