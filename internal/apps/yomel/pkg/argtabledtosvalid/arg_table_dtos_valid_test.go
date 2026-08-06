// Write direct above line for Comment on code
package argtabledtosvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_ArgTableDtosValid verifies that option suffixes follow the alphanumeric pascal case rule.
func Test_ArgTableDtosValid(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtabledtos.ArgTableDto
		wantError string
	}{
		{
			name: "should return nil when option suffixes are valid alphanumeric pascal case",
			input: []argtabledtos.ArgTableDto{
				{
					StageNo: 1,
					IsStage: true,
				},
				{
					StageNo: 1,
					Str:     testutil.Ptr("stage1"),
				},
				{
					StageNo: 1,
					OptStr:  testutil.Ptr("PascalCase"),
				},
				{
					StageNo: 1,
					LoptStr: testutil.Ptr("CamelCase"),
				},
				{
					StageNo: 1,
					ArgStr:  testutil.Ptr("ArgName"),
				},
				{
					StageNo:  1,
					ValueStr: testutil.Ptr("ValueName"),
				},
			},
			wantError: "",
		},
		{
			name: "should return nil when option suffixes are empty",
			input: []argtabledtos.ArgTableDto{
				{
					StageNo: 1,
					IsStage: true,
				},
				{
					StageNo: 1,
					Str:     testutil.Ptr("stage1"),
				},
				{
					StageNo: 1,
					OptStr:  nil,
				},
				{
					StageNo: 1,
					LoptStr: testutil.Ptr(""),
				},
			},
			wantError: "",
		},
		{
			name: "should return error when option suffix starts with lowercase letter",
			input: []argtabledtos.ArgTableDto{
				{
					StageNo: 1,
					IsStage: true,
				},
				{
					StageNo: 1,
					Str:     testutil.Ptr("stage1"),
				},
				{
					StageNo: 1,
					OptStr:  testutil.Ptr("lowerCase"),
				},
			},
			wantError: "Description suffix is must be alphanumeric pascalCase in '--opt' and '--lop' and '--arg' and '--val'\nstageNo: 1",
		},
		{
			name: "should return error when option suffix contains non-alphanumeric characters",
			input: []argtabledtos.ArgTableDto{
				{
					StageNo: 2,
					IsStage: true,
				},
				{
					StageNo: 2,
					Str:     testutil.Ptr("stage2"),
				},
				{
					StageNo: 2,
					LoptStr: testutil.Ptr("Invalid-Name"),
				},
			},
			wantError: "Description suffix is must be alphanumeric pascalCase in '--opt' and '--lop' and '--arg' and '--val'\nstageNo: 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ArgTableDtosValid(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
