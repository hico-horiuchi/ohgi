package ohgi

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hico-horiuchi/ohgibone/sensu"
)

func GetResults(api *sensu.API) string {
	results, err := api.GetResults()
	checkError(err)

	if len(results) == 0 {
		return "No current check results\n"
	}

	table := newUitable()
	table.AddRow("", bold("CLIENT"), bold("CHECK"), bold("EXECUTED"))
	for _, result := range results {
		table.AddRow(
			indicateStatus(result.Check.Status),
			result.Client,
			result.Check.Name,
			utoa(result.Check.Executed),
		)
	}

	return fmt.Sprintln(table)
}

func GetResultsWildcard(api *sensu.API, pattern string) string {
	var match []string
	var matches []int
	var result sensu.ResultStruct

	results, err := api.GetResults()
	checkError(err)

	re := regexp.MustCompile("^" + strings.Replace(pattern, "*", ".*", -1) + "$")
	for i, result := range results {
		match = re.FindStringSubmatch(result.Client)
		if len(match) > 0 {
			matches = append(matches, i)
		}
	}

	if len(matches) == 0 {
		return "No current check results that matches " + pattern + "\n"
	}

	table := newUitable()
	table.AddRow("", bold("CLIENT"), bold("CHECK"), bold("EXECUTED"))
	for _, i := range matches {
		result = results[i]
		table.AddRow(
			indicateStatus(result.Check.Status),
			result.Client,
			result.Check.Name,
			utoa(result.Check.Executed),
		)
	}

	return fmt.Sprintln(table)
}

func GetResultsClient(api *sensu.API, client string) string {
	results, err := api.GetResultsClient(client)
	checkError(err)

	if len(results) == 0 {
		return "No current check results for " + client + "\n"
	}

	table := newUitable()
	table.AddRow("", bold("CHECK"), bold("OUTPUT"), bold("EXECUTED"))
	for _, result := range results {
		table.AddRow(
			indicateStatus(result.Check.Status),
			result.Check.Name,
			formatOutput(result.Check.Output),
			utoa(result.Check.Executed),
		)
	}

	return fmt.Sprintln(table)
}

func GetResultsClientCheck(api *sensu.API, client string, check string) string {
	result, err := api.GetResultsClientCheck(client, check)
	checkError(err)

	table := newUitable(true)
	table.AddRow(bold("CLIENT:"), result.Client)
	table.AddRow(bold("CHECK:"), result.Check.Name)
	table.AddRow(bold("COMMAND:"), result.Check.Command)
	table.AddRow(bold("SUBSCRIBERS:"), strings.Join(result.Check.Subscribers, ", "))
	table.AddRow(bold("INTERVAL:"), strconv.Itoa(result.Check.Interval))
	table.AddRow(bold("HANDLERS:"), strings.Join(result.Check.Handlers, ", "))
	table.AddRow(bold("ISSUED:"), utoa(result.Check.Issued))
	table.AddRow(bold("EXECUTED:"), utoa(result.Check.Executed))
	table.AddRow(bold("OUTPUT:"), formatOutput(result.Check.Output))
	table.AddRow(bold("STATUS:"), colorStatus(result.Check.Status))
	table.AddRow(bold("DURATION:"), strconv.FormatFloat(result.Check.Duration, 'f', 3, 64))

	return fmt.Sprintln(table)
}
