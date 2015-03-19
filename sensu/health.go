package sensu

import (
	"fmt"
)

func GetHealth(consumers int, messages int) string {
	_, status := getAPI(fmt.Sprintf("/health?consumers=%d&messages=%d", consumers, messages))
	checkStatus(status)

	return httpStatus(status) + "\n"
}
