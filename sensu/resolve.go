package sensu

import (
	"log"
	"strings"
)

func PostResolve(client string, check string) string {
	body := `{"client":"` + client + `","check":"` + check + `"}`
	payload := strings.NewReader(body)

	_, status := postAPI("/resolve", payload)
	if status != 202 {
		log.Fatal(httpStatus(status))
	}

	return httpStatus(status) + "\n"
}
