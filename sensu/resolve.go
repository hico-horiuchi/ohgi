package sensu

import (
	"fmt"
	"os"
	"strings"
)

func PostResolve(client string, check string) string {
	body := `{"client":"` + client + `","check":"` + check + `"}`
	payload := strings.NewReader(body)

	_, status := postAPI("/resolve", payload)
	if status != 202 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	return httpStatus(status)
}
