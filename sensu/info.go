package sensu

import (
	"encoding/json"
	"strconv"
)

type infoStruct struct {
	Sensu struct {
		Version string
	}
	Transport struct {
		Connected bool
	}
	Redis struct {
		Connected bool
	}
}

func GetInfo() string {
	var i infoStruct
	var result []byte

	contents, status := getAPI("/info")
	checkStatus(status)

	json.Unmarshal(contents, &i)

	result = append(result, (bold("VERSION    ") + i.Sensu.Version + "\n")...)
	result = append(result, (bold("TRANSPORT  ") + strconv.FormatBool(i.Transport.Connected) + "\n")...)
	result = append(result, (bold("REDIS      ") + strconv.FormatBool(i.Redis.Connected) + "\n")...)

	return string(result)
}
