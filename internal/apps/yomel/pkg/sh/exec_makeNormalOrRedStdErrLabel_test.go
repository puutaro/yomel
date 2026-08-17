package sh

import (
	"testing"
)

// TestMakeNormalOrRedStdErrLabel tests makeNormalOrRedStdErrLabel with various combinations of hasErr and isTerminal.
func TestMakeNormalOrRedStdErrLabel(t *testing.T) {
	tests := []struct {
		name       string
		hasErr     bool
		isTerminal bool
		want       string
	}{
		{
			name:       "Progress label when no error and terminal is false",
			hasErr:     false,
			isTerminal: false,
			want:       "Progress",
		},
		{
			name:       "Error label when error exists and terminal is false",
			hasErr:     true,
			isTerminal: false,
			want:       "Error",
		},
		{
			name:       "Progress label when no error and terminal is true",
			hasErr:     false,
			isTerminal: true,
			want:       "\x1b[4m\x1b[1mP\x1b[22mrogress\x1b[24m",
		},
		{
			name:       "Error label when error exists and terminal is true",
			hasErr:     true,
			isTerminal: true,
			want:       "\x1b[4m\x1b[1mE\x1b[22mrror\x1b[24m",
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
