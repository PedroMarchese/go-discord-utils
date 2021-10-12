package logger

import "github.com/fatih/color"

var (
	info  = color.New(color.FgCyan).PrintlnFunc()
	err   = color.New(color.FgRed).PrintlnFunc()
	fatal = color.New(color.BgRed).PrintlnFunc()
)

func Info(contents interface{}) {
	info("[→]", contents)
}

func Error(contents interface{}) {
	err("[!]", contents)
}

func Fatal(contents interface{}) {
	fatal("[X]", contents)
}
