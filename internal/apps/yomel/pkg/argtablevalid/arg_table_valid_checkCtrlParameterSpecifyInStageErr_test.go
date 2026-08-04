package argtablevalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

func Test_checkCtrlParameterSpecifyInStageErr(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtables.ArgTable
		wantError string
	}{
		{
			name: "should return nil when control parameters are not specified in stage fields",
			input: []argtables.ArgTable{
				{StageNo: 0, IsVersion: true},
				{StageNo: 0, IsHelp: true},
				{StageNo: 0, IsDirect: true},
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, Str: testutilPtr("echo")},
			},
			wantError: "",
		},
		{
			name: "should return error when version is specified in stage field",
			input: []argtables.ArgTable{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsVersion: true},
			},
			wantError: "must be specified'--version' and '--help' and '--gen' and '--direct' in stage 0 field\nstageNo: 1",
		},
		{
			name: "should return error when help is specified in stage field",
			input: []argtables.ArgTable{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsHelp: true},
			},
			wantError: "must be specified'--version' and '--help' and '--gen' and '--direct' in stage 0 field\nstageNo: 1",
		},
		{
			name: "should return error when direct is specified in stage field",
			input: []argtables.ArgTable{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsDirect: true},
			},
			wantError: "must be specified'--version' and '--help' and '--gen' and '--direct' in stage 0 field\nstageNo: 1",
		},
		{
			name: "should return error when gen is specified in stage field",
			input: []argtables.ArgTable{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsGen: true},
			},
			wantError: "must be specified'--version' and '--help' and '--gen' and '--direct' in stage 0 field\nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkCtrlParameterSpecifyInStageErr(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
