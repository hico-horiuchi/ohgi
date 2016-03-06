package ohgi

import (
	"fmt"
	"strings"
	"time"

	"github.com/hico-horiuchi/ohgibone/sensu"
)

type silenceStruct struct {
	sensu.StashStruct
	Content contentStruct `json:"content"`
}

type contentStruct struct {
	Timestamp float64 `json:"timestamp"`
	Source    string  `json:"source"`
	Reason    string  `json:"reason"`
}

func GetSilence(api *sensu.API) string {
	var silences []silenceStruct
	var path []string
	var expire string

	err := api.GetStashes(&silences, -1, -1)
	checkError(err)

	if len(silences) == 0 {
		return "No silence stashes\n"
	}

	table := newUitable()
	table.AddRow(bold("CLIENT"), bold("CHECK"), bold("REASON"), bold("EXPIRATION"))
	for _, silence := range silences {
		path = strings.Split(silence.Path, "/")
		if path[0] != "silence" {
			continue
		} else if len(path) == 2 {
			path = append(path, "")
		}

		if silence.Expire == -1 {
			expire = "Never"
		} else {
			expire = utoa(time.Now().Unix() + silence.Expire)
		}

		table.AddRow(
			path[1],
			path[2],
			silence.Content.Reason,
			expire,
		)
	}

	return fmt.Sprintln(table)
}

func PostSilence(api *sensu.API, client string, check string, expiration string, reason string) string {
	var path string

	if check == "" {
		path = "silence/" + client
	} else {
		path = "silence/" + client + "/" + check
	}

	silence := silenceStruct{
		StashStruct: sensu.StashStruct{
			Expire: stoe(expiration),
			Path:   path,
		},
		Content: contentStruct{
			Timestamp: float64(time.Now().Unix()),
			Source:    "ohgi",
			Reason:    reason,
		},
	}

	err := api.PostStashes(&silence)
	checkError(err)

	return "Created\n"
}

func DeleteSilence(api *sensu.API, client string, check string) string {
	var path string

	if check == "" {
		path = "silence/" + client
	} else {
		path = "silence/" + client + "/" + check
	}

	err := api.DeleteStashesPath(path)
	checkError(err)

	return "No Content\n"
}
