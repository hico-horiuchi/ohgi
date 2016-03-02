package ohgi

import "github.com/hico-horiuchi/ohgibone/sensu"

func PostResolve(api *sensu.API, client string, check string) string {
	err := api.PostResolve(client, check)
	checkError(err)

	return "Accepted\n"
}
