package ohgi

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

func indicateStatus(status int) string {
	if !escapeSequence {
		return strconv.Itoa(status) + " "
	}

	return bgColor(" ", status) + " "
}

func colorStatus(status int) string {
	switch status {
	case 0:
		return fgColor("OK", 0)
	case 1:
		return fgColor("WARNIG", 1)
	case 2:
		return fgColor("CRITICAL", 2)
	default:
		return fgColor("UNKNOWN", 3)
	}
}

func colorHistory(history string) string {
	var format *regexp.Regexp

	format = regexp.MustCompile("(^| )(0)(|,)")
	history = format.ReplaceAllString(history, fgColor("$1$2$3", 0))
	format = regexp.MustCompile("(^| )(1)(|,)")
	history = format.ReplaceAllString(history, fgColor("$1$2$3", 1))
	format = regexp.MustCompile("(^| )(2)(|,)")
	history = format.ReplaceAllString(history, fgColor("$1$2$3", 2))

	return history
}
