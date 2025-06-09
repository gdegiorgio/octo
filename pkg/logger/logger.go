package logger

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Reset = "\033[0m"
 	Red = "\033[31m"
	Yellow = "\033[33m"
	Cyan = "\033[36m"
	Verbose bool
)


func Debug(cmd *cobra.Command, msg string){
	if(Verbose){
		log(cmd, fmt.Sprintf(`DEBUG %s`, msg))
	}
}

func Info(cmd *cobra.Command, msg string){
	log(cmd, fmt.Sprintf(`%s INFO %s %s`, Cyan, msg, Reset))
}

func Warn(cmd *cobra.Command, msg string){
	log(cmd, fmt.Sprintf(`%s WARN %s %s`,Yellow, msg, Reset))
}

func Error(cmd *cobra.Command, msg string) error{
	logErr(cmd, fmt.Sprintf(`%s ERR %s %s`, Red, msg, Reset))
	return errors.New(msg)
}

func Fatal(cmd *cobra.Command, msg string){
	logErr(cmd, fmt.Sprintf(`%s FATAL %s %s`, Red, msg, Reset))
	os.Exit(-1)
}

func log(cmd *cobra.Command, msg string){
	cmd.Println(msg)
}

func logErr(cmd *cobra.Command, msg string){
	cmd.PrintErrln(msg)
}
