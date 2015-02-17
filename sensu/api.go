package sensu

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func makeRequest(method string, namespace string, payload io.Reader) *http.Request {
	conf := loadConfig()
	url := "http://" + conf.Host + ":" + strconv.Itoa(conf.Port) + namespace
	request, _ := http.NewRequest(method, url, payload)

	if conf.User != "" && conf.Password != "" {
		request.SetBasicAuth(conf.User, conf.Password)
	}

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return request
}

func doAPI(method string, namespace string, payload io.Reader) ([]byte, int) {
	client := &http.Client{}
	request := makeRequest(method, namespace, payload)
	response, _ := client.Do(request)

	if response == nil {
		fmt.Println("Connection refused")
		os.Exit(1)
	}

	status := response.StatusCode
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return body, status
}

func getAPI(namespace string) ([]byte, int) {
	return doAPI("GET", namespace, nil)
}

func postAPI(namespace string, payload io.Reader) ([]byte, int) {
	return doAPI("POST", namespace, payload)
}

func deleteAPI(namespace string) ([]byte, int) {
	return doAPI("DELETE", namespace, nil)
}
