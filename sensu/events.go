package sensu

import (
	"encoding/json"
	"fmt"
	"os"
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
	if status != 200 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	json.Unmarshal(contents, &events)
	if len(events) == 0 {
		return "No current events\n"
	}

	result = append(result, bold("  CLIENT              CHECK               #         TIMESTAMP\n")...)
	for i := range events {
		e := events[i]
		occurrences := strconv.Itoa(e.Occurrences)
		line := statusBg(e.Check.Status) + fillSpace(e.Client.Name, 20) + fillSpace(e.Check.Name, 20) + fillSpace(occurrences, 10) + utoa(e.Client.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetEventsClient(client string) string {
	var events []eventStruct
	var result []byte

	contents, status := getAPI("/events/" + client)
	if status != 200 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	json.Unmarshal(contents, &events)
	if len(events) == 0 {
		return "No current events for " + client + "\n"
	}

	result = append(result, bold("  CHECK               OUTPUT                            TIMESTAMP\n")...)
	for i := range events {
		e := events[i]
		output := strings.Replace(e.Check.Output, "\n", " ", -1)
		line := statusBg(e.Check.Status) + fillSpace(e.Check.Name, 20) + fillSpace(output, 35) + utoa(e.Client.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetEventsClientCheck(client string, check string) string {
	var e eventStruct
	var result []byte

	contents, status := getAPI("/events/" + client + "/" + check)
	if status != 200 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	json.Unmarshal(contents, &e)

	result = append(result, (bold("CLIENT         ") + fillSpace(e.Client.Name, 60) + "\n")...)
	result = append(result, (bold("ADDRESS        ") + fillSpace(e.Client.Address, 60) + "\n")...)
	result = append(result, (bold("SUBSCRIPTIONS  ") + fillSpace(strings.Join(e.Client.Subscriptions, ", "), 60) + "\n")...)
	result = append(result, (bold("TIMESTAMP      ") + utoa(e.Client.Timestamp) + "\n")...)
	result = append(result, (bold("CHECK          ") + fillSpace(e.Check.Name, 60) + "\n")...)
	result = append(result, (bold("COMMAND        ") + fillSpace(e.Check.Command, 60) + "\n")...)
	result = append(result, (bold("SUBSCRIBERS    ") + fillSpace(strings.Join(e.Check.Subscribers, ", "), 60) + "\n")...)
	result = append(result, (bold("INTERVAL       ") + strconv.Itoa(e.Check.Interval) + "\n")...)
	result = append(result, (bold("OUTPUT         ") + fillSpace(strings.Replace(e.Check.Output, "\n", " ", -1), 60) + "\n")...)
	result = append(result, (bold("STATUS         ") + statusFg(e.Check.Status) + "\n")...)
	result = append(result, (bold("HISTORY        ") + fillSpace(strings.Join(e.Check.History, ", "), 60) + "\n")...)
	result = append(result, (bold("OCCURRENCES    ") + strconv.Itoa(e.Occurrences) + "\n")...)

	return string(result)
}
