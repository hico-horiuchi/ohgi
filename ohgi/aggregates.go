package ohgi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hico-horiuchi/ohgi/sensu"
)

func GetAggregates(api *sensu.API, limit int, offset int) string {
	aggregates, err := api.GetAggregates(limit, offset)
	checkError(err)

	if len(aggregates) == 0 {
		return "No aggregates\n"
	}

	table := newUitable()
	table.AddRow(bold("CHECK"), bold("ISSUES"))
	for _, aggregate := range aggregates {
		table.AddRow(
			aggregate.Check,
			strconv.Itoa(len(aggregate.Issued)),
		)
	}

	return fmt.Sprintln(table)
}

func GetAggregatesCheck(api *sensu.API, check string, age int) string {
	issues, err := api.GetAggregatesCheck(check, age)
	checkError(err)

	if len(issues) == 0 {
		return "No aggregates\n"
	}

	table := newUitable()
	table.AddRow(bold("TIMESTAMP"), bold("ISSUED"))
	for _, issued := range issues {
		table.AddRow(
			utoa(issued),
			strconv.FormatInt(issued, 10),
		)
	}

	return fmt.Sprintln(table)
}

func DeleteAggregatesCheck(api *sensu.API, check string) string {
	err := api.DeleteAggregatesCheck(check)
	checkError(err)

	return "No Content\n"
}

func GetAggregatesCheckIssued(api *sensu.API, check string, issued string, summarize string, results bool) string {
	if issued == "latest" {
		issues, err := api.GetAggregatesCheck(check, -1)
		checkError(err)

		if len(issues) == 0 {
			return "No aggregates\n"
		}
		issued = strconv.FormatInt(issues[0], 10)
	}

	i, err := strconv.ParseInt(issued, 10, 64)
	checkError(err)

	aggregate, err := api.GetAggregatesCheckIssued(check, i, summarize, results)
	checkError(err)

	table := newUitable(true)
	if summarize == "output" {
		for output, j := range aggregate.Outputs {
			table.AddRow(
				bold(formatOutput(output)+":"),
				strconv.Itoa(j),
			)
		}

		return fmt.Sprintln(table)
	}

	table.AddRow(bold("OK:"), fgColor(strconv.Itoa(aggregate.Ok), 0))
	table.AddRow(bold("WARNING:"), fgColor(strconv.Itoa(aggregate.Warning), 1))
	table.AddRow(bold("CRITICAL:"), fgColor(strconv.Itoa(aggregate.Critical), 2))
	table.AddRow(bold("UNKNOWN:"), fgColor(strconv.Itoa(aggregate.Unknown), 3))
	table.AddRow(bold("TOTAL:"), strconv.Itoa(aggregate.Total))

	if results {
		clients := make([]string, len(aggregate.Results))
		for j, result := range aggregate.Results {
			clients[j] = result.Client
		}

		table.AddRow(bold("CLIENTS:"), strings.Join(clients, ", "))
	}

	return fmt.Sprintln(table)
}
