package ohgi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var datacenter dcStruct
var timeout = 3 * time.Second

type configStruct struct {
	Datacenters []dcStruct
	Timeout     int
}

type dcStruct struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}

func LoadConfig(dc string) {
	var c configStruct

	bytes, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ohgi.json")
	checkError(err)

	json.Unmarshal(bytes, &c)
	datacenter, err = c.selectDatacenter(dc)
	checkError(err)

	if c.Timeout > 0 {
		timeout = time.Duration(c.Timeout) * time.Second
	}
	http.DefaultClient.Timeout = timeout
}

func (c configStruct) selectDatacenter(name string) (dcStruct, error) {
	switch {
	case len(c.Datacenters) < 1:
		return dcStruct{}, errors.New("ohgi: no datacenters in config")
	case name == "":
		return c.Datacenters[0], nil
	}

	for _, dc := range c.Datacenters {
		if dc.Name == name {
			return dc, nil
		}
	}

	return dcStruct{}, errors.New("ohgi: no such datacenter in config")
}
