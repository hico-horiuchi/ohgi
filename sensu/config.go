package sensu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	Host     string
	Port     int
	User     string
	Password string
}

func loadConfig() config {
	var conf config
	contents, _ := ioutil.ReadFile(os.Getenv("HOME") + "/.sensu.json")

	if contents == nil {
		fmt.Println("Cannot open ~/.sensu.json")
		os.Exit(1)
	}

	json.Unmarshal(contents, &conf)
	return conf
}
