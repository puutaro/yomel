package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseOptions(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{Str: &s} }
	tCmd := func() argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{IsCmd: true} }
	tSvc := func() argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{IsSvc: true} }
	// tAct := func() args.ArgTable { return args.ArgTable{IsAct: true} }
	tOpt := func() argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{IsOpt: true} }
	tLopt := func() argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{IsLopt: true} }
	tVal := func() argtabledtos.ArgTableDto { return argtabledtos.ArgTableDto{IsValue: true} }
	// tArg := func() args.ArgTable { return args.ArgTable{IsArg: true} }

	tests := []struct {
		name            string
		nextStartIndex  int
		input           []argtabledtos.ArgTableDto
		isTargetMainArg func(argtabledtos.ArgTableDto) bool
		isNextMainArg   func(argtabledtos.ArgTableDto) bool
		isTargetOpt     func(argtabledtos.ArgTableDto) bool
		want            []OptParam
	}{
		{
			name:           "should parse short options with values correctly",
			nextStartIndex: 0,
			input: []argtabledtos.ArgTableDto{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: argtables.SingleQuote}, tStr("file.txt"),
			},
			isTargetMainArg: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			isNextMainArg:   func(a argtabledtos.ArgTableDto) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtabledtos.ArgTableDto) bool { return a.IsOpt },
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
			input: []argtabledtos.ArgTableDto{
				tCmd(),
				tLopt(), tStr("verbose"),
			},
			isTargetMainArg: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			isNextMainArg:   func(a argtabledtos.ArgTableDto) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtabledtos.ArgTableDto) bool { return a.IsLopt },
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
			input: []argtabledtos.ArgTableDto{
				tCmd(),
				tOpt(), tStr("f"),
				tVal(), {QuoteTypeSignal: argtables.NoQuote}, tStr("file1"),
				tSvc(), // Boundary to next main arg
				tOpt(), tStr("ignored"),
			},
			isTargetMainArg: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			isNextMainArg:   func(a argtabledtos.ArgTableDto) bool { return a.IsSvc || a.IsAct },
			isTargetOpt:     func(a argtabledtos.ArgTableDto) bool { return a.IsOpt },
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
			input: []argtabledtos.ArgTableDto{
				tSvc(),
				tOpt(), tStr("v"),
			},
			isTargetMainArg: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			isNextMainArg:   func(a argtabledtos.ArgTableDto) bool { return a.IsAct },
			isTargetOpt:     func(a argtabledtos.ArgTableDto) bool { return a.IsOpt },
			want:            nil,
		},
		{
			name:           "should skip elements before nextStartIndex",
			nextStartIndex: 4,
			input: []argtabledtos.ArgTableDto{
				tCmd(), tOpt(), tStr("a"), tVal(), tStr("val-skip"),
				tCmd(), tOpt(), tStr("b"), tVal(), {QuoteTypeSignal: argtables.NoQuote}, tStr("val-target"),
			},
			isTargetMainArg: func(a argtabledtos.ArgTableDto) bool { return a.IsCmd },
			isNextMainArg:   func(a argtabledtos.ArgTableDto) bool { return a.IsAct },
			isTargetOpt:     func(a argtabledtos.ArgTableDto) bool { return a.IsOpt },
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
