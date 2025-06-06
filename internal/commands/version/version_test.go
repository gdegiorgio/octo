package version

import (
	"strings"
	"testing"

	"github.com/gdegiorgio/octo/utils"
)

func TestVersionCommand_PrintCorrectVersion(t *testing.T) {
	cmd := NewVersionCmd("testing")
	output := utils.RunCommandAndCaptureOutput(cmd)
	if strings.TrimSpace(output) != "🐙 testing" {
		t.Errorf("Expected testing but got %s\n", output)
	}
}
