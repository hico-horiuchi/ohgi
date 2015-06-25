package ohgi

import "../sensu"

func GetHealth(api *sensu.API, consumers int, messages int) string {
	err := api.GetHealth(consumers, messages)
	checkError(err)

	return "No Content\n"
}
