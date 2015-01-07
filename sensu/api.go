package sensu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func getAPI(namespace string) ([]byte, int) {
	var status int
	var body []byte

	conf := loadConfig()
	url := "http://" + conf.Host + ":" + strconv.Itoa(conf.Port) + namespace
	request, _ := http.NewRequest("GET", url, nil)

	if conf.User != "" && conf.Password != "" {
		request.SetBasicAuth(conf.User, conf.Password)
	}

	client := &http.Client{}
	response, _ := client.Do(request)

	if response == nil {
		fmt.Println("Connection refused")
		os.Exit(1)
		return body, status
	} else {
		status := response.StatusCode
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		return body, status
	}
}
