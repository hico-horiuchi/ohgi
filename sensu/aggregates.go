package sensu

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AggregateList struct {
	Check  string  `json:"check"`
	Issued []int64 `json:"issued"`
}

type AggregateStruct struct {
	Ok       int `json:"ok"`
	Warning  int `json:"warning"`
	Critical int `json:"critical"`
	Unknown  int `json:"unknown"`
	Total    int `json:"total"`
	Results  []struct {
		Client string `json:"client"`
		Output string `json:"output"`
		Status int    `json:"status"`
	}
}

// Returns the list of aggregates.
func (api API) GetAggregates(limit int, offset int) ([]AggregateList, error) {
	var aggregates []AggregateList

	response, err := api.get(fmt.Sprintf("/aggregates?limit=%d&offset=%d", limit, offset))
	if err != nil {
		return aggregates, err
	} else if response.StatusCode != 200 {
		return aggregates, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &aggregates)
	if err != nil {
		return aggregates, err
	}

	return aggregates, nil
}

// Returns the list of aggregates for a given check.
func (api API) GetAggregatesCheck(check string, age int) ([]int64, error) {
	var issues []int64

	response, err := api.get(fmt.Sprintf("/aggregates/%s?age=%d", check, age))
	if err != nil {
		return issues, err
	} else if response.StatusCode != 200 {
		return issues, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}

// Deletes all aggregates for a check.
func (api API) DeleteAggregatesCheck(check string) error {
	response, err := api.delete("/aggregates/" + check)
	if err != nil {
		return err
	} else if response.StatusCode != 204 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}

// Returns an aggregate.
func (api API) GetAggregatesCheckIssued(check string, issued int64, results bool) (AggregateStruct, error) {
	var aggregate AggregateStruct

	response, err := api.get(fmt.Sprintf("/aggregates/%s/%d?results=%t", check, issued, results))
	if err != nil {
		return aggregate, err
	} else if response.StatusCode != 200 {
		return aggregate, errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	err = json.Unmarshal([]byte(response.Body), &aggregate)
	if err != nil {
		return aggregate, err
	}

	return aggregate, nil
}
