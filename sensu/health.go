package sensu

import (
	"errors"
	"fmt"
)

// Returns health information on transport & Redis connections.
//
//   consumers: The minimum number of transport consumers to be considered healthy.
//   messages:  The maximum ammount of transport queued messages to be considered healthy.
//
func (api API) GetHealth(consumers int, messages int) error {
	response, err := api.get(fmt.Sprintf("/health?consumers=%d&messages=%d", consumers, messages))
	if err != nil {
		return err
	} else if response.StatusCode != 204 {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}
