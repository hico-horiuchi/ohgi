package ohgi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hico-horiuchi/ohgi/sensu"
)

type configStruct struct {
	Datacenters []datacenterStruct `json:"datacenters"`
}

type datacenterStruct struct {
	sensu.API
	Name string `json:"name"`
}

func LoadConfig(path string, name string) *sensu.API {
	var config configStruct

	bytes, err := getConfig(path)
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

func getConfig(path string) ([]byte, error) {
	home := os.Getenv("HOME")

	if path == "" {
		path = filepath.Join(home, ".ohgi.json")
	} else if len(path) > 1 && path[:2] == "~/" {
		path = filepath.Join(home, path[2:])
	}

	return ioutil.ReadFile(path)
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
