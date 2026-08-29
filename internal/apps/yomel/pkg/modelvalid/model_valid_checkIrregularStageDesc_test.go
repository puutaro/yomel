package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_checkIrregularStageDesc(t *testing.T) {
	tests := []struct {
		name      string
		input     model.StageModel
		wantError string
	}{
		{
			name: "should return nil when stage description is valid",
			input: model.StageModel{
				No:   1,
				Desc: "stage1",
				Cmd:  "echo",
			},
			wantError: "",
		},
		{
			name: "should return error when stage description is empty",
			input: model.StageModel{
				No:   1,
				Desc: "",
				Cmd:  "echo",
			},
			wantError: "'//' description must be meaning sentence\nstageNo: 1\ndesc: ''",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkIrregularStageDesc(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
