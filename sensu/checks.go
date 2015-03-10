package sensu

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type checkStruct struct {
	Name        string
	Command     string
	Subscribers []string
	Interval    int
	Issued      int
	Executed    int
	Output      string
	Status      int
	Duration    float32
	History     []string
}

func GetChecks() string {
	var checks []checkStruct
	var result []byte

	contents, status := getAPI("/checks")
	if status != 200 {
		log.Fatal(httpStatus(status))
	}

	json.Unmarshal(contents, &checks)
	if len(checks) == 0 {
		return "No checks\n"
	}

	result = append(result, bold("NAME                          COMMAND                                                     INTERVAL\n")...)
	for i := range checks {
		c := checks[i]
		line := fillSpace(c.Name, 30) + fillSpace(c.Command, 60) + strconv.Itoa(c.Interval) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetChecksWildcard(pattern string) string {
	var checks []checkStruct
	var result []byte
	var matches []int
	re := regexp.MustCompile(strings.Replace(pattern, "*", ".*", -1))

	contents, status := getAPI("/checks")
	if status != 200 {
		log.Fatal(httpStatus(status))
	}

	json.Unmarshal(contents, &checks)
	for i := range checks {
		c := checks[i]
		match := re.FindStringSubmatch(c.Name)
		if len(match) > 0 {
			matches = append(matches, i)
		}
	}

	if len(matches) == 0 {
		return "No checks\n"
	}

	result = append(result, bold("NAME                          COMMAND                                                     INTERVAL\n")...)
	for _, i := range matches {
		c := checks[i]
		line := fillSpace(c.Name, 30) + fillSpace(c.Command, 60) + strconv.Itoa(c.Interval) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetChecksCheck(check string) string {
	var c checkStruct
	var result []byte

	contents, status := getAPI("/checks/" + check)
	if status != 200 {
		log.Fatal(httpStatus(status))
	}

	json.Unmarshal(contents, &c)

	result = append(result, (bold("NAME         ") + c.Name + "\n")...)
	result = append(result, (bold("COMMAND      ") + c.Command + "\n")...)
	result = append(result, (bold("SUBSCRIBERS  ") + strings.Join(c.Subscribers, ", ") + "\n")...)
	result = append(result, (bold("INTERVAL     ") + strconv.Itoa(c.Interval) + "\n")...)

	return string(result)
}
