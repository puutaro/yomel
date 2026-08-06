// Write direct above line for Comment on code
package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_capitalizeFirst verifies that capitalizeFirst correctly capitalizes the first letter of a string.
func Test_capitalizeFirst(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "should capitalize the first letter of a lowercase string",
			input: "hello world",
			want:  "Hello world",
		},
		{
			name:  "should keep the first letter capitalized if already uppercase",
			input: "Hello World",
			want:  "Hello World",
		},
		{
			name:  "should handle single character string correctly",
			input: "a",
			want:  "A",
		},
		{
			name:  "should return empty string when input is empty",
			input: "",
			want:  "",
		},
		{
			name:  "should handle Japanese or multibyte characters gracefully",
			input: "テスト",
			want:  "テスト",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := capitalizeFirst(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
