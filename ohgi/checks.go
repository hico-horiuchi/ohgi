package ohgi

import (
	"regexp"
	"strconv"
	"strings"

	"../sensu"
)

func GetChecks(api *sensu.API) string {
	var line string

	checks, err := api.GetChecks()
	checkError(err)

	if len(checks) == 0 {
		return "No checks\n"
	}

	result := []byte(bold("NAME                          COMMAND                                                     INTERVAL\n"))
	for _, check := range checks {
		line = fillSpace(check.Name, 30) + fillSpace(check.Command, 60) + strconv.Itoa(check.Interval) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetChecksWildcard(api *sensu.API, pattern string) string {
	var match []string
	var matches []int
	var line string

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
		return "No checks\n"
	}

	result := []byte(bold("NAME                          COMMAND                                                     INTERVAL\n"))
	for _, i := range matches {
		check := checks[i]
		line = fillSpace(check.Name, 30) + fillSpace(check.Command, 60) + strconv.Itoa(check.Interval) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetChecksCheck(api *sensu.API, name string) string {
	var result []byte

	check, err := api.GetChecksCheck(name)
	checkError(err)

	result = append(result, (bold("NAME         ") + check.Name + "\n")...)
	result = append(result, (bold("COMMAND      ") + check.Command + "\n")...)
	result = append(result, (bold("SUBSCRIBERS  ") + strings.Join(check.Subscribers, ", ") + "\n")...)
	result = append(result, (bold("INTERVAL     ") + strconv.Itoa(check.Interval) + "\n")...)
	result = append(result, (bold("HANDLERS     ") + strings.Join(check.Handlers, ", ") + "\n")...)

	return string(result)
}
