// Test_printYomelLogStartHolder verifies that printYomelLogStartHolder correctly outputs the formatted start holder log with timestamp.
package sh

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_printYomelLogStartHolder(t *testing.T) {
	tests := []struct {
		name      string
		startTime time.Time
		want      string
	}{
		{
			name:      "should print yomel log start holder with formatted timestamp",
			startTime: time.Date(2026, 6, 15, 12, 30, 45, 123456000, time.UTC),
			want:      "\n\x1b[4m\x1b[1mYomel-log_2026/06/15-12:30:45.123456\x1b[22m\x1b[24m\n",
		},
		{
			name:      "should handle zero time correctly",
			startTime: time.Time{},
			want:      "\n\x1b[4m\x1b[1mYomel-log_0001/01/01-00:00:00.000000\x1b[22m\x1b[24m\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			yl := &yomelLog{
				startTime: tt.startTime,
			}

			yl.printYomelLogStartHolder(&buf)

			assert.Equal(t, tt.want, buf.String())
		})
	}
}
