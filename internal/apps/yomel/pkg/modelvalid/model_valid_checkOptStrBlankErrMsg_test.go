// Write direct above line for Comment on code
package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

// Test_checkOptStrBlankErrMsg verifies that checkOptStrBlankErrMsg correctly checks for blank option strings across all stages.
func Test_checkOptStrBlankErrMsg(t *testing.T) {
	tests := []struct {
		name      string
		stModels  []model.StageModel
		wantError string
	}{
		{
			name: "should return nil when all option strings are valid and not blank",
			stModels: []model.StageModel{
				{
					No:   1,
					Desc: "stage1",
					Cmd:  "aws",
					CmdOps: []model.OptParam{
						{OptStr: "profile"},
					},
				},
			},
			wantError: "",
		},
		{
			name: "should return error when command option string is blank",
			stModels: []model.StageModel{
				{
					No:   1,
					Desc: "stage1",
					Cmd:  "aws",
					CmdOps: []model.OptParam{
						{OptStr: ""},
					},
				},
			},
			wantError: "'--opt' and '--lop' str must not be blank\nstageNo: 1",
		},
		{
			name:      "should return nil when stages slice is empty",
			stModels:  []model.StageModel{},
			wantError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkOptStrBlankErrMsg(tt.stModels)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
