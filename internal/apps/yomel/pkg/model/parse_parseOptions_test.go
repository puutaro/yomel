package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_parseOptions verifies that parseOptions correctly parses short/long options with various quote types, boundary conditions, and filters.
func Test_parseOptions(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) argtables.ArgTable { return argtables.ArgTable{Str: &s} }
	tCmd := func() argtables.ArgTable { return argtables.ArgTable{IsCmd: true} }
	tSvc := func() argtables.ArgTable { return argtables.ArgTable{IsSvc: true} }
	// tAct := func() argtables.ArgTable { return argtables.ArgTable{IsAct: true} }
	tOpt := func() argtables.ArgTable { return argtables.ArgTable{IsOpt: true} }
	tLopt := func() argtables.ArgTable { return argtables.ArgTable{IsLopt: true} }
	tVal := func() argtables.ArgTable { return argtables.ArgTable{IsValue: true} }

	tests := []struct {
		name            string
		nextStartIndex  int
		input           []argtables.ArgTable
		isTargetMainArg func(argtables.ArgTable) bool
		isNextMainArg   func(argtables.ArgTable) bool
		isTargetOpt     func(argtables.ArgTable) bool
		want            []OptParam
	}{
		{
			name:           "should parse short options with values correctly",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: argtables.SingleQuote}, tStr("file.txt"),
			},
			isTargetMainArg: func(a argtables.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtables.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtables.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "f",
					Param: ParamType{
						Str:       testutil.Ptr("file.txt"),
						QuoteType: argtables.SingleQuote,
					},
				},
			},
		},
		{
			name:           "should parse long options without values correctly",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tCmd(),
				tLopt(), tStr("verbose"),
			},
			isTargetMainArg: func(a argtables.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtables.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtables.ArgTable) bool { return a.IsLopt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "verbose",
					Param:  ParamType{},
				},
			},
		},
		{
			name:           "should parse double quoted option value correctly",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tCmd(),
				tOpt(), tStr("m"),
				tVal(), tStr("double-val"), // DoubleQuote signal is implicit or handled by getQuoteStr
			},
			isTargetMainArg: func(a argtables.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtables.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtables.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "m",
					Param: ParamType{
						Str:       testutil.Ptr("double-val"),
						QuoteType: argtables.DoubleQuote,
					},
				},
			},
		},
		{
			name:           "should parse no-quote option value correctly",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tCmd(),
				tOpt(), tStr("c"),
				tVal(), {QuoteTypeSignal: argtables.NoQuote}, tStr("count"),
			},
			isTargetMainArg: func(a argtables.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtables.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtables.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "c",
					Param: ParamType{
						Str:       testutil.Ptr("count"),
						QuoteType: argtables.NoQuote,
					},
				},
			},
		},
		{
			name:           "should stop parsing when next main argument is encountered",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: argtables.NoQuote}, tStr("file1"),
				tSvc(), // Boundary to next main arg
				tOpt(), tStr("ignored"),
			},
			isTargetMainArg: func(a argtables.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtables.ArgTable) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtables.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  2,
					OptStr: "f",
					Param: ParamType{
						Str:       testutil.Ptr("file1"),
						QuoteType: argtables.NoQuote,
					},
				},
			},
		},
		{
			name:           "should return empty slice when target main argument is not found",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tSvc(),
				tOpt(), tStr("v"),
			},
			isTargetMainArg: func(a argtables.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtables.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a argtables.ArgTable) bool { return a.IsOpt },
			want:            nil,
		},
		{
			name:           "should skip elements before nextStartIndex",
			nextStartIndex: 4,
			input: []argtables.ArgTable{
				tCmd(), tOpt(), tStr("a"), tVal(), tStr("val-skip"),
				tCmd(), tOpt(), tStr("b"), tVal(), {QuoteTypeSignal: argtables.NoQuote}, tStr("val-target"),
			},
			isTargetMainArg: func(a argtables.ArgTable) bool { return a.IsCmd },
			isNextMainArg:   func(a argtables.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a argtables.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  7,
					OptStr: "b",
					Param: ParamType{
						Str:       testutil.Ptr("val-target"),
						QuoteType: argtables.NoQuote,
					},
				},
			},
		},
		{
			name:           "should parse service options when target main arg is svc",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tCmd(), tStr("aws"),
				tSvc(), tStr("s3"),
				tOpt(), tStr("r"),
			},
			isTargetMainArg: func(a argtables.ArgTable) bool { return a.IsSvc },
			isNextMainArg:   func(a argtables.ArgTable) bool { return a.IsAct },
			isTargetOpt:     func(a argtables.ArgTable) bool { return a.IsOpt },
			want: []OptParam{
				{
					Index:  5,
					OptStr: "r",
					Param:  ParamType{},
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
