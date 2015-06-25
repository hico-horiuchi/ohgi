package ohgi

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var EscapeSequence bool = true

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func utoa(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006/01/02 15:04:05")
}

func stoa(arr []int, sep string) string {
	var result []byte

	for _, i := range arr {
		result = append(result, (strconv.Itoa(i) + sep)...)
	}

	return string(result)
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

func fillSpace(str string, max int) string {
	length := len(str)
	padding := 2
	width := max - padding

	if length > width {
		return str[0:width] + strings.Repeat(" ", padding)
	} else {
		return str + strings.Repeat(" ", max-length)
	}
}

func bold(str string) string {
	if !EscapeSequence {
		return str
	}

	return "\x1b[1m" + str + "\x1b[0m"
}

func indicateStatus(status int) string {
	if !EscapeSequence {
		return "  "
	}

	switch status {
	case 0:
		return "\x1b[42m \x1b[0m "
	case 1:
		return "\x1b[43m \x1b[0m "
	case 2:
		return "\x1b[41m \x1b[0m "
	default:
		return "\x1b[47m \x1b[0m "
	}
}

func paintStatus(status int) string {
	if !EscapeSequence {
		switch status {
		case 0:
			return "OK "
		case 1:
			return "WARNING "
		case 2:
			return "CRITICAL "
		default:
			return "UNKNOWN "
		}
	}

	switch status {
	case 0:
		return "\x1b[32mOK\x1b[0m "
	case 1:
		return "\x1b[33mWARNING\x1b[0m "
	case 2:
		return "\x1b[31mCRITICAL\x1b[0m "
	default:
		return "\x1b[37mUNKNOWN\x1b[0m "
	}
}

func paintHistory(history string) string {
	var format *regexp.Regexp

	if !EscapeSequence {
		return history
	}

	format = regexp.MustCompile("(^| )(0)(|,)")
	history = format.ReplaceAllString(history, "\x1b[32m$1$2$3\x1b[0m")
	format = regexp.MustCompile("(^| )(1)(|,)")
	history = format.ReplaceAllString(history, "\x1b[33m$1$2$3\x1b[0m")
	format = regexp.MustCompile("(^| )(2)(|,)")
	history = format.ReplaceAllString(history, "\x1b[31m$1$2$3\x1b[0m")

	return history
}
