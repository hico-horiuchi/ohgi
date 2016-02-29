package sensu

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type StashStruct struct {
	Expire int64  `json:"expire"`
	Path   string `json:"path"`
}

// Returns a list of stashes.
//
//   limit:  The number of stashes to return.
//   offset: The number of stashes to offset before returning items.
//
func (api API) GetStashes(stashes interface{}, limit int, offset int) error {
	response, err := api.get(fmt.Sprintf("/stashes?limit=%d&offset=%d", limit, offset))
	if err != nil {
		return err
	} else if response.StatusCode != 200 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), stashes)
	if err != nil {
		return err
	}

	return nil
}

// Create a stash.
func (api API) PostStashes(stash interface{}) error {
	body, err := json.Marshal(stash)
	if err != nil {
		return err
	}
	payload := strings.NewReader(strings.Replace(string(body), `"expire":-1,`, "", -1))

	response, err := api.post("/stashes", payload)
	if err != nil {
		return err
	} else if response.StatusCode != 201 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}

// Create a stash.
func (api API) PostStashesPath(path string, content interface{}) error {
	body, err := json.Marshal(content)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(body))

	response, err := api.post("/stashes/"+path, payload)
	if err != nil {
		return err
	} else if response.StatusCode != 201 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}

// Get a stash.
func (api API) GetStashesPath(path string, content interface{}) error {
	response, err := api.get("/stashes/" + path)
	if err != nil {
		return err
	} else if response.StatusCode != 200 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), content)
	if err != nil {
		return err
	}

	return nil
}

// Delete a stash.
func (api API) DeleteStashesPath(path string) error {
	response, err := api.delete("/stashes/" + path)
	if err != nil {
		return err
	} else if response.StatusCode != 204 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}
