package color

import (
	"bytes"
	"fmt"
)

const (
	Reset = "Reset"

	Bright     = "Bright"
	Dim        = "Dim"
	Underscore = "Underscore"
	Blink      = "Blink"
	Reverse    = "Reverse"
	Hidden     = "Hidden"

	FgBlack   = "FgBlack"
	FgRed     = "FgRed"
	FgGreen   = "FgGreen"
	FgYellow  = "FgYellow"
	FgBlue    = "FgBlue"
	FgMagenta = "FgMagenta"
	FgCyan    = "FgCyan"
	FgWhite   = "FgWhite"

	BgBlack   = "BgBlack"
	BgRed     = "BgRed"
	BgGreen   = "BgGreen"
	BgYellow  = "BgYellow"
	BgBlue    = "BgBlue"
	BgMagenta = "BgMagenta"
	BgCyan    = "BgCyan"
	BgWhite   = "BgWhite"
)

var (
	colors = map[string]string{
		"Reset": "\x1b[0m",

		"Bright":     "\x1b[1m",
		"Dim":        "\x1b[2m",
		"Underscore": "\x1b[4m",
		"Blink":      "\x1b[5m",
		"Reverse":    "\x1b[7m",
		"Hidden":     "\x1b[8m",

		"FgBlack":   "\x1b[30m",
		"FgRed":     "\x1b[31m",
		"FgGreen":   "\x1b[32m",
		"FgYellow":  "\x1b[33m",
		"FgBlue":    "\x1b[34m",
		"FgMagenta": "\x1b[35m",
		"FgCyan":    "\x1b[36m",
		"FgWhite":   "\x1b[37m",

		"BgBlack":   "\x1b[40m",
		"BgRed":     "\x1b[41m",
		"BgGreen":   "\x1b[42m",
		"BgYellow":  "\x1b[43m",
		"BgBlue":    "\x1b[44m",
		"BgMagenta": "\x1b[45m",
		"BgCyan":    "\x1b[46m",
		"BgWhite":   "\x1b[47m",
	}
)

// Colorize provides output of the string with color
func Colorize(colorsSlice []string, format string, a ...interface{}) (s string) {
	buf := new(bytes.Buffer)
	for _, color := range colorsSlice {
		buf.WriteString(colors[color])
	}

	if len(a) == 0 {
		buf.WriteString(format)
	} else {
		buf.WriteString(fmt.Sprintf(format, a...))
	}
	buf.WriteString(colors[Reset])
	s = buf.String()
	return
}

// Blue returns string on red color
func Blue(format string, a ...interface{}) string {
	return colored(FgBlue, format, a...)
}

// Red returns string on red color
func Red(format string, a ...interface{}) string {
	return colored(FgRed, format, a...)
}

// Green returns string on green color
func Green(format string, a ...interface{}) string {
	return colored(FgGreen, format, a...)
}

// Yellow returns string on yellow color
func Yellow(format string, a ...interface{}) string {
	return colored(FgYellow, format, a...)
}

func colored(color string, format string, a ...interface{}) string {
	if len(a) == 0 {
		return Colorize([]string{color}, format)
	}

	return Colorize([]string{color}, fmt.Sprintf(format, a...))
}
