package sensu

import (
	"encoding/json"
	"fmt"
	"os"
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
	if status != 200 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	json.Unmarshal(contents, &histories)
	if len(histories) == 0 {
		return "No historiess\n"
	}

	result = append(result, bold("CHECK                         HISTORY                                         TIMESTAMP\n")...)
	for i := range histories {
		h := histories[i]
		history := stoa(h.History, ", ")
		line := fillSpace(h.Check, 30) + fillSpace(history, 48) + utoa(h.Last_execution) + "\n"
		result = append(result, line...)
	}

	return string(result)
}
