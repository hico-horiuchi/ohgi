package ohgi

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/hico-horiuchi/uitable"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newUitable(wrap ...bool) *uitable.Table {
	table := uitable.New()
	table.MaxColWidth = uint(terminalWidth / 2)
	if wrap != nil && len(wrap) > 0 && wrap[0] {
		table.Wrap = true
	}

	return table
}

func formatOutput(output string) string {
	return strings.TrimSpace(strings.Replace(output, "\n", " ", -1))
}

func indicateStatus(status int) string {
	if !escapeSequence {
		return strconv.Itoa(status)
	}

	return bgColor(" ", status)
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
