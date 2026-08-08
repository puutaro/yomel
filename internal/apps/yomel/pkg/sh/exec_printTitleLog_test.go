// Test_printTitleLog verifies that printTitleLog correctly outputs the title.
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
		wantOutputSubstr []string
	}{
		{
			name:  "should print title log when title is provided",
			title: "Initialize-environment",
			wantOutputSubstr: []string{
				"Title",
				"Initialize-environment",
			},
		},
		{
			name:             "should output nothing when title is empty",
			title:            "",
			wantOutputSubstr: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var combinedLog bytes.Buffer
			yl := &yomelLog{}

			yl.printTitleLog(
				&combinedLog,
				tt.title,
			)

			output := stripTitleANSI(combinedLog.String())

			if len(tt.wantOutputSubstr) == 0 {
				assert.Empty(t, output)
				return
			}

			for _, substr := range tt.wantOutputSubstr {
				assert.Contains(t, output, substr)
			}
		})
	}
}
