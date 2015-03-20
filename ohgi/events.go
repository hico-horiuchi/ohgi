package ohgi

import (
	"encoding/json"
	"strconv"
	"strings"
)

type eventStruct struct {
	Id          string
	Client      clientStruct
	Check       checkStruct
	Occurrences int
	Action      string
}

func GetEvents() string {
	var events []eventStruct
	var result []byte

	contents, status := getAPI("/events")
	checkStatus(status)

	json.Unmarshal(contents, &events)
	if len(events) == 0 {
		return "No current events\n"
	}

	result = append(result, bold("  CLIENT                                  CHECK                         #         TIMESTAMP\n")...)
	for i := range events {
		e := events[i]
		occurrences := strconv.Itoa(e.Occurrences)
		line := statusBg(e.Check.Status) + fillSpace(e.Client.Name, 40) + fillSpace(e.Check.Name, 30) + fillSpace(occurrences, 10) + utoa(e.Client.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetEventsClient(client string) string {
	var events []eventStruct
	var result []byte

	contents, status := getAPI("/events/" + client)
	checkStatus(status)

	json.Unmarshal(contents, &events)
	if len(events) == 0 {
		return "No current events for " + client + "\n"
	}

	result = append(result, bold("  CHECK                         OUTPUT                                            TIMESTAMP\n")...)
	for i := range events {
		e := events[i]
		output := strings.Replace(e.Check.Output, "\n", " ", -1)
		line := statusBg(e.Check.Status) + fillSpace(e.Check.Name, 30) + fillSpace(output, 50) + utoa(e.Client.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetEventsClientCheck(client string, check string) string {
	var e eventStruct
	var result []byte

	contents, status := getAPI("/events/" + client + "/" + check)
	checkStatus(status)

	json.Unmarshal(contents, &e)

	result = append(result, (bold("CLIENT         ") + e.Client.Name + "\n")...)
	result = append(result, (bold("ADDRESS        ") + e.Client.Address + "\n")...)
	result = append(result, (bold("SUBSCRIPTIONS  ") + strings.Join(e.Client.Subscriptions, ", ") + "\n")...)
	result = append(result, (bold("TIMESTAMP      ") + utoa(e.Client.Timestamp) + "\n")...)
	result = append(result, (bold("CHECK          ") + e.Check.Name + "\n")...)
	result = append(result, (bold("COMMAND        ") + e.Check.Command + "\n")...)
	result = append(result, (bold("SUBSCRIBERS    ") + strings.Join(e.Check.Subscribers, ", ") + "\n")...)
	result = append(result, (bold("INTERVAL       ") + strconv.Itoa(e.Check.Interval) + "\n")...)
	result = append(result, (bold("OUTPUT         ") + strings.Replace(e.Check.Output, "\n", " ", -1) + "\n")...)
	result = append(result, (bold("STATUS         ") + statusFg(e.Check.Status) + "\n")...)
	result = append(result, (bold("HISTORY        ") + strings.Join(e.Check.History, ", ") + "\n")...)
	result = append(result, (bold("OCCURRENCES    ") + strconv.Itoa(e.Occurrences) + "\n")...)

	return string(result)
}

func DeleteEventsClientCheck(client string, check string) string {
	_, status := deleteAPI("/events/" + client + "/" + check)
	checkStatus(status)

	return httpStatus(status) + "\n"
}
