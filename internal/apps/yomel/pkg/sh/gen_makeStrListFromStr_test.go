package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_makeStrListFromStr verifies that makeStrListFromStr correctly converts a string into a string slice or returns nil when empty.
func Test_makeStrListFromStr(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "should return string slice with single element when input string is not empty",
			input: "test-string",
			want:  []string{"test-string"},
		},
		{
			name:  "should return nil when input string is empty",
			input: "",
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeStrListFromStr(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
