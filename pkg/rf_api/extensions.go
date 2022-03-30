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

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + " " + string(bodyBytes))
	}

	data := &[]rf.Extension{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e ExtensionsApi) GetOwned() (*[]rf.Extension, error) {
	resp, err := e.Client.Get(e.Options.BaseURL + "/api/extensions/owned")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + " " + string(bodyBytes))
	}

	data := &[]rf.Extension{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e ExtensionsApi) Get(id string) (*rf.Extension, error) {
	resp, err := e.Client.Get(e.Options.BaseURL + "/api/extensions/" + id)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + " " + string(bodyBytes))
	}

	data, err := rf.UnmarshalExtension(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (e ExtensionsApi) Create(ext *rf.Extension) (*rf.Extension, error) {
	var payloadBuf bytes.Buffer

	err := json.NewEncoder(&payloadBuf).Encode(ext)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", e.Options.BaseURL+"/api/extensions", &payloadBuf)
	resp, err := e.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + " " + string(bodyBytes))
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

	req, _ := http.NewRequest("POST", e.Options.BaseURL+"/api/extensions/"+ext.ID, &payloadBuf)
	resp, err := e.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + " " + string(bodyBytes))
	}

	data, err := rf.UnmarshalExtension(bodyBytes)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
