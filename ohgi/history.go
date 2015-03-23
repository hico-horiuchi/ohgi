package ohgi

import (
	"encoding/json"
)

type historyStruct struct {
	Check          string
	History        []int
	Last_execution int64
	Last_status    int
}

func GetHistory(client string) string {
	var histories []historyStruct
	var result []byte

	contents, status := getAPI("/clients/" + client + "/history")
	checkStatus(status)

	json.Unmarshal(contents, &histories)
	if len(histories) == 0 {
		return "No historiess\n"
	}

	result = append(result, bold("CHECK                         HISTORY                                         TIMESTAMP\n")...)
	for _, h := range histories {
		history := stoa(h.History, ", ")
		line := fillSpace(h.Check, 30) + fillSpace(history, 48) + utoa(h.Last_execution) + "\n"
		result = append(result, line...)
	}

	return string(result)
}
