package ohgi

import "github.com/hico-horiuchi/ohgi/sensu"

func PostRequest(api *sensu.API, check string, subscribers []string) string {
	err := api.PostRequest(check, subscribers)
	checkError(err)

	return "Accepted\n"
}
