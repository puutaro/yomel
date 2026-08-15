package sh

import (
	"testing"
)

// TestMakeNormalOrRedStdErrLabel tests makeNormalOrRedStdErrLabel with various conditions.
func TestMakeNormalOrRedStdErrLabel(t *testing.T) {
	tests := []struct {
		name       string
		hasErr     bool
		isTerminal bool
		want       string
	}{
		{
			name:       "Progress without terminal",
			hasErr:     false,
			isTerminal: false,
			want:       "Progress",
		},
		{
			name:       "Progress with terminal",
			hasErr:     false,
			isTerminal: true,
			want:       "\x1b[4m\x1b[1mP\x1b[22mrogress\x1b[24m",
		},
		{
			name:       "Error without terminal",
			hasErr:     true,
			isTerminal: false,
			want:       "\x1b[31mError\x1b[39m",
		},
		{
			name:       "Error with terminal",
			hasErr:     true,
			isTerminal: true,
			want:       "\x1b[31m\x1b[4m\x1b[1mE\x1b[22mrror\x1b[24m\x1b[39m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeNormalOrRedStdErrLabel(tt.hasErr, tt.isTerminal)
			if got != tt.want {
				t.Errorf("makeNormalOrRedStdErrLabel() = %q, want %q", got, tt.want)
			}
		})
	}
}
