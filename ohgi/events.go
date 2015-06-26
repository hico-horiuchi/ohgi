package ohgi

import (
	"strconv"
	"strings"

	"../sensu"
)

func GetEvents(api *sensu.API) string {
	var line string

	events, err := api.GetEvents()
	checkError(err)

	if len(events) == 0 {
		return "No current events\n"
	}

	print := []byte(bold("  CLIENT                                  CHECK                         #         EXECUTED\n"))
	for _, event := range events {
		line = indicateStatus(event.Check.Status) + fillSpace(event.Client.Name, 40) + fillSpace(event.Check.Name, 30) + fillSpace(strconv.Itoa(event.Occurrences), 10) + utoa(event.Check.Executed) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func GetEventsClient(api *sensu.API, client string) string {
	var line, output string

	events, err := api.GetEventsClient(client)
	checkError(err)

	if len(events) == 0 {
		return "No current events for " + client + "\n"
	}

	print := []byte(bold("  CHECK                         OUTPUT                                            EXECUTED\n"))
	for _, event := range events {
		output = strings.Replace(event.Check.Output, "\n", " ", -1)
		line = indicateStatus(event.Check.Status) + fillSpace(event.Check.Name, 30) + fillSpace(output, 50) + utoa(event.Check.Executed) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func GetEventsClientCheck(api *sensu.API, client string, check string) string {
	var print []byte

	event, err := api.GetEventsClientCheck(client, check)
	checkError(err)

	print = append(print, (bold("CLIENT         ") + event.Client.Name + "\n")...)
	print = append(print, (bold("ADDRESS        ") + event.Client.Address + "\n")...)
	print = append(print, (bold("SUBSCRIPTIONS  ") + strings.Join(event.Client.Subscriptions, ", ") + "\n")...)
	print = append(print, (bold("TIMESTAMP      ") + utoa(event.Client.Timestamp) + "\n")...)
	print = append(print, (bold("CHECK          ") + event.Check.Name + "\n")...)
	print = append(print, (bold("COMMAND        ") + event.Check.Command + "\n")...)
	print = append(print, (bold("SUBSCRIBERS    ") + strings.Join(event.Check.Subscribers, ", ") + "\n")...)
	print = append(print, (bold("INTERVAL       ") + strconv.Itoa(event.Check.Interval) + "\n")...)
	print = append(print, (bold("HANDLERS       ") + strings.Join(event.Check.Handlers, ", ") + "\n")...)
	print = append(print, (bold("ISSUED         ") + utoa(event.Check.Issued) + "\n")...)
	print = append(print, (bold("EXECUTED       ") + utoa(event.Check.Executed) + "\n")...)
	print = append(print, (bold("OUTPUT         ") + strings.Replace(event.Check.Output, "\n", " ", -1) + "\n")...)
	print = append(print, (bold("STATUS         ") + paintStatus(event.Check.Status) + "\n")...)
	print = append(print, (bold("DURATION       ") + strconv.FormatFloat(event.Check.Duration, 'f', 3, 64) + "\n")...)
	print = append(print, (bold("HISTORY        ") + paintHistory(strings.Join(event.Check.History, ", ")) + "\n")...)
	print = append(print, (bold("OCCURRENCES    ") + strconv.Itoa(event.Occurrences) + "\n")...)
	print = append(print, (bold("ACTION         ") + event.Action + "\n")...)

	return string(print)
}

func DeleteEventsClientCheck(api *sensu.API, client string, check string) string {
	err := api.DeleteEventsClientCheck(client, check)
	checkError(err)

	return "Accepted\n"
}
