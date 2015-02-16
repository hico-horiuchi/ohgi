package sensu

import (
	"strings"
)

func PostResolve(client string, check string) string {
	body := `{"client":"` + client + `","check":"` + check + `"}`
	payload := strings.NewReader(body)
	_, status := postAPI("/resolve", payload)

	return httpStatus(status)
}
