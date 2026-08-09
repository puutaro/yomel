package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

// Test_checkOptStrBlankErrMsg verifies that checkOptStrBlankErrMsg correctly returns an error when any option or long option string is blank.
func Test_checkOptStrBlankErrMsg(t *testing.T) {
	tests := []struct {
		name      string
		stModels  []model.StageModel
		wantError string
	}{
		{
			name: "should return nil when all option strings are valid and not blank in Cmd",
			stModels: []model.StageModel{
				{
					No: 1,
					CmdOps: []model.OptParam{
						{OptStr: "f"},
					},
					CmdLops: []model.OptParam{
						{OptStr: "verbose"},
					},
				},
			},
			wantError: "",
		},
		{
			name: "should return nil when all option strings are valid and not blank in Svc",
			stModels: []model.StageModel{
				{
					No: 1,
					SvcOps: []model.OptParam{
						{OptStr: "f"},
					},
					SvcLops: []model.OptParam{
						{OptStr: "verbose"},
					},
				},
			},
			wantError: "",
		},
		{
			name: "should return nil when all option strings are valid and not blank in Act",
			stModels: []model.StageModel{
				{
					No: 1,
					ActOps: []model.OptParam{
						{OptStr: "f"},
					},
					ActLops: []model.OptParam{
						{OptStr: "verbose"},
					},
				},
			},
			wantError: "",
		},
		{
			name: "should return error when CmdOps has a blank option string",
			stModels: []model.StageModel{
				{
					No: 1,
					CmdOps: []model.OptParam{
						{OptStr: ""},
					},
				},
			},
			wantError: "'--opt' and '--lop' str is required\nstageNo: 1",
		},
		{
			name: "should return error when SvcLops has a blank option string in stage 2",
			stModels: []model.StageModel{
				{
					No: 1,
					CmdOps: []model.OptParam{
						{OptStr: "v"},
					},
				},
				{
					No: 2,
					SvcLops: []model.OptParam{
						{OptStr: ""},
					},
				},
			},
			wantError: "'--opt' and '--lop' str is required\nstageNo: 2",
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
