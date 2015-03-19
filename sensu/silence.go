package sensu

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type silenceStruct struct {
	Path    string
	Content struct {
		Reason    string
		Source    string
		Timestamp float32
	}
	Expire int64
}

func GetSilence() string {
	var silences []silenceStruct
	var result []byte
	var expire string

	contents, status := getAPI("/stashes")
	checkStatus(status)

	json.Unmarshal(contents, &silences)
	if len(silences) == 0 {
		return "No silences\n"
	}

	result = append(result, bold("CLIENT                                  CHECK                         REASON                        EXPIRATION\n")...)
	for i := range silences {
		s := silences[i]
		path := strings.Split(s.Path, "/")

		if path[0] != "silence" {
			continue
		} else if len(path) == 2 {
			path = append(path, "")
		}

		if s.Expire == -1 {
			expire = "Never"
		} else {
			expire = utoa(time.Now().Unix() + s.Expire)
		}

		line := fillSpace(path[1], 40) + fillSpace(path[2], 30) + fillSpace(s.Content.Reason, 30) + expire + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func PostSilence(client string, check string, expiration string, reason string) string {
	var body string
	var path string

	if check == "" {
		path = "silence/" + client
	} else {
		path = "silence/" + client + "/" + check
	}

	now := strconv.FormatInt(time.Now().Unix(), 10)
	expire := stoe(expiration)
	if expire == -1 {
		body = fmt.Sprintf(`{"path":"%s","content":{"reason":"%s","source":"ohgi","timestamp":%s}}`, path, reason, now)
	} else {
		body = fmt.Sprintf(`{"path":"%s","content":{"reason":"%s","source":"ohgi","timestamp":%s},"expire":%d}`, path, reason, now, expire)
	}
	payload := strings.NewReader(body)

	_, status := postAPI("/stashes", payload)
	checkStatus(status)

	return httpStatus(status) + "\n"
}

func DeleteSilence(client string, check string) string {
	var path string

	if check == "" {
		path = "silence/" + client
	} else {
		path = "silence/" + client + "/" + check
	}

	_, status := deleteAPI("/stashes/" + path)
	checkStatus(status)

	return httpStatus(status) + "\n"
}
