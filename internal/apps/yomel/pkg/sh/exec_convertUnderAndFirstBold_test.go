// Write direct above line for Comment on code
package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_convertUnderAndFirstBold verifies that convertUnderAndFirstBold correctly adds underline and bold formatting to the first character of the string.
func Test_convertUnderAndFirstBold(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "should format first character with underline and bold when string is provided",
			input: "Stage",
			want:  "\x1b[4m\x1b[1mS\x1b[22mtage\x1b[24m",
		},
		{
			name:  "should handle single character string correctly",
			input: "C",
			want:  "\x1b[4m\x1b[1mC\x1b[22m\x1b[24m",
		},
		{
			name:  "should return empty string when input is empty",
			input: "",
			want:  "",
		},
		{
			name:  "should handle Japanese characters gracefully",
			input: "テスト",
			want:  "\x1b[4m\x1b[1mテ\x1b[22mスト\x1b[24m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertUnderAndFirstBold(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
