package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argTable"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseOptions(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) argTable.ArgTable { return argTable.ArgTable{Str: &s} }
	tCmd := func() argTable.ArgTable { return argTable.ArgTable{IsCmd: true} }
	tSvc := func() argTable.ArgTable { return argTable.ArgTable{IsSvc: true} }
	// tAct := func() args.ArgTable { return args.ArgTable{IsAct: true} }
	tOpt := func() argTable.ArgTable { return argTable.ArgTable{IsOpt: true} }
	tLopt := func() argTable.ArgTable { return argTable.ArgTable{IsLopt: true} }
	tVal := func() argTable.ArgTable { return argTable.ArgTable{IsValue: true} }
	// tArg := func() args.ArgTable { return args.ArgTable{IsArg: true} }

	tests := []struct {
		name            string
		nextStartIndex  int
		input           []argTable.ArgTable
		isTargetMainArg func(argTable.ArgTable) bool
		isNextMainArg   func(argTable.ArgTable) bool
		isTargetOpt     func(argTable.ArgTable) bool
		want            []OptParam
	}{
		{
			name:           "should parse short options with values correctly",
			nextStartIndex: 0,
			input: []argTable.ArgTable{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: argTable.SingleQuote}, tStr("file.txt"),
			},
			isTargetMainArg: func(a argTable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argTable.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argTable.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "f",
					Param: ParamType{
						Str:       testutil.Ptr("file.txt"),
						QuoteType: argTable.SingleQuote,
					},
				},
			},
		},
		{
			name:           "should parse long options without values correctly",
			nextStartIndex: 0,
			input: []argTable.ArgTable{
				tCmd(),
				tLopt(), tStr("verbose"),
			},
			isTargetMainArg: func(a argTable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argTable.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argTable.ArgTable) bool { return a.IsLopt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "verbose",
					Param:  ParamType{},
				},
			},
		},
		{
			name:           "should stop parsing when next main argument is encountered",
			nextStartIndex: 0,
			input: []argTable.ArgTable{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: argTable.NoQuote}, tStr("file1"),
				tSvc(), // Boundary to next main arg
				tOpt(), tStr("ignored"),
			},
			isTargetMainArg: func(a argTable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argTable.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argTable.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "f",
					Param: ParamType{
						Str:       testutil.Ptr("file1"),
						QuoteType: argTable.NoQuote,
					},
				},
			},
		},
		{
			name:           "should return empty slice when target main argument is not found",
			nextStartIndex: 0,
			input: []argTable.ArgTable{
				tSvc(),
				tOpt(), tStr("v"),
			},
			isTargetMainArg: func(a argTable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argTable.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a argTable.ArgTable) bool { return a.IsOpt },
			want:            nil,
		},
		{
			name:           "should skip elements before nextStartIndex",
			nextStartIndex: 4,
			input: []argTable.ArgTable{
				tCmd(), tOpt(), tStr("a"), tVal(), tStr("val-skip"),
				tCmd(), tOpt(), tStr("b"), tVal(), {QuoteTypeSignal: argTable.NoQuote}, tStr("val-target"),
			},
			isTargetMainArg: func(a argTable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argTable.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a argTable.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  7,
					OptStr: "b",
					Param: ParamType{
						Str:       testutil.Ptr("val-target"),
						QuoteType: argTable.NoQuote,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []OptParam
			parseOptions(
				tt.nextStartIndex,
				tt.input,
				tt.isTargetMainArg,
				tt.isNextMainArg,
				tt.isTargetOpt,
				func(p OptParam) {
					got = append(got, p)
				},
			)
			assert.Equal(t, tt.want, got)
		})
	}
}
