package ohgi

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func makeRequest(method string, namespace string, payload io.Reader) *http.Request {
	url := "http://" + datacenter.Host + ":" + strconv.Itoa(datacenter.Port) + namespace
	request, err := http.NewRequest(method, url, payload)
	checkError(err)

	if datacenter.User != "" && datacenter.Password != "" {
		request.SetBasicAuth(datacenter.User, datacenter.Password)
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
