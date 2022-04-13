package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

const (
	DebugLevel = iota
	InfoLevel
)

var (
	Level = InfoLevel

	debugPrefix = color.New(color.FgCyan).Sprint("DEBUG ")
	warnPrefix  = color.New(color.FgYellow).Sprint("WARN ")
	errorPrefix = color.New(color.FgRed).Sprint("ERROR ")
)

func Info(format string, args ...interface{}) {
	if Level > InfoLevel {
		return
	}
	fmt.Printf(format+"\n", args...)
}

func Debug(format string, args ...interface{}) {
	if Level > DebugLevel {
		return
	}
	fmt.Printf(debugPrefix+format+"\n", args...)
}

func Warn(format string, args ...interface{}) {
	fmt.Printf(warnPrefix+format+"\n", args...)
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, errorPrefix+"%s\n", err.Error())
}
