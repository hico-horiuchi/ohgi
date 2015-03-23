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

func utoa(timestamp int64) string {
	format := "2006/01/02 15:04:05"
	return time.Unix(timestamp, 0).Format(format)
}

func stoa(arr []int, sep string) string {
	var result []byte

	for _, i := range arr {
		line := strconv.Itoa(i) + sep
		result = append(result, line...)
	}

	return string(result)
}

func stoe(expiration string) int64 {
	str := []byte(expiration)
	format := regexp.MustCompile("([0-9]+)([smhd])")
	group := format.FindSubmatch(str)
	var expire int64 = -1

	if len(group) == 3 {
		num, _ := strconv.ParseInt(string(group[1]), 10, 0)
		switch string(group[2]) {
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

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkStatus(status int) {
	if status >= 300 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}
}

func httpStatus(status int) string {
	switch status {
	case 200:
		return strconv.Itoa(status) + " OK"
	case 201:
		return strconv.Itoa(status) + " Created"
	case 202:
		return strconv.Itoa(status) + " Accepted"
	case 204:
		return strconv.Itoa(status) + " No Content"
	case 400:
		return strconv.Itoa(status) + " Bad Request"
	case 401:
		return strconv.Itoa(status) + " Unauthorized"
	case 404:
		return strconv.Itoa(status) + " Not Found"
	case 500:
		return strconv.Itoa(status) + " Internal Server Error"
	case 503:
		return strconv.Itoa(status) + " Service Unavailable"
	}
	return ""
}

func bold(str string) string {
	if !EscapeSequence {
		return str
	}

	return "\x1b[1m" + str + "\x1b[0m"
}

func statusFg(status int) string {
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

func statusBg(status int) string {
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
	}
	return "\x1b[47m \x1b[0m "
}

func fillSpace(str string, max int) string {
	padding := 2
	width := max - padding
	length := len(str)
	if length > width {
		return str[0:width] + strings.Repeat(" ", padding)
	} else {
		return str + strings.Repeat(" ", max-length)
	}
}
