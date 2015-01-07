package sensu

import (
	"fmt"
	"os"
)

func GetHealth(consumers int, messages int) string {
	_, status := getAPI(fmt.Sprintf("/health?consumers=%d&messages=%d", consumers, messages))
	if status != 204 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	return httpStatus(status) + "\n"
}
