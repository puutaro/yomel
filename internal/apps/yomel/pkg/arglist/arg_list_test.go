package arglist

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Gen verifies that Gen correctly returns command line arguments from os.Args.
func Test_Gen(t *testing.T) {
	tests := []struct {
		name     string
		mockArgs []string
		want     []string
	}{
		{
			name:     "should return command line arguments successfully",
			mockArgs: []string{"yomel", "stage", "test1", "-cmd", "echo"},
			want:     []string{"stage", "test1", "-cmd", "echo"},
		},
		{
			name:     "should handle empty arguments slice gracefully",
			mockArgs: []string{"yomel"},
			want:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = tt.mockArgs

			got := Gen()
			assert.Equal(t, tt.want, got)
		})
	}
}
