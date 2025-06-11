package logger

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Verbose bool

	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

type Logger struct {
	cmd *cobra.Command
}

func NewLogger(cmd *cobra.Command) *Logger {
	return &Logger{cmd: cmd}
}

func (l *Logger) Debug(msg string) {
	if Verbose {
		l.log(fmt.Sprintf("DEBUG %s", msg))
	}
}

func (l *Logger) Info(msg string) {
	l.log(fmt.Sprintf("%sINFO %s%s", Cyan, msg, Reset))
}

func (l *Logger) Warn(msg string) {
	l.log(fmt.Sprintf("%sWARN %s%s", Yellow, msg, Reset))
}

func (l *Logger) Error(msg string) {
	l.logErr(fmt.Sprintf("%sERR %s%s", Red, msg, Reset))
}

func (l *Logger) Fatal(msg string) {
	l.logErr(fmt.Sprintf("%sFATAL %s%s", Red, msg, Reset))
	os.Exit(-1)
}

func (l *Logger) log(msg string) {
	l.cmd.Println(msg)
}

func (l *Logger) logErr(msg string) {
	l.cmd.PrintErrln(msg)
}
