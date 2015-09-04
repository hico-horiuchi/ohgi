package ohgi

import (
	"strings"
	"time"

	"github.com/hico-horiuchi/ohgi/sensu"
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
	var line string
	var path []string
	var expire string

	err := api.GetStashes(&silences, -1, -1)
	checkError(err)

	if len(silences) == 0 {
		return "No silence stashes\n"
	}

	print := []byte(bold("CLIENT                                  CHECK                         REASON                        EXPIRATION\n"))
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

		line = fillSpace(path[1], 40) + fillSpace(path[2], 30) + fillSpace(silence.Content.Reason, 30) + expire + "\n"
		print = append(print, line...)
	}

	return string(print)
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

	err := api.PostStashes(silence)
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
