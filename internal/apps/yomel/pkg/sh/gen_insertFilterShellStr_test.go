// Write direct above line for Comment on code
package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_insertFilterShellStr verifies that insertFilterShellStr returns the logFilter when it is not empty, and falls back to the globalFilter otherwise.
func Test_insertFilterShellStr(t *testing.T) {
	tests := []struct {
		name         string
		globalFilter string
		logFilter    string
		want         string
	}{
		{
			name:         "should return logFilter when logFilter is not empty",
			globalFilter: "global-filter",
			logFilter:    "stage-filter",
			want:         "stage-filter",
		},
		{
			name:         "should return globalFilter when logFilter is empty",
			globalFilter: "global-filter",
			logFilter:    "",
			want:         "global-filter",
		},
		{
			name:         "should return empty string when both filters are empty",
			globalFilter: "",
			logFilter:    "",
			want:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertFilterShellStr(tt.globalFilter, tt.logFilter)
			assert.Equal(t, tt.want, got)
		})
	}
}
