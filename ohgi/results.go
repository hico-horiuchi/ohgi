package ohgi

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/hico-horiuchi/ohgi/sensu"
)

func GetResults(api *sensu.API) string {
	var line string

	results, err := api.GetResults()
	checkError(err)

	if len(results) == 0 {
		return "No current check results\n"
	}

	print := []byte(bold("  CLIENT                                  CHECK                         EXECUTED\n"))
	for _, result := range results {
		line = indicateStatus(result.Check.Status) + fillSpace(result.Client, 40) + fillSpace(result.Check.Name, 30) + utoa(result.Check.Executed) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func GetResultsWildcard(api *sensu.API, pattern string) string {
	var match []string
	var matches []int
	var line string

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
		return "No current check results that match " + pattern + "\n"
	}

	print := []byte(bold("  CLIENT                                  CHECK                         EXECUTED\n"))
	for _, i := range matches {
		result := results[i]
		line = indicateStatus(result.Check.Status) + fillSpace(result.Client, 40) + fillSpace(result.Check.Name, 30) + utoa(result.Check.Executed) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func GetResultsClient(api *sensu.API, client string) string {
	var line, output string

	results, err := api.GetResultsClient(client)
	checkError(err)

	if len(results) == 0 {
		return "No current check results for " + client + "\n"
	}

	print := []byte(bold("  CHECK                         OUTPUT                                            EXECUTED\n"))
	for _, result := range results {
		output = strings.Replace(result.Check.Output, "\n", " ", -1)
		line = indicateStatus(result.Check.Status) + fillSpace(result.Check.Name, 30) + fillSpace(output, 50) + utoa(result.Check.Executed) + "\n"
		print = append(print, line...)
	}

	return string(print)
}

func GetResultsClientCheck(api *sensu.API, client string, check string) string {
	var print []byte

	result, err := api.GetResultsClientCheck(client, check)
	checkError(err)

	print = append(print, (bold("CLIENT         ") + result.Client + "\n")...)
	print = append(print, (bold("CHECK          ") + result.Check.Name + "\n")...)
	print = append(print, (bold("COMMAND        ") + result.Check.Command + "\n")...)
	print = append(print, (bold("SUBSCRIBERS    ") + strings.Join(result.Check.Subscribers, ", ") + "\n")...)
	print = append(print, (bold("INTERVAL       ") + strconv.Itoa(result.Check.Interval) + "\n")...)
	print = append(print, (bold("HANDLERS       ") + strings.Join(result.Check.Handlers, ", ") + "\n")...)
	print = append(print, (bold("ISSUED         ") + utoa(result.Check.Issued) + "\n")...)
	print = append(print, (bold("EXECUTED       ") + utoa(result.Check.Executed) + "\n")...)
	print = append(print, (bold("OUTPUT         ") + strings.Replace(result.Check.Output, "\n", " ", -1) + "\n")...)
	print = append(print, (bold("STATUS         ") + paintStatus(result.Check.Status) + "\n")...)
	print = append(print, (bold("DURATION       ") + strconv.FormatFloat(result.Check.Duration, 'f', 3, 64) + "\n")...)

	return string(print)
}
