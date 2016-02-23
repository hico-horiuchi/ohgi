package ohgi

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	isatty "github.com/mattn/go-isatty"
)

var escapeSequence bool
var terminalWidth int

func Initialize() {
	escapeSequence = isTerminal()
	terminalWidth = getTerminalWidth()
}

func isTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	checkError(err)

	arr := strings.Split(strings.TrimRight(string(out), "\n"), " ")
	if len(arr) != 2 {
		fmt.Println("Can not get the terminal width with stty")
		os.Exit(1)
	}

	w, err := strconv.ParseInt(arr[1], 10, 0)
	checkError(err)

	return int(w)
}
