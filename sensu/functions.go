package sensu

import (
	"strconv"
	"strings"
	"time"
)

func utoa(timestamp int64) string {
	format := "2006/01/02 15:04:05"
	return time.Unix(timestamp, 0).Format(format)
}

func bold(str string) string {
	return "\x1b[1m" + str + "\x1b[0m"
}

func httpStatus(status int) string {
	switch status {
	case 401:
		return strconv.Itoa(status) + " Unauthorized"
	case 404:
		return strconv.Itoa(status) + " Not Found"
	case 500:
		return strconv.Itoa(status) + " Internal Server Error"
	}
	return ""
}

func statusFg(status int) string {
	switch status {
	case 0:
		return "\x1b[32mOK\x1b[0m "
	case 1:
		return "\x1b[33mWARNING\x1b[0m "
	case 2:
		return "\x1b[31mCRITICAL\x1b[0m "
	}
	return "\x1b[37mUNKNOWN\x1b[0m "
}

func statusBg(status int) string {
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
