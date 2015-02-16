package sensu

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type clientStruct struct {
	Name          string
	Address       string
	Subscriptions []string
	Timestamp     int64
}

func GetClients(limit int, offset int) string {
	var clients []clientStruct
	var result []byte

	contents, status := getAPI(fmt.Sprintf("/clients?limit=%d&offset=%d", limit, offset))
	if status != 200 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	json.Unmarshal(contents, &clients)
	if len(clients) == 0 {
		return "No clients\n"
	}

	result = append(result, bold("NAME                ADDRESS             TIMESTAMP\n")...)
	for i := range clients {
		c := clients[i]
		line := fillSpace(c.Name, 20) + fillSpace(c.Address, 20) + utoa(c.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetClientsClient(client string) string {
	var c clientStruct
	var result []byte

	contents, status := getAPI("/clients/" + client)
	if status != 200 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	json.Unmarshal(contents, &c)

	result = append(result, (bold("NAME           ") + fillSpace(c.Name, 60) + "\n")...)
	result = append(result, (bold("ADDRESS        ") + fillSpace(c.Address, 60) + "\n")...)
	result = append(result, (bold("SUBSCRIPTIONS  ") + fillSpace(strings.Join(c.Subscriptions, ", "), 60) + "\n")...)
	result = append(result, (bold("TIMESTAMP      ") + utoa(c.Timestamp) + "\n")...)

	return string(result)
}
