package sensu

import (
	"encoding/json"
	"strconv"
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

	result = append(result, bold("  CLIENT              CHECK               #         TIME\n")...)
	for i := range events {
		e := events[i]
		line := statusColor(e.Check.Status) + fillSpace(e.Client.Name, 20) + fillSpace(e.Check.Name, 20) + fillSpace(strconv.Itoa(e.Occurrences), 10) + utoa(e.Client.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}
