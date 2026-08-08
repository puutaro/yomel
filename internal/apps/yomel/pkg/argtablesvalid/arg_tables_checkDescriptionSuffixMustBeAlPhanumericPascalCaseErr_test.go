package argtablesvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

func Test_checkDescriptionSuffixMustBeAlPhanumericPascalCaseErr(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtables.ArgTable
		wantError string
	}{
		{
			name: "should return nil when comment suffix is valid PascalCase",
			input: []argtables.ArgTable{
				{
					StageNo: 1,
					IsStage: true,
				},
				{
					StageNo: 1,
					IsOpt:   true,
					Comment: "ValidPascal123",
				},
			},
			wantError: "",
		},
		{
			name: "should return nil when comment suffix is empty",
			input: []argtables.ArgTable{
				{
					StageNo: 1,
					IsStage: true,
				},
				{
					StageNo: 1,
					IsOpt:   true,
					Comment: "",
				},
			},
			wantError: "",
		},
		{
			name: "should return error when comment suffix starts with lowercase letter",
			input: []argtables.ArgTable{
				{
					StageNo: 1,
					IsStage: true,
				},
				{
					StageNo: 1,
					IsLopt:  true,
					Comment: "invalidPascal",
				},
			},
			wantError: "Description suffix is must be alphanumeric pascalCase in '--opt' and '--lop' and '--arg' and '--val'\nstageNo: 1",
		},
		{
			name: "should return error when comment suffix contains non-alphanumeric characters",
			input: []argtables.ArgTable{
				{
					StageNo: 2,
					IsStage: true,
				},
				{
					StageNo: 2,
					IsArg:   true,
					Comment: "Invalid-Name",
				},
			},
			wantError: "Description suffix is must be alphanumeric pascalCase in '--opt' and '--lop' and '--arg' and '--val'\nstageNo: 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDescriptionSuffixMustBeAlPhanumericPascalCaseErr(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
