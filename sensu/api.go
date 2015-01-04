package sensu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func requestAPI(method string, namespace string) []byte {
	var body []byte

	conf := loadConfig()
	url := "http://" + conf.Host + ":" + strconv.Itoa(conf.Port) + "/" + namespace
	request, _ := http.NewRequest(method, url, nil)

	if conf.User != "" && conf.Password != "" {
		request.SetBasicAuth(conf.User, conf.Password)
	}

	client := &http.Client{}
	response, _ := client.Do(request)

	if response == nil {
		fmt.Println("Connection refused")
		os.Exit(1)
		return body
	} else {
		status := response.StatusCode
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()

		if status != 200 {
			fmt.Println(httpStatus(status))
			os.Exit(1)
		}

		return body
	}
}
