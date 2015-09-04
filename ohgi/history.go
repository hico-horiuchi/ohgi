package ohgi

import "github.com/hico-horiuchi/ohgi/sensu"

func GetClientsHistory(api *sensu.API, client string) string {
	var line string

	histories, err := api.GetClientsHistory(client)
	checkError(err)

	if len(histories) == 0 {
		return "No histories\n"
	}

	print := []byte(bold("CHECK                         HISTORY                                         TIMESTAMP\n"))
	for _, history := range histories {
		line = fillSpace(history.Check, 30) + paintHistory(fillSpace(stoa(history.History, ", "), 48)) + utoa(history.LastExecution) + "\n"
		print = append(print, line...)
	}

	return string(print)
}
