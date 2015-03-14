package sensu

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type configStruct struct {
	Host     string
	Port     int
	User     string
	Password string
}

func loadConfig() configStruct {
	var config configStruct
	bytes, err := ioutil.ReadFile(os.Getenv("HOME") + "/.sensu.json")

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(bytes, &config)
	return config
}
