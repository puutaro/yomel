// Test_printTitleLog verifies that printTitleLog correctly outputs the title and pipeline command log.
package sh

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// stripTitleANSI removes ANSI escape sequences (colors, underlines, bold, etc.) from the string.
func stripTitleANSI(str string) string {
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[mK]`)
	return ansiRegex.ReplaceAllString(str, "")
}

func Test_printTitleLog(t *testing.T) {
	tests := []struct {
		name             string
		title            string
		totalPipeCmdStr  string
		wantOutputSubstr []string
	}{
		{
			name:            "should_print_title_log_with_title_and_pipeline_command",
			title:           "Initialize-environment",
			totalPipeCmdStr: "go mod download",
			wantOutputSubstr: []string{
				"Yomel-log-title",
				"Initialize-environment",
				"Total-cmd",
				"go mod download",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var combinedLog bytes.Buffer

			printTitleLog(
				&combinedLog,
				tt.title,
				tt.totalPipeCmdStr,
			)

			output := stripTitleANSI(combinedLog.String())

			for _, substr := range tt.wantOutputSubstr {
				assert.Contains(t, output, substr)
			}
		})
	}
}
