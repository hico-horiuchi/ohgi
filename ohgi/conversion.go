package ohgi

import (
	"regexp"
	"strconv"
	"time"
)

func utoa(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006/01/02 15:04:05")
}

func stoa(arr []int, sep string) string {
	var print []byte

	for _, i := range arr {
		print = append(print, (strconv.Itoa(i) + sep)...)
	}

	return string(print)
}

func stoe(expiration string) int64 {
	var expire int64 = -1

	str := []byte(expiration)
	format := regexp.MustCompile("([0-9]+)([smhd])")
	matches := format.FindSubmatch(str)

	if len(matches) == 3 {
		num, err := strconv.ParseInt(string(matches[1]), 10, 0)
		if err != nil {
			return expire
		}

		switch string(matches[2]) {
		case "s":
			expire = num
		case "m":
			expire = num * int64(time.Minute) / int64(time.Second)
		case "h":
			expire = num * int64(time.Hour) / int64(time.Second)
		case "d":
			expire = num * int64(time.Hour) * 24 / int64(time.Second)
		}
	}

	return expire
}
