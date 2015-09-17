package sensu

import (
	"encoding/json"
	"errors"
)

type InfoStruct struct {
	Sensu struct {
		Version string `json:"version"`
	} `json:"sensu"`
	Transport struct {
		Keepalives struct {
			Messages  int `json:"messages"`
			Consumers int `json:"consumers"`
		} `json:"keepalives"`
		Results struct {
			Messages  int `json:"messages"`
			Consumers int `json:"consumers"`
		} `json:"results"`
		Connected bool `json:"connected"`
	}
	Redis struct {
		Connected bool `json:"connected"`
	} `json:"redis"`
}

// Returns information on the API.
func (api API) GetInfo() (InfoStruct, error) {
	var info InfoStruct

	response, err := api.get("/info")
	if err != nil {
		return info, err
	} else if response.StatusCode != 200 {
		return info, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &info)
	if err != nil {
		return info, err
	}

	return info, nil
}
