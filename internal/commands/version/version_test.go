package version

import (
	"strings"
	"testing"

	"github.com/gdegiorgio/octo/utils"
)

func TestVersionCommand_PrintCorrectVersion(t *testing.T) {
	cmd := NewVersionCmd()
	output := utils.RunCommandAndCaptureOutput(cmd)
	if strings.TrimSpace(output) != Version {
		t.Errorf("Expected %s but got %s\n", Version, output)
	}
}
