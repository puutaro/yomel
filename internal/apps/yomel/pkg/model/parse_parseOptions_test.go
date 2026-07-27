package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtable"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseOptions(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) argtable.ArgTable { return argtable.ArgTable{Str: &s} }
	tCmd := func() argtable.ArgTable { return argtable.ArgTable{IsCmd: true} }
	tSvc := func() argtable.ArgTable { return argtable.ArgTable{IsSvc: true} }
	// tAct := func() args.ArgTable { return args.ArgTable{IsAct: true} }
	tOpt := func() argtable.ArgTable { return argtable.ArgTable{IsOpt: true} }
	tLopt := func() argtable.ArgTable { return argtable.ArgTable{IsLopt: true} }
	tVal := func() argtable.ArgTable { return argtable.ArgTable{IsValue: true} }
	// tArg := func() args.ArgTable { return args.ArgTable{IsArg: true} }

	tests := []struct {
		name            string
		nextStartIndex  int
		input           []argtable.ArgTable
		isTargetMainArg func(argtable.ArgTable) bool
		isNextMainArg   func(argtable.ArgTable) bool
		isTargetOpt     func(argtable.ArgTable) bool
		want            []OptParam
	}{
		{
			name:           "should parse short options with values correctly",
			nextStartIndex: 0,
			input: []argtable.ArgTable{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: argtable.SingleQuote}, tStr("file.txt"),
			},
			isTargetMainArg: func(a argtable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtable.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtable.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "f",
					Param: ParamType{
						Str:       testutil.Ptr("file.txt"),
						QuoteType: argtable.SingleQuote,
					},
				},
			},
		},
		{
			name:           "should parse long options without values correctly",
			nextStartIndex: 0,
			input: []argtable.ArgTable{
				tCmd(),
				tLopt(), tStr("verbose"),
			},
			isTargetMainArg: func(a argtable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtable.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtable.ArgTable) bool { return a.IsLopt },
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
			input: []argtable.ArgTable{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: argtable.NoQuote}, tStr("file1"),
				tSvc(), // Boundary to next main arg
				tOpt(), tStr("ignored"),
			},
			isTargetMainArg: func(a argtable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtable.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtable.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "f",
					Param: ParamType{
						Str:       testutil.Ptr("file1"),
						QuoteType: argtable.NoQuote,
					},
				},
			},
		},
		{
			name:           "should return empty slice when target main argument is not found",
			nextStartIndex: 0,
			input: []argtable.ArgTable{
				tSvc(),
				tOpt(), tStr("v"),
			},
			isTargetMainArg: func(a argtable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtable.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a argtable.ArgTable) bool { return a.IsOpt },
			want:            nil,
		},
		{
			name:           "should skip elements before nextStartIndex",
			nextStartIndex: 4,
			input: []argtable.ArgTable{
				tCmd(), tOpt(), tStr("a"), tVal(), tStr("val-skip"),
				tCmd(), tOpt(), tStr("b"), tVal(), {QuoteTypeSignal: argtable.NoQuote}, tStr("val-target"),
			},
			isTargetMainArg: func(a argtable.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtable.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a argtable.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  7,
					OptStr: "b",
					Param: ParamType{
						Str:       testutil.Ptr("val-target"),
						QuoteType: argtable.NoQuote,
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
