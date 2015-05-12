package ohgi

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type clientStruct struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Timestamp     int64
}

func GetClients(limit int, offset int) string {
	var clients []clientStruct
	var result []byte

	contents, status := getAPI(fmt.Sprintf("/clients?limit=%d&offset=%d", limit, offset))
	checkStatus(status)

	json.Unmarshal(contents, &clients)
	if len(clients) == 0 {
		return "No clients\n"
	}

	result = append(result, bold("NAME                                    ADDRESS                                 TIMESTAMP\n")...)
	for _, c := range clients {
		line := fillSpace(c.Name, 40) + fillSpace(c.Address, 40) + utoa(c.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetClientsWildcard(pattern string) string {
	var clients []clientStruct
	var result []byte
	var matches []int
	re := regexp.MustCompile("^" + strings.Replace(pattern, "*", ".*", -1) + "$")

	contents, status := getAPI("/clients")
	checkStatus(status)

	json.Unmarshal(contents, &clients)
	for i, c := range clients {
		match := re.FindStringSubmatch(c.Name)
		if len(match) > 0 {
			matches = append(matches, i)
		}
	}

	if len(matches) == 0 {
		return "No clients\n"
	}

	result = append(result, bold("NAME                                    ADDRESS                                 TIMESTAMP\n")...)
	for _, i := range matches {
		c := clients[i]
		line := fillSpace(c.Name, 40) + fillSpace(c.Address, 40) + utoa(c.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetClientsClient(client string) string {
	var c clientStruct
	var result []byte

	contents, status := getAPI("/clients/" + client)
	checkStatus(status)

	json.Unmarshal(contents, &c)

	result = append(result, (bold("NAME           ") + c.Name + "\n")...)
	result = append(result, (bold("ADDRESS        ") + c.Address + "\n")...)
	result = append(result, (bold("SUBSCRIPTIONS  ") + strings.Join(c.Subscriptions, ", ") + "\n")...)
	result = append(result, (bold("TIMESTAMP      ") + utoa(c.Timestamp) + "\n")...)

	return string(result)
}

func PostClients(name string, address string, subscriptions string) string {
	c := clientStruct{
		Name:          name,
		Address:       address,
		Subscriptions: strings.Split(subscriptions, ","),
	}

	body, err := json.Marshal(c)
	checkError(err)

	payload := strings.NewReader(string(body))
	_, status := postAPI("/clients", payload)
	checkStatus(status)

	return httpStatus(status) + "\n"
}

func DeleteClientsClient(client string) string {
	_, status := deleteAPI("/clients/" + client)
	checkStatus(status)

	return httpStatus(status) + "\n"
}
