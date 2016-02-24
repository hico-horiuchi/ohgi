package ohgi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hico-horiuchi/ohgi/sensu"
)

func GetClients(api *sensu.API, limit int, offset int) string {
	clients, err := api.GetClients(limit, offset)
	checkError(err)

	if len(clients) == 0 {
		return "No clients\n"
	}

	table := newUitable()
	table.AddRow(bold("NAME"), bold("ADDRESS"), bold("TIMESTAMP"))
	for _, client := range clients {
		table.AddRow(
			client.Name,
			client.Address,
			client.Timestamp,
		)
	}

	return fmt.Sprintln(table)
}

func GetClientsWildcard(api *sensu.API, pattern string) string {
	var match []string
	var matches []int
	var client sensu.ClientStruct

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
		return "No clients that matches " + pattern + "\n"
	}

	table := newUitable()
	table.AddRow(bold("NAME"), bold("ADDRESS"), bold("TIMESTAMP"))
	for _, i := range matches {
		client = clients[i]
		table.AddRow(
			client.Name,
			client.Address,
			client.Timestamp,
		)
	}

	return fmt.Sprintln(table)
}

func GetClientsClient(api *sensu.API, name string) string {
	client, err := api.GetClientsClient(name)
	checkError(err)

	table := newUitable(true)
	table.AddRow(bold("NAME:"), client.Name)
	table.AddRow(bold("ADDRESS:"), client.Address)
	table.AddRow(bold("SUBSCRIPTIONS:"), strings.Join(client.Subscriptions, ", "))
	table.AddRow(bold("TIMESTAMP:"), utoa(client.Timestamp))
	table.AddRow(bold("VERSION:"), client.Version)

	return fmt.Sprintln(table)
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
