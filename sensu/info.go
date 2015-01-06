package sensu

import (
	"encoding/json"
	"strconv"
)

type sensuStruct struct {
	Version string
}

type connectionStruct struct {
	Connected bool
}

type infoStruct struct {
	Sensu     sensuStruct
	Transport connectionStruct
	Redis     connectionStruct
}

func GetInfo() string {
	var i infoStruct
	var result []byte

	contents := getAPI("/info")
	json.Unmarshal(contents, &i)

	result = append(result, (bold("VERSION    ") + i.Sensu.Version + "\n")...)
	result = append(result, (bold("TRANSPORT  ") + strconv.FormatBool(i.Transport.Connected) + "\n")...)
	result = append(result, (bold("REDIS      ") + strconv.FormatBool(i.Redis.Connected) + "\n")...)

	return string(result)
}
