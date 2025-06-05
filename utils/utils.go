package utils

import (
	"bytes"

	"github.com/spf13/cobra"
)

func RunCommandAndCaptureOutput(cmd *cobra.Command) (string){

	var buf bytes.Buffer

	cmd.SetOut(&buf)
	cmd.SetErr(&buf)

	_ = cmd.Execute()

	return buf.String()
}
