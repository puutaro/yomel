package argtables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_removePrefix verifies that removePrefix correctly strips the specified prefix from the input string.
func Test_removePrefix(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		prefix string
		want   string
	}{
		{
			name:   "should remove prefix when string starts with prefix",
			str:    "--opt",
			prefix: "--",
			want:   "opt",
		},
		{
			name:   "should remove prefix for single character prefix",
			str:    "-cmd",
			prefix: "-",
			want:   "cmd",
		},
		{
			name:   "should return original string without changes when prefix does not match",
			str:    "opt",
			prefix: "--",
			want:   "opt",
		},
		{
			name:   "should handle empty string correctly",
			str:    "",
			prefix: "--",
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removePrefix(tt.str, tt.prefix)
			assert.Equal(t, tt.want, got)
		})
	}
}
