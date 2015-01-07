package sensu

import (
	"encoding/json"
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

	contents := getAPI("/checks")
	json.Unmarshal(contents, &checks)

	if len(checks) == 0 {
		return "No checks\n"
	}

	result = append(result, bold("NAME                COMMAND                                 INTERVAL\n")...)
	for i := range checks {
		c := checks[i]
		line := fillSpace(c.Name, 20) + fillSpace(strings.Replace(c.Command, "\n", " ", -1), 40) + strconv.Itoa(c.Interval) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetChecksCheck(check string) string {
	var c checkStruct
	var result []byte

	contents := getAPI("/checks/" + check)
	json.Unmarshal(contents, &c)

	result = append(result, (bold("NAME         ") + fillSpace(c.Name, 60) + "\n")...)
	result = append(result, (bold("COMMAND      ") + fillSpace(c.Command, 60) + "\n")...)
	result = append(result, (bold("SUBSCRIBERS  ") + fillSpace(strings.Join(c.Subscribers, ", "), 60) + "\n")...)
	result = append(result, (bold("INTERVAL     ") + strconv.Itoa(c.Interval) + "\n")...)

	return string(result)
}
