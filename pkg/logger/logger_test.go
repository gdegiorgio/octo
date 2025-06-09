package logger

import (
	"fmt"
	"io"
	"testing"

	"github.com/gdegiorgio/octo/utils"
	"github.com/spf13/cobra"
)





var testCmd = &cobra.Command {
	Use:  "test",
	Long: "Used for testing",
	Short: "Used for testing",
}


func TestDebug_NoVerbose(t *testing.T){

	var buf io.Writer

	testCmd.SetOut(buf)
	Verbose = false

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		Debug(testCmd, "Test")
	}

	out := utils.RunCommandAndCaptureOutput(testCmd)
	expected := ""
	if out != expected{
		t.Errorf("Expected %s, found %s", expected, out )
	}
}


func TestDebug_Verbose(t *testing.T){

	var buf io.Writer

	testCmd.SetOut(buf)
	Verbose = true

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		Debug(testCmd, "Test")
	}

	out := utils.RunCommandAndCaptureOutput(testCmd)
	expected := "DEBUG Test\n"
	if out != expected{
		t.Errorf("Expected %s, found %s", expected, out )
	}

}
func TestInfo(t *testing.T){
	var buf io.Writer

	testCmd.SetOut(buf)
	Verbose = true

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		Info(testCmd, "Test")
	}

	out := utils.RunCommandAndCaptureOutput(testCmd)
	expected := fmt.Sprintf("%s INFO Test %s\n", Cyan, Reset)
	if out != expected{
		t.Errorf("Expected %s, found %s", expected, out )
	}
}
func TestWarn(t *testing.T){

	var buf io.Writer

	testCmd.SetOut(buf)
	Verbose = true

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		Warn(testCmd, "Test")
	}

	out := utils.RunCommandAndCaptureOutput(testCmd)
	expected := fmt.Sprintf("%s WARN Test %s\n", Yellow, Reset)

	if out != expected{
		t.Errorf("Expected %s, found %s", expected, out )
	}
}

func TestError(t *testing.T){

	var buf io.Writer


	testCmd.SetOut(buf)
	Verbose = true

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		Error(testCmd, "Test")
	}

	out := utils.RunCommandAndCaptureOutput(testCmd)
	expected := fmt.Sprintf("%s ERR Test %s\n", Red, Reset)

	if out != expected{
		t.Errorf("Expected %s, found %s", expected, out )
	}
}
