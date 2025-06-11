package logger

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Long:  "Used for testing",
	Short: "Used for testing",
}

var logger = NewLogger(testCmd)

func TestDebug_NoVerbose(t *testing.T) {
	var buf bytes.Buffer
	testCmd.SetOut(&buf)
	Verbose = false

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		logger.Debug("Test")
	}

	_ = testCmd.Execute()
	expected := ""
	if buf.String() != expected {
		t.Errorf("Expected %q, found %q", expected, buf.String())
	}
}

func TestDebug_Verbose(t *testing.T) {
	var buf bytes.Buffer
	testCmd.SetOut(&buf)
	Verbose = true

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		logger.Debug("Test")
	}

	_ = testCmd.Execute()
	expected := "DEBUG Test\n"
	if buf.String() != expected {
		t.Errorf("Expected %q, found %q", expected, buf.String())
	}
}

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	testCmd.SetOut(&buf)
	Verbose = true

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		logger.Info("Test")
	}

	_ = testCmd.Execute()
	expected := fmt.Sprintf("%sINFO Test%s\n", Cyan, Reset)
	if buf.String() != expected {
		t.Errorf("Expected %q, found %q", expected, buf.String())
	}
}

func TestWarn(t *testing.T) {
	var buf bytes.Buffer
	testCmd.SetOut(&buf)
	Verbose = true

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		logger.Warn("Test")
	}

	_ = testCmd.Execute()
	expected := fmt.Sprintf("%sWARN Test%s\n", Yellow, Reset)
	if buf.String() != expected {
		t.Errorf("Expected %q, found %q", expected, buf.String())
	}
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	testCmd.SetErr(&buf)
	Verbose = true

	testCmd.Run = func(cmd *cobra.Command, args []string) {
		logger.Error("Test")
	}

	_ = testCmd.Execute()
	expected := fmt.Sprintf("%sERR Test%s\n", Red, Reset)
	if buf.String() != expected {
		t.Errorf("Expected %q, found %q", expected, buf.String())
	}
}
