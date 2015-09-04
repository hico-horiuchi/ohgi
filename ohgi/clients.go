package ohgi

import (
	"regexp"
	"strings"

	"github.com/hico-horiuchi/ohgi/sensu"
)

func GetClients(api *sensu.API, limit int, offset int) string {
	var line string

	clients, err := api.GetClients(limit, offset)
	checkError(err)

	if len(clients) == 0 {
		return "No clients\n"
	}

	print := []byte(bold("NAME                                    ADDRESS                                 TIMESTAMP\n"))
	for _, client := range clients {
		line = fillSpace(client.Name, 40) + fillSpace(client.Address, 40) + utoa(client.Timestamp) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func GetClientsWildcard(api *sensu.API, pattern string) string {
	var match []string
	var matches []int
	var line string

	clients, err := api.GetClients(-1, -1)
	checkError(err)

	re := regexp.MustCompile("^" + strings.Replace(pattern, "*", ".*", -1) + "$")
	for i, client := range clients {
		match = re.FindStringSubmatch(client.Name)
		if len(match) > 0 {
			matches = append(matches, i)
		}
	}

	if len(matches) == 0 {
		return "No clients that match " + pattern + "\n"
	}

	print := []byte(bold("NAME                                    ADDRESS                                 TIMESTAMP\n"))
	for _, i := range matches {
		client := clients[i]
		line = fillSpace(client.Name, 40) + fillSpace(client.Address, 40) + utoa(client.Timestamp) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func GetClientsClient(api *sensu.API, name string) string {
	var print []byte

	client, err := api.GetClientsClient(name)
	checkError(err)

	print = append(print, (bold("NAME           ") + client.Name + "\n")...)
	print = append(print, (bold("ADDRESS        ") + client.Address + "\n")...)
	print = append(print, (bold("SUBSCRIPTIONS  ") + strings.Join(client.Subscriptions, ", ") + "\n")...)
	print = append(print, (bold("TIMESTAMP      ") + utoa(client.Timestamp) + "\n")...)
	print = append(print, (bold("VERSION        ") + client.Version + "\n")...)

	return string(print)
}

func PostClients(api *sensu.API, name string, address string, subscriptions []string) string {
	err := api.PostClients(name, address, subscriptions)
	checkError(err)

	return "Created\n"
}

func DeleteClientsClient(api *sensu.API, name string) string {
	err := api.DeleteClientsClient(name)
	checkError(err)

	return "Accepted\n"
}
