package ohgi

import (
	"os"

	isatty "github.com/mattn/go-isatty"
)

var escapeSequence bool

func Initialize() {
	escapeSequence = isTerminal()
}

func isTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}
