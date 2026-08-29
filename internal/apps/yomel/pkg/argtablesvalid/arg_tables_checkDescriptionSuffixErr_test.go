package argtablesvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

// Test_checkDescriptionSuffixErr verifies that description suffixes correctly validate PascalCase rules and irregular repetitions.
func Test_checkDescriptionSuffixErr(t *testing.T) {
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
			wantError: "Description suffix of '-o' and '--o' and '-a' and '-v' must be alphanumeric pascalCase in \nstageNo: 1",
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
			wantError: "Description suffix of '-o' and '--o' and '-a' and '-v' must be alphanumeric pascalCase in \nstageNo: 2",
		},
		{
			name: "should return error when comment suffix consists of repeated single characters",
			input: []argtables.ArgTable{
				{
					StageNo: 1,
					IsStage: true,
				},
				{
					StageNo: 1,
					IsOpt:   true,
					Comment: "aaa",
				},
			},
			wantError: "Description suffix of '-o' and '--o' and '-a' and '-v' must be alphanumeric pascalCase in \nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkDescriptionSuffixErr(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
