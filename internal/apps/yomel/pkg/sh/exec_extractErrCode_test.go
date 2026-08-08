package sh

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractErrCode(t *testing.T) {

	tests := []struct {
		name     string
		err      error
		wantCode int
	}{
		{
			name:     "should return ExitErrGeneral when error is a standard generic error",
			err:      errors.New("generic error"),
			wantCode: ExitErrGeneral,
		},
		{
			name: "should extract correct exit code when error is an exec.ExitError",
			err: func() error {
				cmd := exec.Command("sh", "-c", "exit 123")
				err := cmd.Run()
				return err
			}(),
			wantCode: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode := extractErrCode(tt.err)
			assert.Equal(t, tt.wantCode, gotCode)
		})
	}
}
