// The Sensu API provides access to the data that Sensu servers collect, such as client information & current events.
// The API can be used to resolve events and request check executions, among other things.
package sensu

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var DefaultAPI *API = &API{
	Host: "localhost",
	Port: 4567,
}

type API struct {
	Host     string
	Port     int
	User     string
	Password string
}

type apiResponse struct {
	Body       string
	StatusCode int
}

func (api API) get(namespace string) (*apiResponse, error) {
	return api.do("GET", namespace, nil)
}

func (api API) post(namespace string, payload io.Reader) (*apiResponse, error) {
	return api.do("POST", namespace, payload)
}

func (api API) delete(namespace string) (*apiResponse, error) {
	return api.do("DELETE", namespace, nil)
}

func (api API) do(method string, namespace string, payload io.Reader) (*apiResponse, error) {
	request, err := api.newRequest(method, namespace, payload)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &apiResponse{
		Body:       string(body),
		StatusCode: response.StatusCode,
	}, nil
}

func (api API) newRequest(method string, namespace string, payload io.Reader) (*http.Request, error) {
	url := "http://" + api.Host + ":" + strconv.Itoa(api.Port) + namespace

	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	if api.User != "" && api.Password != "" {
		request.SetBasicAuth(api.User, api.Password)
	}

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return request, nil
}
