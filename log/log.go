package log

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Fatalf(format string, v ...any) {
	output(color.FgHiRed, "[!]", fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatal(message string) {
	Fatalf("%s", message)
}

func Errorf(format string, v ...any) {
	output(color.FgRed, "[!]", fmt.Sprintf(format, v...))
}

func Error(message string) {
	Errorf("%s", message)
}

func Infof(format string, v ...any) {
	output(color.FgBlue, "[+]", fmt.Sprintf(format, v...))
}

func Info(message string) {
	Infof("%s", message)
}

func Warnf(format string, v ...any) {
	output(color.FgYellow, "[?]", fmt.Sprintf(format, v...))
}

func Warn(message string) {
	Warnf("%s", message)
}

func Successf(format string, v ...any) {
	output(color.FgHiGreen, "[#]", fmt.Sprintf(format, v...))
}

func Success(message string) {
	Successf("%s", message)
}

func output(colorAttribute color.Attribute, prefix, message string) {
	prefixColor := color.New(colorAttribute).SprintFunc()
	fmt.Println(fmt.Sprintf("%s %s", prefixColor(prefix), message))
}
