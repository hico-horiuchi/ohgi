package ohgi

import (
	"strconv"
	"strings"

	"../sensu"
)

func GetAggregates(api *sensu.API, limit int, offset int) string {
	var line string

	aggregates, err := api.GetAggregates(limit, offset)
	checkError(err)

	if len(aggregates) == 0 {
		return "No aggregates\n"
	}

	print := []byte(bold("CHECK                         ISSUES\n"))
	for _, aggregate := range aggregates {
		line = fillSpace(aggregate.Check, 30) + strconv.Itoa(len(aggregate.Issued)) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func GetAggregatesCheck(api *sensu.API, check string, age int) string {
	var line string

	issues, err := api.GetAggregatesCheck(check, age)
	checkError(err)

	if len(issues) == 0 {
		return "No aggregates\n"
	}

	print := []byte(bold("TIMESTAMP            ISSUED\n"))
	for _, issued := range issues {
		line = utoa(issued) + "  " + strconv.FormatInt(issued, 10) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func DeleteAggregatesCheck(api *sensu.API, check string) string {
	err := api.DeleteAggregatesCheck(check)
	checkError(err)

	return "No Content\n"
}

func GetAggregatesCheckIssued(api *sensu.API, check string, issued string, results bool) string {
	var print []byte

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

	aggregate, err := api.GetAggregatesCheckIssued(check, i, results)
	checkError(err)

	print = append(print, (bold("OK        ") + frontColor(strconv.Itoa(aggregate.Ok), 0) + "\n")...)
	print = append(print, (bold("WARNING   ") + frontColor(strconv.Itoa(aggregate.Warning), 1) + "\n")...)
	print = append(print, (bold("CRITICAL  ") + frontColor(strconv.Itoa(aggregate.Critical), 2) + "\n")...)
	print = append(print, (bold("UNKNOWN   ") + frontColor(strconv.Itoa(aggregate.Unknown), 3) + "\n")...)
	print = append(print, (bold("TOTAL     ") + strconv.Itoa(aggregate.Total) + "\n")...)

	if results {
		clients := make([]string, len(aggregate.Results))
		for j, result := range aggregate.Results {
			clients[j] = result.Client
		}

		print = append(print, (bold("CLIENTS   ") + strings.Join(clients, ", ") + "\n")...)
	}

	return string(print)
}
