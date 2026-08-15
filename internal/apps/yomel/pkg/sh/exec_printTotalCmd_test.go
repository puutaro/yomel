package sh

import (
	"bytes"
	"testing"
)

// Test printTotalCmd with various terminal and color configurations
func TestPrintTotalCmd(t *testing.T) {
	tests := []struct {
		name            string
		isTerminal      bool
		totalPipeCmdStr string
		titleStartColor string
		foregroundColor string
		commentColor    string
	}{
		{
			name:            "Terminal disabled case",
			isTerminal:      false,
			totalPipeCmdStr: "echo 'test' \\\n| cat",
			titleStartColor: "",
			foregroundColor: "",
			commentColor:    "",
		},
		{
			name:            "Terminal enabled with color and comments",
			isTerminal:      true,
			totalPipeCmdStr: "echo 'test' SIED_COmMetN_BAlNk # side comment \\\n| cat",
			titleStartColor: "\x1b[48;5;52m",
			foregroundColor: "\x1b[38;5;244m",
			commentColor:    "\x1b[38;5;208m",
		},
		{
			name:            "Empty command string case",
			isTerminal:      true,
			totalPipeCmdStr: "",
			titleStartColor: "\x1b[48;5;52m",
			foregroundColor: "\x1b[38;5;244m",
			commentColor:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize yomelLog for testing printTotalCmd
			yl := yomelLog{
				yomelInfo: YomelInfo{
					Title:                       "Sample Title",
					ForegroundColor:             tt.foregroundColor,
					TitleCommentForegroundColor: tt.commentColor,
				},
				isTerminal: tt.isTerminal,
			}

			var buf bytes.Buffer
			// Execute printTotalCmd method
			yl.printTotalCmd(&buf, tt.totalPipeCmdStr, tt.titleStartColor)

			if buf.Len() == 0 && tt.totalPipeCmdStr != "" {
				t.Errorf("printTotalCmd() produced unexpected empty output for %s", tt.name)
			}
		})
	}
}
