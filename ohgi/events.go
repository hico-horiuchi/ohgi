package ohgi

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hico-horiuchi/ohgibone/sensu"
)

func GetEvents(api *sensu.API) string {
	events, err := api.GetEvents()
	checkError(err)

	if len(events) == 0 {
		return "No current events\n"
	}

	table := newUitable()
	table.AddRow("", bold("CLIENT"), bold("CHECK"), bold("#"), bold("EXECUTED"))
	for _, event := range events {
		table.AddRow(
			indicateStatus(event.Check.Status),
			event.Client.Name,
			event.Check.Name,
			strconv.Itoa(event.Occurrences),
			utoa(event.Check.Executed),
		)
	}

	return fmt.Sprintln(table)
}

func GetEventsWildcard(api *sensu.API, pattern string) string {
	var match []string
	var matches []int
	var event sensu.EventStruct

	events, err := api.GetEvents()
	checkError(err)

	re := regexp.MustCompile("^" + strings.Replace(pattern, "*", ".*", -1) + "$")
	for i, event := range events {
		match = re.FindStringSubmatch(event.Client.Name)
		if len(match) > 0 {
			matches = append(matches, i)
		}
	}

	if len(matches) == 0 {
		return "No current events that matches " + pattern + "\n"
	}

	table := newUitable()
	table.AddRow("", bold("CLIENT"), bold("CHECK"), bold("#"), bold("EXECUTED"))
	for _, i := range matches {
		event = events[i]
		table.AddRow(
			indicateStatus(event.Check.Status),
			event.Client.Name,
			event.Check.Name,
			strconv.Itoa(event.Occurrences),
			utoa(event.Check.Executed),
		)
	}

	return fmt.Sprintln(table)
}

func GetEventsClient(api *sensu.API, client string) string {
	events, err := api.GetEventsClient(client)
	checkError(err)

	if len(events) == 0 {
		return "No current events for " + client + "\n"
	}

	table := newUitable()
	table.AddRow("", bold("CHECK"), bold("OUTPUT"), bold("EXECUTED"))
	for _, event := range events {
		table.AddRow(
			indicateStatus(event.Check.Status),
			event.Check.Name,
			formatOutput(event.Check.Output),
			utoa(event.Check.Executed),
		)
	}

	return fmt.Sprintln(table)
}

func GetEventsClientCheck(api *sensu.API, client string, check string) string {
	event, err := api.GetEventsClientCheck(client, check)
	checkError(err)

	table := newUitable(true)
	table.AddRow(bold("CLIENT:"), event.Client.Name)
	table.AddRow(bold("ADDRESS:"), event.Client.Address)
	table.AddRow(bold("SUBSCRIPTIONS:"), strings.Join(event.Client.Subscriptions, ", "))
	table.AddRow(bold("TIMESTAMP:"), utoa(event.Client.Timestamp))
	table.AddRow(bold("CHECK:"), event.Check.Name)
	table.AddRow(bold("COMMAND:"), event.Check.Command)
	table.AddRow(bold("SUBSCRIBERS:"), strings.Join(event.Check.Subscribers, ", "))
	table.AddRow(bold("INTERVAL:"), strconv.Itoa(event.Check.Interval))
	table.AddRow(bold("HANDLERS:"), strings.Join(event.Check.Handlers, ", "))
	table.AddRow(bold("ISSUED:"), utoa(event.Check.Issued))
	table.AddRow(bold("EXECUTED:"), utoa(event.Check.Executed))
	table.AddRow(bold("OUTPUT:"), formatOutput(event.Check.Output))
	table.AddRow(bold("STATUS:"), colorStatus(event.Check.Status))
	table.AddRow(bold("DURATION:"), strconv.FormatFloat(event.Check.Duration, 'f', 3, 64))
	table.AddRow(bold("HISTORY:"), colorHistory(strings.Join(event.Check.History, ", ")))
	table.AddRow(bold("OCCURRENCES:"), strconv.Itoa(event.Occurrences))
	table.AddRow(bold("ACTION:"), event.Action)

	return fmt.Sprintln(table)
}

func DeleteEventsClientCheck(api *sensu.API, client string, check string) string {
	err := api.DeleteEventsClientCheck(client, check)
	checkError(err)

	return "Accepted\n"
}
