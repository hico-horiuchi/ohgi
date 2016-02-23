package ohgi

import "github.com/fatih/color"

func fgColor(str string, status int) string {
	if !escapeSequence {
		return str
	}

	switch status {
	case 0:
		return color.GreenString(str)
	case 1:
		return color.YellowString(str)
	case 2:
		return color.RedString(str)
	default:
		return color.WhiteString(str)
	}
}

func bgColor(str string, status int) string {
	if !escapeSequence {
		return str
	}

	switch status {
	case 0:
		return color.New(color.BgGreen).SprintFunc()(str)
	case 1:
		return color.New(color.BgYellow).SprintFunc()(str)
	case 2:
		return color.New(color.BgRed).SprintFunc()(str)
	default:
		return color.New(color.BgWhite).SprintFunc()(str)
	}
}

func bold(str string) string {
	if !escapeSequence {
		return str
	}

	return color.New(color.Bold).SprintFunc()(str)
}
