package argtabledtosvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/stretchr/testify/assert"
)

func Test_checkCmdSvcActOrderErr(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtabledtos.ArgTableDto
		wantError string
	}{
		{
			name: "should return nil when order is correct with cmd, svc, and act",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, IsSvc: true},
				{StageNo: 1, IsAct: true},
			},
			wantError: "",
		},
		{
			name: "should return nil when order is correct with only cmd and act",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, IsAct: true},
			},
			wantError: "",
		},
		{
			name: "should return error when act appears before cmd",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsAct: true},
				{StageNo: 1, IsCmd: true},
			},
			wantError: "'-cmd' and '-svc' and '-act' must be'-cmd' -> '-svc' -> '-act' order \nstageNo: 1",
		},
		{
			name: "should return error when svc appears before cmd",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsSvc: true},
				{StageNo: 1, IsCmd: true},
			},
			wantError: "'-cmd' and '-svc' and '-act' must be'-cmd' -> '-svc' -> '-act' order \nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkCmdSvcActOrderErr(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
