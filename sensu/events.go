package sensu

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

	contents := requestAPI("GET", "events")
	json.Unmarshal(contents, &events)

	if len(events) == 0 {
		return "No current events\n"
	}

	result = append(result, bold("  CLIENT              CHECK               #         TIME\n")...)
	for i := range events {
		e := events[i]
		occurrences := strconv.Itoa(e.Occurrences)
		line := statusColor(e.Check.Status) + fillSpace(e.Client.Name, 20) + fillSpace(e.Check.Name, 20) + fillSpace(occurrences, 10) + utoa(e.Client.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetEventsClient(client string) string {
	var events []eventStruct
	var result []byte

	contents := requestAPI("GET", "events/"+client)
	json.Unmarshal(contents, &events)

	if len(events) == 0 {
		return "No current events for " + client + "\n"
	}

	result = append(result, bold("  CHECK               OUTPUT                            TIME\n")...)
	for i := range events {
		e := events[i]
		output := strings.Replace(e.Check.Output, "\n", " ", -1)
		line := statusColor(e.Check.Status) + fillSpace(e.Check.Name, 20) + fillSpace(output, 35) + utoa(e.Client.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}
