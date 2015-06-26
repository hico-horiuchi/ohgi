package sensu

import (
	"encoding/json"
	"errors"
)

type historyStruct struct {
	Check         string `json:"check"`
	History       []int  `json:"history"`
	LastExecution int64  `json:"last_execution"`
	LastStatus    int    `json:"last_status"`
}

// Returns the history for a client.
func (api API) GetClientsHistory(client string) ([]historyStruct, error) {
	var histories []historyStruct

	response, err := api.get("/clients/" + client + "/history")
	if err != nil {
		return histories, err
	} else if response.StatusCode != 200 {
		return histories, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &histories)
	if err != nil {
		return histories, err
	}

	return histories, nil
}
