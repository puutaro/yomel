// Test_convertTimeStampStr verifies that convertTimeStampStr correctly formats a time.Time into the expected timestamp string format.
package sh

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_convertTimeStampStr(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "should format time correctly into standard timestamp string",
			input: time.Date(2026, 6, 15, 12, 30, 45, 123456000, time.UTC),
			want:  "2026/06/15-12:30:45.123456",
		},
		{
			name:  "should handle zero time correctly",
			input: time.Time{},
			want:  "0001/01/01-00:00:00.000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertTimeStampStr(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
