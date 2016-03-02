package ohgi

import (
	"fmt"

	"github.com/hico-horiuchi/ohgibone/sensu"
)

func GetClientsHistory(api *sensu.API, client string) string {
	histories, err := api.GetClientsHistory(client)
	checkError(err)

	if len(histories) == 0 {
		return "No histories\n"
	}

	table := newUitable()
	table.AddRow(bold("CHECK"), bold("HISTORY"), bold("TIMESTAMP"))
	for _, history := range histories {
		table.AddRow(
			history.Check,
			colorHistory(stoa(history.History, ", ")),
			utoa(history.LastExecution),
		)
	}

	return fmt.Sprintln(table)
}
