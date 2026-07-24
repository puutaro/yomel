package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_getFlagByPtr(t *testing.T) {
	tests := []struct {
		name           string
		nextStartIndex int
		input          []args.ArgTable
		isCheckFn      func(args.ArgTable) bool
		returnBool     bool
		want           *bool
	}{
		{
			name:           "should return pointer to true when target flag exists",
			nextStartIndex: 0,
			input: []args.ArgTable{
				{IsLog: true},
			},
			isCheckFn:  func(t args.ArgTable) bool { return t.IsLog },
			returnBool: true,
			want:       testutil.Ptr(true),
		},
		{
			name:           "should return pointer to false when target flag exists and returnBool is false",
			nextStartIndex: 0,
			input: []args.ArgTable{
				{IsNoLog: true},
			},
			isCheckFn:  func(t args.ArgTable) bool { return t.IsNoLog },
			returnBool: false,
			want:       testutil.Ptr(false),
		},
		{
			name:           "should return nil when target flag does not exist",
			nextStartIndex: 0,
			input: []args.ArgTable{
				{IsLog: false},
			},
			isCheckFn:  func(t args.ArgTable) bool { return t.IsLog },
			returnBool: true,
			want:       nil,
		},
		{
			name:           "should respect nextStartIndex and ignore flags before it",
			nextStartIndex: 1,
			input: []args.ArgTable{
				{IsLog: true},
				{IsLog: true},
			},
			isCheckFn:  func(t args.ArgTable) bool { return t.IsLog },
			returnBool: true,
			want:       testutil.Ptr(true),
		},
		{
			name:           "should return nil when flag exists before nextStartIndex",
			nextStartIndex: 1,
			input: []args.ArgTable{
				{IsLog: true},
				{IsLog: false},
			},
			isCheckFn:  func(t args.ArgTable) bool { return t.IsLog },
			returnBool: true,
			want:       nil,
		},
		{
			name:           "should return nil when input slice is empty",
			nextStartIndex: 0,
			input:          []args.ArgTable{},
			isCheckFn:      func(t args.ArgTable) bool { return t.IsLog },
			returnBool:     true,
			want:           nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFlagByPtr(
				tt.nextStartIndex,
				tt.input,
				tt.isCheckFn,
				tt.returnBool,
			)
			assert.Equal(t, tt.want, got)
		})
	}
}
