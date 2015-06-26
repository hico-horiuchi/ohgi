package sensu

import (
	"encoding/json"
	"errors"
)

type ResultStruct struct {
	Client string      `json:"client"`
	Check  CheckStruct `json:"check"`
}

// Returns a list of current check results for all clients.
func (api API) GetResults() ([]ResultStruct, error) {
	var results []ResultStruct

	response, err := api.get("/results")
	if err != nil {
		return results, err
	} else if response.StatusCode != 200 {
		return results, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

// Returns a list of current check results for a given client.
func (api API) GetResultsClient(client string) ([]ResultStruct, error) {
	var results []ResultStruct

	response, err := api.get("/results/" + client)
	if err != nil {
		return results, err
	} else if response.StatusCode != 200 {
		return results, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

// Returns a check result for a given client & check name.
func (api API) GetResultsClientCheck(client string, check string) (ResultStruct, error) {
	var result ResultStruct

	response, err := api.get("/result/" + client + "/" + check)
	if err != nil {
		return result, err
	} else if response.StatusCode != 200 {
		return result, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
