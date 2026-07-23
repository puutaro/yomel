package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_getOneStr(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) args.ArgTable { return args.ArgTable{Str: &s} }
	tCmd := func() args.ArgTable { return args.ArgTable{IsCmd: true} }
	tStage := func() args.ArgTable { return args.ArgTable{IsStage: true} }

	tests := []struct {
		name           string
		nextStartIndex int
		input          []args.ArgTable
		isCheckFn      func(args.ArgTable) bool
		want           *string
	}{
		{
			name:           "should return string pointer when target flag is matched and followed by a string",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tCmd(), tStr("aws"),
			},
			isCheckFn: func(a args.ArgTable) bool { return a.IsCmd },
			want:      testutil.Ptr("aws"),
		},
		{
			name:           "should skip elements before nextStartIndex",
			nextStartIndex: 2,
			input: []args.ArgTable{
				tStage(), tStr("skipped-stage"),
				tStage(), tStr("target-stage"),
			},
			isCheckFn: func(a args.ArgTable) bool { return a.IsStage },
			want:      testutil.Ptr("target-stage"),
		},
		{
			name:           "should return nil when matched flag is at the end of the slice",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tCmd(),
			},
			isCheckFn: func(a args.ArgTable) bool { return a.IsCmd },
			want:      nil,
		},
		{
			name:           "should return nil when the next element's Str field is nil",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tCmd(),
				tStage(), // Str field is nil
			},
			isCheckFn: func(a args.ArgTable) bool { return a.IsCmd },
			want:      nil,
		},
		{
			name:           "should return nil when target flag does not exist in the slice",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tCmd(), tStr("aws"),
			},
			isCheckFn: func(a args.ArgTable) bool { return a.IsStage },
			want:      nil,
		},
		{
			name:           "should return nil immediately when nextStartIndex exceeds input slice length",
			nextStartIndex: 3,
			input: []args.ArgTable{
				tCmd(), tStr("aws"),
			},
			isCheckFn: func(a args.ArgTable) bool { return a.IsCmd },
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getOneStr(tt.nextStartIndex, tt.input, tt.isCheckFn)
			assert.Equal(t, tt.want, got)
		})
	}
}
