package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_DetectColorStrIrregularErrForCtrl verifies that DetectColorStrIrregularErrForCtrl correctly checks irregular color strings.
func Test_DetectColorStrIrregularErrForCtrl(t *testing.T) {
	defaultErrMsg := "invalid color string for control"

	tests := []struct {
		name          string
		colorStr      string
		errMsg        string
		expectedIsErr bool
		expectedError string
	}{
		{
			name:          "should return no error when color string is valid hex",
			colorStr:      "#ff0000",
			errMsg:        defaultErrMsg,
			expectedIsErr: false,
		},
		{
			name:          "should return no error when color string is empty",
			colorStr:      "",
			errMsg:        defaultErrMsg,
			expectedIsErr: false,
		},
		{
			name:          "should return error when color string is irregular",
			colorStr:      "invalid_color",
			errMsg:        defaultErrMsg,
			expectedIsErr: true,
			expectedError: "invalid color string for control%!(EXTRA int=0, *errors.errorString=invalid hex color length or macro: invalid_color)",
		},
		{
			name:          "should return error with custom error message when color string is invalid",
			colorStr:      "badcolor",
			errMsg:        "custom invalid color error",
			expectedIsErr: true,
			expectedError: "custom invalid color error%!(EXTRA int=0, *errors.errorString=invalid hex color length or macro: badcolor)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DetectColorStrIrregularErrForCtrl(tt.colorStr, tt.errMsg)
			if tt.expectedIsErr {
				assert.Error(t, err)
				if err != nil {
					assert.Equal(t, tt.expectedError, err.Error())
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
