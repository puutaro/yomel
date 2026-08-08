package argtablesvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

func Test_ArgTableValidate(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtables.ArgTable
		wantError string
	}{
		{
			name: "should return nil when arg tables are valid",
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
				{StageNo: 0, IsCmd: true},
			},
			wantError: "'-cmd' must be specfied in stage field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ArgTableValidate(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}

func testutilPtr[T any](v T) *T {
	return &v
}
