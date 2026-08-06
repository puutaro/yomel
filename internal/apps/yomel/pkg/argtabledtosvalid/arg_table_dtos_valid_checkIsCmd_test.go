package argtabledtosvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/stretchr/testify/assert"
)

func Test_checkIsCmd(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtabledtos.ArgTableDto
		wantError string
	}{
		{
			name: "should return nil when cmd is specified in each stage",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, Str: testutilPtr("echo")},
			},
			wantError: "",
		},
		{
			name: "should return error when stage is defined but cmd is missing",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
			},
			wantError: "'-cmd' not found\nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkIsCmd(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
