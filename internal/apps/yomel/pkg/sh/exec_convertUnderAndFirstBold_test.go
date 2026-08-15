// internal/apps/yomel/pkg/sh/exec_convertUnderAndFirstBold_test.go
package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_convertUnderAndFirstBold verifies that convertUnderAndFirstBold correctly adds underline and bold formatting to the first character when isTerminal is true, and returns plain text when isTerminal is false.
func Test_convertUnderAndFirstBold(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		isTerminal bool
		want       string
	}{
		{
			name:       "should format first character with underline and bold when isTerminal is true",
			input:      "Stage",
			isTerminal: true,
			want:       "\x1b[4m\x1b[1mS\x1b[22mtage\x1b[24m",
		},
		{
			name:       "should handle single character string correctly when isTerminal is true",
			input:      "C",
			isTerminal: true,
			want:       "\x1b[4m\x1b[1mC\x1b[22m\x1b[24m",
		},
		{
			name:       "should return plain text when isTerminal is false",
			input:      "Stage",
			isTerminal: false,
			want:       "Stage",
		},
		{
			name:       "should return empty string when input is empty regardless of isTerminal",
			input:      "",
			isTerminal: true,
			want:       "",
		},
		{
			name:       "should handle Japanese characters gracefully when isTerminal is true",
			input:      "テスト",
			isTerminal: true,
			want:       "\x1b[4m\x1b[1mテ\x1b[22mスト\x1b[24m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertUnderAndFirstBold(tt.input, tt.isTerminal)
			assert.Equal(t, tt.want, got)
		})
	}
}
