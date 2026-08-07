package sh

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_printTitleLog verifies that printTitleLog correctly formats and prints title logs to stderr.
func Test_printTitleLog(t *testing.T) {
	tests := []struct {
		name             string
		title            string
		cmdPipeline      string
		wantOutputSubstr []string
	}{
		{
			name:        "should print title log with title and pipeline command",
			title:       "initialize-environment",
			cmdPipeline: "go mod download",
			wantOutputSubstr: []string{
				"YOMEL-LOG-TITLE:",
				"Initialize-environment",
				"TotalCmd:",
				"go mod download",
			},
		},
		{
			name:             "should handle empty title gracefully",
			title:            "",
			cmdPipeline:      "some-cmd",
			wantOutputSubstr: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStderr := os.Stderr
			rErr, wErr, _ := os.Pipe()
			os.Stderr = wErr

			printTitleLog(tt.title, tt.cmdPipeline)

			wErr.Close()
			os.Stderr = oldStderr

			var bufErr bytes.Buffer
			_, _ = bufErr.ReadFrom(rErr)
			output := bufErr.String()

			if tt.title == "" {
				assert.Equal(t, "", output)
				return
			}

			for _, substr := range tt.wantOutputSubstr {
				assert.Contains(t, output, substr)
			}
		})
	}
}
