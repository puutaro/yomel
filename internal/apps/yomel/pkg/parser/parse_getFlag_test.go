package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/stretchr/testify/assert"
)

func Test_getFlag(t *testing.T) {
	tests := []struct {
		name           string                   // English test case description
		nextStartIndex int                      // Index to start searching from
		input          []args.ArgTable          // Slice of input arguments
		isCheckFn      func(args.ArgTable) bool // Condition check function
		defaultBool    bool                     // Default value if not found
		want           bool                     // Expected return value
	}{
		{
			name:           "should return true when target flag exists and default is false",
			nextStartIndex: 0,
			input: []args.ArgTable{
				{IsLog: true},
			},
			isCheckFn:   func(t args.ArgTable) bool { return t.IsLog },
			defaultBool: false,
			want:        true,
		},
		{
			name:           "should return false when target flag exists and default is true",
			nextStartIndex: 0,
			input: []args.ArgTable{
				{IsLog: true},
			},
			isCheckFn:   func(t args.ArgTable) bool { return t.IsLog },
			defaultBool: true,
			want:        false,
		},
		{
			name:           "should return false (default) when target flag does not exist and default is false",
			nextStartIndex: 0,
			input: []args.ArgTable{
				{IsLog: false},
			},
			isCheckFn:   func(t args.ArgTable) bool { return t.IsLog },
			defaultBool: false,
			want:        false,
		},
		{
			name:           "should return true (default) when target flag does not exist and default is true",
			nextStartIndex: 0,
			input: []args.ArgTable{
				{IsLog: false},
			},
			isCheckFn:   func(t args.ArgTable) bool { return t.IsLog },
			defaultBool: true,
			want:        true,
		},
		{
			name:           "should detect target flag when it exists at or after nextStartIndex",
			nextStartIndex: 1,
			input: []args.ArgTable{
				{IsLog: false},
				{IsLog: true},
			},
			isCheckFn:   func(t args.ArgTable) bool { return t.IsLog },
			defaultBool: false,
			want:        true,
		},
		{
			name:           "should ignore target flag and return default when it exists before nextStartIndex",
			nextStartIndex: 1,
			input: []args.ArgTable{
				{IsLog: true},
				{IsLog: false},
			},
			isCheckFn:   func(t args.ArgTable) bool { return t.IsLog },
			defaultBool: false,
			want:        false,
		},
		{
			name:           "should return default value when input slice is empty",
			nextStartIndex: 0,
			input:          []args.ArgTable{},
			isCheckFn:      func(t args.ArgTable) bool { return t.IsLog },
			defaultBool:    false,
			want:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFlag(
				tt.nextStartIndex,
				tt.input,
				tt.isCheckFn,
				tt.defaultBool,
			)
			// アサーション前に tt.want = got で上書きしていたバグを完全に修正
			assert.Equal(t, tt.want, got)
		})
	}
}
