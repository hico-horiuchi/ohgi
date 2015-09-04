package ohgi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/hico-horiuchi/ohgi/sensu"
)

type configStruct struct {
	Datacenters []datacenterStruct `json:"datacenters"`
}

type datacenterStruct struct {
	sensu.API
	Name string `json:"name"`
}

func LoadConfig(name string) *sensu.API {
	var config configStruct

	bytes, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ohgi.json")
	checkError(err)

	err = json.Unmarshal(bytes, &config)
	checkError(err)

	datacenter, err := config.selectDatacenter(name)
	checkError(err)

	return &sensu.API{
		Host:     datacenter.Host,
		Port:     datacenter.Port,
		User:     datacenter.User,
		Password: datacenter.Password,
	}
}

func (config configStruct) selectDatacenter(name string) (*datacenterStruct, error) {
	switch {
	case len(config.Datacenters) == 0:
		return nil, errors.New("ohgi: no datacenters in config")
	case name == "":
		return &config.Datacenters[0], nil
	}

	for _, datacenter := range config.Datacenters {
		if datacenter.Name == name {
			return &datacenter, nil
		}
	}

	return nil, errors.New("ohgi: no such datacenter in config")
}
