package sensu

import (
	"fmt"
	"log"
)

func GetHealth(consumers int, messages int) string {
	_, status := getAPI(fmt.Sprintf("/health?consumers=%d&messages=%d", consumers, messages))
	if status != 204 {
		log.Fatal(httpStatus(status))
	}

	return httpStatus(status) + "\n"
}
