package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseOptions(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) args.ArgTable { return args.ArgTable{Str: &s} }
	tCmd := func() args.ArgTable { return args.ArgTable{IsCmd: true} }
	tSvc := func() args.ArgTable { return args.ArgTable{IsSvc: true} }
	// tAct := func() args.ArgTable { return args.ArgTable{IsAct: true} }
	tOpt := func() args.ArgTable { return args.ArgTable{IsOpt: true} }
	tLopt := func() args.ArgTable { return args.ArgTable{IsLopt: true} }
	tVal := func() args.ArgTable { return args.ArgTable{IsValue: true} }
	// tArg := func() args.ArgTable { return args.ArgTable{IsArg: true} }

	tests := []struct {
		name            string
		nextStartIndex  int
		input           []args.ArgTable
		isTargetMainArg func(args.ArgTable) bool
		isNextMainArg   func(args.ArgTable) bool
		isTargetOpt     func(args.ArgTable) bool
		want            []optParam
	}{
		{
			name:           "should parse short options with values correctly",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: args.SingleQuote}, tStr("file.txt"),
			},
			isTargetMainArg: func(a args.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a args.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a args.ArgTable) bool { return a.IsOpt },
			want: []optParam{
				{
					index:  2,
					optStr: "f",
					param: paramType{
						str:       testutil.Ptr("file.txt"),
						quoteType: args.SingleQuote,
					},
				},
			},
		},
		{
			name:           "should parse long options without values correctly",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tCmd(),
				tLopt(), tStr("verbose"),
			},
			isTargetMainArg: func(a args.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a args.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a args.ArgTable) bool { return a.IsLopt },
			want: []optParam{
				{
					index:  2,
					optStr: "verbose",
					param:  paramType{},
				},
			},
		},
		{
			name:           "should stop parsing when next main argument is encountered",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: args.NoQuote}, tStr("file1"),
				tSvc(), // Boundary to next main arg
				tOpt(), tStr("ignored"),
			},
			isTargetMainArg: func(a args.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a args.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a args.ArgTable) bool { return a.IsOpt },
			want: []optParam{
				{
					index:  2,
					optStr: "f",
					param: paramType{
						str:       testutil.Ptr("file1"),
						quoteType: args.NoQuote,
					},
				},
			},
		},
		{
			name:           "should return empty slice when target main argument is not found",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tSvc(),
				tOpt(), tStr("v"),
			},
			isTargetMainArg: func(a args.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a args.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a args.ArgTable) bool { return a.IsOpt },
			want:            nil,
		},
		{
			name:           "should skip elements before nextStartIndex",
			nextStartIndex: 4,
			input: []args.ArgTable{
				tCmd(), tOpt(), tStr("a"), tVal(), tStr("val-skip"),
				tCmd(), tOpt(), tStr("b"), tVal(), {QuoteTypeSignal: args.NoQuote}, tStr("val-target"),
			},
			isTargetMainArg: func(a args.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a args.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a args.ArgTable) bool { return a.IsOpt },
			want: []optParam{
				{
					index:  7,
					optStr: "b",
					param: paramType{
						str:       testutil.Ptr("val-target"),
						quoteType: args.NoQuote,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []optParam
			parseOptions(
				tt.nextStartIndex,
				tt.input,
				tt.isTargetMainArg,
				tt.isNextMainArg,
				tt.isTargetOpt,
				func(p optParam) {
					got = append(got, p)
				},
			)
			assert.Equal(t, tt.want, got)
		})
	}
}
