package argtablevalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

func Test_checkUnkownOptionSpecifyedErrMsg(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtables.ArgTable
		wantError string
	}{
		{
			name: "should return nil when there are no unknown options",
			input: []argtables.ArgTable{
				{StageNo: 0, IsLog: true},
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, Str: testutilPtr("echo")},
			},
			wantError: "",
		},
		{
			name: "should return error when unknown option is specified in control section (stage 0)",
			input: []argtables.ArgTable{
				{StageNo: 0, UnkownOption: "--unknown-ctrl"},
			},
			wantError: "'--unknown-ctrl' is unknown option\nstageNo: 0",
		},
		{
			name: "should return error with correct stage number when unknown option is specified in a stage",
			input: []argtables.ArgTable{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, UnkownOption: "--unknown-stage"},
			},
			wantError: "'--unknown-stage' is unknown option\nstageNo: 1",
		},
		{
			name: "should return error with subsequent stage number when unknown option is specified in later stage",
			input: []argtables.ArgTable{
				{StageNo: 1, IsStage: true},
				{StageNo: 2, IsStage: true},
				{StageNo: 2, UnkownOption: "-u"},
			},
			wantError: "'-u' is unknown option\nstageNo: 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUnkownOptionSpecifyedErrMsg(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
