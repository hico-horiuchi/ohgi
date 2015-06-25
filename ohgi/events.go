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

	result := []byte(bold("  CLIENT                                  CHECK                         #         EXECUTED\n"))
	for _, event := range events {
		line = indicateStatus(event.Check.Status) + fillSpace(event.Client.Name, 40) + fillSpace(event.Check.Name, 30) + fillSpace(strconv.Itoa(event.Occurrences), 10) + utoa(event.Check.Executed) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetEventsClient(api *sensu.API, client string) string {
	var line, output string

	events, err := api.GetEvents()
	checkError(err)

	if len(events) == 0 {
		return "No current events for " + client + "\n"
	}

	result := []byte(bold("  CHECK                         OUTPUT                                            EXECUTED\n"))
	for _, event := range events {
		output = strings.Replace(event.Check.Output, "\n", " ", -1)
		line = indicateStatus(event.Check.Status) + fillSpace(event.Check.Name, 30) + fillSpace(output, 50) + utoa(event.Check.Executed) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetEventsClientCheck(api *sensu.API, client string, check string) string {
	var result []byte

	event, err := api.GetEventsClientCheck(client, check)
	checkError(err)

	result = append(result, (bold("CLIENT         ") + event.Client.Name + "\n")...)
	result = append(result, (bold("ADDRESS        ") + event.Client.Address + "\n")...)
	result = append(result, (bold("SUBSCRIPTIONS  ") + strings.Join(event.Client.Subscriptions, ", ") + "\n")...)
	result = append(result, (bold("TIMESTAMP      ") + utoa(event.Client.Timestamp) + "\n")...)
	result = append(result, (bold("CHECK          ") + event.Check.Name + "\n")...)
	result = append(result, (bold("COMMAND        ") + event.Check.Command + "\n")...)
	result = append(result, (bold("SUBSCRIBERS    ") + strings.Join(event.Check.Subscribers, ", ") + "\n")...)
	result = append(result, (bold("INTERVAL       ") + strconv.Itoa(event.Check.Interval) + "\n")...)
	result = append(result, (bold("HANDLERS       ") + strings.Join(event.Check.Handlers, ", ") + "\n")...)
	result = append(result, (bold("ISSUED         ") + utoa(event.Check.Issued) + "\n")...)
	result = append(result, (bold("EXECUTED       ") + utoa(event.Check.Executed) + "\n")...)
	result = append(result, (bold("OUTPUT         ") + strings.Replace(event.Check.Output, "\n", " ", -1) + "\n")...)
	result = append(result, (bold("STATUS         ") + paintStatus(event.Check.Status) + "\n")...)
	result = append(result, (bold("DURATION       ") + strconv.FormatFloat(event.Check.Duration, 'f', 3, 64) + "\n")...)
	result = append(result, (bold("HISTORY        ") + paintHistory(strings.Join(event.Check.History, ", ")) + "\n")...)
	result = append(result, (bold("OCCURRENCES    ") + strconv.Itoa(event.Occurrences) + "\n")...)
	result = append(result, (bold("ACTION         ") + event.Action + "\n")...)

	return string(result)
}

func DeleteEventsClientCheck(api *sensu.API, client string, check string) string {
	err := api.DeleteEventsClientCheck(client, check)
	checkError(err)

	return "Accepted\n"
}
