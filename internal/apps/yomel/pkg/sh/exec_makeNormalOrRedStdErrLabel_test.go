package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeNormalOrRedStdErrLabel(t *testing.T) {
	tests := []struct {
		name   string
		hasErr bool
		want   string
	}{
		{
			name:   "should return red error label when hasErr is true",
			hasErr: true,
			want:   "#\x1b[31m error:\x1b[0m\n",
		},
		{
			name:   "should return normal progress label when hasErr is false",
			hasErr: false,
			want:   "# progress:\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeNormalOrRedStdErrLabel(tt.hasErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
