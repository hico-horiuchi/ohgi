package sensu

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type configStruct struct {
	Host     string
	Port     int
	User     string
	Password string
}

func loadConfig() configStruct {
	bytes, err := ioutil.ReadFile(os.Getenv("HOME") + "/.sensu.json")
	checkError(err)

	var config configStruct
	json.Unmarshal(bytes, &config)

	return config
}
