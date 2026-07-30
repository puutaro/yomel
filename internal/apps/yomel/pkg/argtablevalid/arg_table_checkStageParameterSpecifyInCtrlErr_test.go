package argtablevalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

func Test_checkStageParameterSpecifyInCtrlErr(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtables.ArgTable
		wantError string
	}{
		{
			name: "should return nil when stage parameters are not in control section",
			input: []argtables.ArgTable{
				{StageNo: 0, IsLog: true},
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, Str: testutilPtr("echo")},
			},
			wantError: "",
		},
		{
			name: "should return error when cmd is specified in control section",
			input: []argtables.ArgTable{
				{StageNo: 0, IsCmd: true},
			},
			wantError: "'-cmd' must be specfied in stage field",
		},
		{
			name: "should return error when svc is specified in control section",
			input: []argtables.ArgTable{
				{StageNo: 0, IsSvc: true},
			},
			wantError: "'-svc' must be specfied in stage field",
		},
		{
			name: "should return error when act is specified in control section",
			input: []argtables.ArgTable{
				{StageNo: 0, IsAct: true},
			},
			wantError: "'-act' must be specfied in stage field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkStageParameterSpecifyInCtrlErr(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
