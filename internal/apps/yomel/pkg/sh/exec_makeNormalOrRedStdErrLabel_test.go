// Write direct above line for Comment on code
package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_makeNormalOrRedStdErrLabel verifies that makeNormalOrRedStdErrLabel returns the correct decorated progress or error label based on the hasErr flag.
func Test_makeNormalOrRedStdErrLabel(t *testing.T) {
	tests := []struct {
		name   string
		hasErr bool
		want   string
	}{
		{
			name:   "should return red error label when hasErr is true",
			hasErr: true,
			want:   "\x1b[31m\x1b[4m\x1b[1mE\x1b[22mrror\x1b[24m\x1b[39m\n",
		},
		{
			name:   "should return normal progress label when hasErr is false",
			hasErr: false,
			want:   "\x1b[4m\x1b[1mP\x1b[22mrogress\x1b[24m\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeNormalOrRedStdErrLabel(tt.hasErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
