package sensu

import (
	"log"
	"strings"
)

func PostRequest(check string, subscriber string) string {
	body := `{"check":"` + check + `","subscribers":["` + subscriber + `"]}`
	payload := strings.NewReader(body)

	_, status := postAPI("/request", payload)
	if status != 202 {
		log.Fatal(httpStatus(status))
	}

	return httpStatus(status) + "\n"
}
