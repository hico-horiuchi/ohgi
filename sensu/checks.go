package sensu

import (
	"encoding/json"
	"errors"
)

type CheckStruct struct {
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	Subscribers []string `json:"subscribers"`
	Interval    int      `json:"interval"`
	Handlers    []string `json:"handlers"`
	Issued      int64    `json:"issued"`
	Executed    int64    `json:"executed"`
	Output      string   `json:"output"`
	Status      int      `json:"status"`
	Duration    float64  `json:"duration"`
	History     []string `json:"history"`
}

// Returns the list of checks.
func (api API) GetChecks() ([]CheckStruct, error) {
	var checks []CheckStruct

	response, err := api.get("/checks")
	if err != nil {
		return checks, err
	} else if response.StatusCode != 200 {
		return checks, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &checks)
	if err != nil {
		return checks, err
	}

	return checks, nil
}

// Returns a check.
func (api API) GetChecksCheck(name string) (CheckStruct, error) {
	var check CheckStruct

	response, err := api.get("/checks/" + name)
	if err != nil {
		return check, err
	} else if response.StatusCode != 200 {
		return check, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &check)
	if err != nil {
		return check, err
	}

	return check, nil
}
