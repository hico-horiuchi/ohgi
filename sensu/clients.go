package sensu

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ClientStruct struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Timestamp     int64    `json:"timestamp"`
	Version       string   `json:"version"`
}

// Returns a list of clients.
//
//   limit:  The number of clients to return.
//   offset: The number of clients to offset before returning items.
//
func (api API) GetClients(limit int, offset int) ([]ClientStruct, error) {
	var clients []ClientStruct

	response, err := api.get(fmt.Sprintf("/clients?limit=%d&offset=%d", limit, offset))
	if err != nil {
		return clients, err
	} else if response.StatusCode != 200 {
		return clients, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &clients)
	if err != nil {
		return clients, err
	}

	return clients, nil
}

// Returns a client.
func (api API) GetClientsClient(name string) (ClientStruct, error) {
	var client ClientStruct

	response, err := api.get("/clients/" + name)
	if err != nil {
		return client, err
	} else if response.StatusCode != 200 {
		return client, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &client)
	if err != nil {
		return client, err
	}

	return client, nil
}

// Create or update client data (e.g. Sensu JIT clients).
func (api API) PostClients(name string, address string, subscriptions []string) error {
	client := ClientStruct{
		Name:          name,
		Address:       address,
		Subscriptions: subscriptions,
	}

	body, err := json.Marshal(client)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(body))

	response, err := api.post("/clients", payload)
	if err != nil {
		return err
	} else if response.StatusCode != 201 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}

// Removes a client, resolving its current events.
func (api API) DeleteClientsClient(name string) error {
	response, err := api.delete("/client/" + name)
	if err != nil {
		return err
	} else if response.StatusCode != 202 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}
