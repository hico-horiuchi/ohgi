package sensu

import (
	"encoding/json"
	"errors"
	"strings"
)

type RequestStruct struct {
	Check       string   `json:"check"`
	Subscribers []string `json:"subscribers"`
}

// Issues a check execution request.
func (api API) PostRequest(check string, subscribers []string) error {
	request := RequestStruct{
		Check:       check,
		Subscribers: subscribers,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(body))

	response, err := api.post("/request", payload)
	if err != nil {
		return err
	} else if response.StatusCode != 202 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}
