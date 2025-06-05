package version

import (
	"os"
	"strings"
	"testing"

	"github.com/gdegiorgio/octo/utils"
)

func TestVersionCommand_MissingEnvVar(t *testing.T) {
	cmd := NewVersionCmd()
	output := utils.RunCommandAndCaptureOutput(cmd)
	if strings.TrimSpace(output) != "Unknown Version" {
		t.Errorf("Expected \"Unkown Version\" but got %s\n", output)
	}
}

func TestVersionCommand_EnvVarSet(t *testing.T) {
	err := os.Setenv("OCTO_CLI_VERSION", "v1.0.0")
	if err != nil {
		t.Errorf("Could not set env for tests : %v+", err)
	}
	cmd := NewVersionCmd()
	output := utils.RunCommandAndCaptureOutput(cmd)
	if strings.TrimSpace(output) != "v1.0.0" {
		t.Errorf("Expected \"v1.0.0\" but got %s\n", output)
	}
}
