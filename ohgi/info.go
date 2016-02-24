package ohgi

import (
	"fmt"
	"strconv"

	"github.com/hico-horiuchi/ohgi/sensu"
)

func GetInfo(api *sensu.API) string {
	info, err := api.GetInfo()
	checkError(err)

	table := newUitable(true)
	table.AddRow(bold("VERSION:"), info.Sensu.Version)
	table.AddRow(bold("TRANSPORT:"), strconv.FormatBool(info.Transport.Connected))
	table.AddRow(bold("REDIS:"), strconv.FormatBool(info.Redis.Connected))

	return fmt.Sprintln(table)
}
