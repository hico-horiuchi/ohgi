package sensu

import (
	"encoding/json"
	"errors"
)

type EventStruct struct {
	ID          string       `json:"id"`
	Client      ClientStruct `json:"client"`
	Check       CheckStruct  `json:"check"`
	Occurrences int          `json:"occurrences"`
	Action      string       `json:"action"`
}

// Returns the list of current events.
func (api API) GetEvents() ([]EventStruct, error) {
	var events []EventStruct

	response, err := api.get("/events")
	if err != nil {
		return events, err
	} else if response.StatusCode != 200 {
		return events, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &events)
	if err != nil {
		return events, err
	}

	return events, nil
}

// Returns the list of current events for a given client.
func (api API) GetEventsClient(client string) ([]EventStruct, error) {
	var events []EventStruct

	response, err := api.get("/events/" + client)
	if err != nil {
		return events, err
	} else if response.StatusCode != 200 {
		return events, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &events)
	if err != nil {
		return events, err
	}

	return events, nil
}

// Returns an event for a given client & check name.
func (api API) GetEventsClientCheck(client string, check string) (EventStruct, error) {
	var event EventStruct

	response, err := api.get("/event/" + client + "/" + check)
	if err != nil {
		return event, err
	} else if response.StatusCode != 200 {
		return event, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &event)
	if err != nil {
		return event, err
	}

	return event, nil
}

// Resolves an event for a given check on a given client.
func (api API) DeleteEventsClientCheck(client string, check string) error {
	response, err := api.delete("/event/" + client + "/" + check)
	if err != nil {
		return err
	} else if response.StatusCode != 202 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}
