package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_checkStageDescriptionDuplicate(t *testing.T) {
	tests := []struct {
		name      string
		input     []model.StageModel
		wantError string
	}{
		{
			name: "should return nil when stage descriptions are unique",
			input: []model.StageModel{
				{No: 1, Desc: "stage1"},
				{No: 2, Desc: "stage2"},
			},
			wantError: "",
		},
		{
			name: "should return error when stage descriptions are duplicate",
			input: []model.StageModel{
				{No: 1, Desc: "same-desc"},
				{No: 2, Desc: "same-desc"},
			},
			wantError: "'//' description must be unique across stages\nstageNo: 2\ndesc: 'same-desc'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkStageDescriptionDuplicate(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
