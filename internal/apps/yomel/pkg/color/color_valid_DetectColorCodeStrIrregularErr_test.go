package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DetectColorCodeStrIrregularErr(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		wantError bool
	}{
		{
			name:      "should return nil when string is empty",
			str:       "",
			wantError: false,
		},
		{
			name:      "should return nil when single valid 6-digit hex color is passed",
			str:       "#FF0000",
			wantError: false,
		},
		{
			name:      "should return nil when valid hex color without hash is passed",
			str:       "00FF00",
			wantError: false,
		},
		{
			name:      "should return nil when multiple valid colors joined by AndOperator",
			str:       "#FF0000 & #00FF00",
			wantError: false,
		},
		{
			name:      "should return error when hex color length is invalid",
			str:       "#FF00",
			wantError: true,
		},
		{
			name:      "should return error when hex color contains non-hex characters",
			str:       "#GG0000",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DetectColorCodeStrIrregularErr(tt.str)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
