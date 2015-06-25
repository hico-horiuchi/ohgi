package ohgi

import "../sensu"

func PostResolve(api *sensu.API, client string, check string) string {
	err := api.PostResolve(client, check)
	checkError(err)

	return "Accepted\n"
}
