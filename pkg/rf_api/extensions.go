package rf_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/deissh/rf-cli/pkg/rf"
	"io/ioutil"
	"net/http"
)

type ExtensionsApi struct {
	Service
}

func (e ExtensionsApi) GetAll() (*[]rf.Extension, error) {
	resp, err := e.Client.Get(e.Options.BaseURL + "/api/extensions")
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

	data := &[]rf.Extension{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e ExtensionsApi) GetOwned() (*[]rf.Extension, error) {
	resp, err := e.Client.Get(e.Options.BaseURL + "/api/extension/owned")
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

	data := &[]rf.Extension{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e ExtensionsApi) Get(id string) (*rf.Extension, error) {
	resp, err := e.Client.Get(e.Options.BaseURL + "/api/extension/" + id)
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

	data, err := rf.UnmarshalExtension(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (e ExtensionsApi) Update(ext *rf.Extension) (*rf.Extension, error) {
	var payloadBuf bytes.Buffer

	err := json.NewEncoder(&payloadBuf).Encode(ext)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("PATCH", e.Options.BaseURL+"/api/extension/"+ext.ID, &payloadBuf)
	resp, err := e.Client.Do(req)
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

	data, err := rf.UnmarshalExtension(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
