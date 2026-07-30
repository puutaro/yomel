package argtablevalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

func Test_checkIsStage(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtables.ArgTable
		wantError string
	}{
		{
			name: "should return nil when stage is present",
			input: []argtables.ArgTable{
				{StageNo: 0, IsLog: true},
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, Str: testutilPtr("echo")},
			},
			wantError: "",
		},
		{
			name: "should return error when stage is missing",
			input: []argtables.ArgTable{
				{StageNo: 0, IsLog: true},
			},
			wantError: "'stage' not found",
		},
		{
			name:      "should return error when input slice is empty",
			input:     []argtables.ArgTable{},
			wantError: "'stage' not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkIsStage(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
