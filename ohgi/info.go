package ohgi

import (
	"strconv"

	"../sensu"
)

func GetInfo(api *sensu.API) string {
	var result []byte

	info, err := api.GetInfo()
	checkError(err)

	result = append(result, (bold("VERSION    ") + info.Sensu.Version + "\n")...)
	result = append(result, (bold("TRANSPORT  ") + strconv.FormatBool(info.Transport.Connected) + "\n")...)
	result = append(result, (bold("REDIS      ") + strconv.FormatBool(info.Redis.Connected) + "\n")...)

	return string(result)
}
