package rf_api

import (
	"encoding/json"
	"github.com/deissh/rf-cli/pkg/extension"
	"io/ioutil"
	"log"
	"net/http"
)

type ExtensionsApi struct {
	Service
}

func (e ExtensionsApi) GetAll() (*[]extension.Extension, error) {
	resp, err := e.Client.Get(e.Options.BaseURL + "/api/extensions")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := &[]extension.Extension{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e ExtensionsApi) GetOwned() (*[]extension.Extension, error) {
	resp, err := e.Client.Get(e.Options.BaseURL + "/api/extensions/owned")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := &[]extension.Extension{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e ExtensionsApi) Get(id string) (*extension.Extension, error) {
	resp, err := e.Client.Get(e.Options.BaseURL + "/api/extensions/" + id)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := &extension.Extension{}
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
