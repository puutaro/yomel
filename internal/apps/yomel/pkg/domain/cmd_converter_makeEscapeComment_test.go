package domain

import (
	"testing"
)

func Test_makeEscapeComment(t *testing.T) {
	tests := []struct {
		name       string
		comment    string
		isTerminal bool
		want       string
	}{
		{
			name:       "empty comment",
			comment:    "",
			isTerminal: false,
			want:       "",
		},
		{
			name:       "non-empty comment without terminal",
			comment:    "test comment",
			isTerminal: false,
			want:       "`# test comment`",
		},
		{
			name:       "non-empty comment with terminal",
			comment:    "test comment",
			isTerminal: true,
			want:       "\x1b[38;5;244m`# test comment`\x1b[39m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeEscapeComment(tt.comment, tt.isTerminal)
			if got != tt.want {
				t.Errorf("makeEscapeComment() = %q, want %q", got, tt.want)
			}
		})
	}
}
