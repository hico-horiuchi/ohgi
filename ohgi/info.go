package ohgi

import (
	"strconv"

	"github.com/hico-horiuchi/ohgi/sensu"
)

func GetInfo(api *sensu.API) string {
	var print []byte

	info, err := api.GetInfo()
	checkError(err)

	print = append(print, (bold("VERSION    ") + info.Sensu.Version + "\n")...)
	print = append(print, (bold("TRANSPORT  ") + strconv.FormatBool(info.Transport.Connected) + "\n")...)
	print = append(print, (bold("REDIS      ") + strconv.FormatBool(info.Redis.Connected) + "\n")...)

	return string(print)
}
