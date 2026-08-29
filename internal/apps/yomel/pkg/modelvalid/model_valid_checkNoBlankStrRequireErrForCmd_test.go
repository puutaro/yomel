package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_checkNoBlankStrRequireErrForCmd(t *testing.T) {
	tests := []struct {
		name      string
		input     model.StageModel
		wantError string
	}{
		{
			name: "should return nil when command is not blank",
			input: model.StageModel{
				No:  1,
				Cmd: "echo",
			},
			wantError: "",
		},
		{
			name: "should return error when command is blank",
			input: model.StageModel{
				No:  1,
				Cmd: "",
			},
			wantError: "'-c' no blank str is required\nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkNoBlankStrRequireErrForCmd(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
