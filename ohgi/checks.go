package ohgi

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hico-horiuchi/ohgibone/sensu"
)

func GetChecks(api *sensu.API) string {
	checks, err := api.GetChecks()
	checkError(err)

	if len(checks) == 0 {
		return "No checks\n"
	}

	table := newUitable()
	table.AddRow(bold("NAME"), bold("COMMAND"), bold("INTERVAL"))
	for _, check := range checks {
		table.AddRow(
			check.Name,
			check.Command,
			strconv.Itoa(check.Interval),
		)
	}

	return fmt.Sprintln(table)
}

func GetChecksWildcard(api *sensu.API, pattern string) string {
	var match []string
	var matches []int
	var check sensu.CheckStruct

	checks, err := api.GetChecks()
	checkError(err)

	re := regexp.MustCompile("^" + strings.Replace(pattern, "*", ".*", -1) + "$")
	for i, check := range checks {
		match = re.FindStringSubmatch(check.Name)
		if len(match) > 0 {
			matches = append(matches, i)
		}
	}

	if len(matches) == 0 {
		return "No checks that matches " + pattern + "\n"
	}

	table := newUitable()
	table.AddRow(bold("NAME"), bold("COMMAND"), bold("INTERVAL"))
	for _, i := range matches {
		check = checks[i]
		table.AddRow(
			check.Name,
			check.Command,
			strconv.Itoa(check.Interval),
		)
	}

	return fmt.Sprintln(table)
}

func GetChecksCheck(api *sensu.API, name string) string {
	check, err := api.GetChecksCheck(name)
	checkError(err)

	table := newUitable(true)
	table.AddRow(bold("NAME:"), check.Name)
	table.AddRow(bold("COMMAND:"), check.Command)
	table.AddRow(bold("SUBSCRIBERS:"), strings.Join(check.Subscribers, ", "))
	table.AddRow(bold("INTERVAL:"), strconv.Itoa(check.Interval))
	table.AddRow(bold("HANDLERS:"), strings.Join(check.Handlers, ", "))

	return fmt.Sprintln(table)
}
