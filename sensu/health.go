package sensu

import (
	"errors"
	"fmt"
)

// Returns health information on transport & Redis connections.
func (api API) GetHealth(consumers int, messages int) error {
	response, err := api.get(fmt.Sprintf("/health?consumers=%d&messages=%d", consumers, messages))
	if err != nil {
		return err
	} else if response.StatusCode != 204 {
		return errors.New("sensu: " + StatusCodeToString(response.StatusCode))
	}

	return nil
}
