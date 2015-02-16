package sensu

import (
	"fmt"
	"os"
	"strings"
)

func PostRequest(check string, subscriber string) string {
	body := `{"check":"` + check + `","subscribers":["` + subscriber + `"]}`
	payload := strings.NewReader(body)

	_, status := postAPI("/request", payload)
	if status != 202 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	return httpStatus(status)
}
