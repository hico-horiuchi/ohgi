package sensu

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func makeRequest(method string, namespace string, payload io.Reader) *http.Request {
	config := loadConfig()
	url := "http://" + config.Host + ":" + strconv.Itoa(config.Port) + namespace
	request, err := http.NewRequest(method, url, payload)
	checkError(err)

	if config.User != "" && config.Password != "" {
		request.SetBasicAuth(config.User, config.Password)
	}

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return request
}

func doAPI(method string, namespace string, payload io.Reader) ([]byte, int) {
	request := makeRequest(method, namespace, payload)
	response, err := http.DefaultClient.Do(request)
	checkError(err)

	status := response.StatusCode
	body, err := ioutil.ReadAll(response.Body)
	checkError(err)

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
