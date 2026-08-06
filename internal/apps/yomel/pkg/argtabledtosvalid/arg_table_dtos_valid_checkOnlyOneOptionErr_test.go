package argtabledtosvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/stretchr/testify/assert"
)

func Test_checkOnlyOneOptionErr(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtabledtos.ArgTableDto
		wantError string
	}{
		{
			name: "should return nil when options appear only once",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, IsOpt: true},
			},
			wantError: "",
		},
		{
			name: "should return error when cmd is duplicated in a single stage",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, IsCmd: true},
			},
			wantError: "'-cmd' are only one in each stage field\nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkOnlyOneOptionErr(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
