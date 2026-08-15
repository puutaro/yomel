package sh

import (
	"bytes"
	"testing"
)

func Test_printTitleLog(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		isTerminal bool
		fgColor    string
		bgCol      string
		want       string
	}{
		{
			name:       "empty title",
			title:      "",
			isTerminal: false,
			want:       "",
		},
		{
			name:       "non-terminal title",
			title:      "test title",
			isTerminal: false,
			want:       "\nTitle\nTest title",
		},
		{
			name:       "terminal title without bg",
			title:      "my pipeline",
			isTerminal: true,
			fgColor:    "\x1b[31m",
			bgCol:      "",
			want:       "\n\x1b[31m\x1b[4m\x1b[1mT\x1b[22mitle\x1b[24m\x1b[39m\n\x1b[31mMy pipeline\x1b[39m",
		},
		{
			name:       "terminal title with bg",
			title:      "colored title",
			isTerminal: true,
			fgColor:    "\x1b[32m",
			bgCol:      "\x1b[48;5;22m",
			want:       "\n\x1b[48;5;22m\x1b[32m\x1b[4m\x1b[1mT\x1b[22mitle\x1b[24m\x1b[39m\x1b[49m\n\x1b[32mColored title\x1b[39m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			yl := yomelLog{
				yomelInfo: YomelInfo{
					Title:           tt.title,
					ForegroundColor: tt.fgColor,
				},
				isTerminal: tt.isTerminal,
			}
			yl.printTitleLog(&buf, tt.title, tt.bgCol)
			got := buf.String()
			if got != tt.want {
				t.Errorf("printTitleLog() = %q, want %q", got, tt.want)
			}
		})
	}
}
