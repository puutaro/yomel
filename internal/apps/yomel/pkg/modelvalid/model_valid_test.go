package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_ModelValidate(t *testing.T) {
	tests := []struct {
		name      string
		stModels  []model.StageModel
		wantError string
	}{
		{
			name: "should return nil when stage models are valid",
			stModels: []model.StageModel{
				{
					No:   1,
					Desc: "stage1",
					Cmd:  "echo",
				},
			},
			wantError: "",
		},
		{
			name: "should return error when stage command is empty",
			stModels: []model.StageModel{
				{
					No:  1,
					Cmd: "",
				},
			},
			wantError: "'//' description must be meaning sentence\nstageNo: 1\ndesc: ''",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ModelValidate(tt.stModels)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
