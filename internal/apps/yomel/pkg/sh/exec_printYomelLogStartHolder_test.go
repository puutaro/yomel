package sh

import (
	"bytes"
	"testing"
	"time"
)

func Test_yomelLog_printYomelLogStartHolder(t *testing.T) {
	fixedTime := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name                      string
		isTerminal                bool
		foregroundColor           string
		titleStartBackgroundColor string
		expectedContains          string
	}{
		{
			name:                      "Non-terminal mode",
			isTerminal:                false,
			foregroundColor:           "\x1b[31m",
			titleStartBackgroundColor: "",
			expectedContains:          "Yomel-log 2026/01/01-12:00:00.000000",
		},
		{
			name:                      "Terminal mode with background color",
			isTerminal:                true,
			foregroundColor:           "\x1b[31m",
			titleStartBackgroundColor: "\x1b[48;5;52m",
			expectedContains:          "Yomel-log",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yl := &yomelLog{
				yomelInfo: YomelInfo{
					ForegroundColor: tt.foregroundColor,
				},
				startTime:  fixedTime,
				isTerminal: tt.isTerminal,
			}

			var buf bytes.Buffer
			yl.printYomelLogStartHolder(&buf, tt.titleStartBackgroundColor)

			got := buf.String()
			if got == "" {
				t.Errorf("printYomelLogStartHolder() output is empty")
			}
		})
	}
}
