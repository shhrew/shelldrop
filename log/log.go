package log

import (
	"fmt"
	"os"
)

func Fatalf(format string, v ...any) {
	fmt.Println("[!]", fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatal(message string) {
	Fatalf("%s", message)
}

func Errorf(format string, v ...any) {
	fmt.Println("[!]", fmt.Sprintf(format, v...))
}

func Error(message string) {
	Errorf("%s", message)
}

func Infof(format string, v ...any) {
	fmt.Println("[*]", fmt.Sprintf(format, v...))
}

func Info(message string) {
	Infof("%s", message)
}

func Warnf(format string, v ...any) {
	fmt.Println("[?]", fmt.Sprintf(format, v...))
}

func Warn(message string) {
	Warnf("%s", message)
}
