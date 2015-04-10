package ohgi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var config configStruct
var timeout = 1 * time.Second

type configStruct struct {
	Host     string
	Port     int
	User     string
	Password string
	Timeout  int
}

func LoadConfig() {
	bytes, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ohgi.json")
	checkError(err)
	json.Unmarshal(bytes, &config)

	if config.Timeout > 0 {
		timeout = time.Duration(config.Timeout) * time.Second
	}
	http.DefaultClient.Timeout = timeout
}
