package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseArg(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) args.ArgTable { return args.ArgTable{Str: &s} }
	tArg := func() args.ArgTable { return args.ArgTable{IsArg: true} }
	tAct := func() args.ArgTable { return args.ArgTable{IsAct: true} }
	tSvc := func() args.ArgTable { return args.ArgTable{IsSvc: true} }

	tests := []struct {
		name               string
		nextStartIndex     int
		input              []args.ArgTable
		isNextMainArg      func(t args.ArgTable) bool
		isTargetMainArg    func(t args.ArgTable) bool
		appendFn           func(ind int, p paramType) // Added appendFn to table test items
		wantParam          []paramType
		wantIndices        []int
		wantNextStartIndex int
	}{
		{
			name:           "should parse positional arguments correctly when target main arg matches",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tAct(),
				tArg(),
				{QuoteTypeSignal: args.SingleQuote},
				tStr("arg1"),
				tArg(),
				{QuoteTypeSignal: args.NoQuote},
				tStr("arg2"),
			},
			isNextMainArg:   func(t args.ArgTable) bool { return false },
			isTargetMainArg: func(t args.ArgTable) bool { return t.IsAct },
			appendFn: func(ind int, p paramType) {
				// Default append logic can be placed here or handled dynamically
			},
			wantParam: []paramType{
				{str: testutil.Ptr("arg1"), quoteType: args.SingleQuote},
				{str: testutil.Ptr("arg2"), quoteType: args.NoQuote},
			},
			wantIndices:        []int{3, 6},
			wantNextStartIndex: 6,
		},
		{
			name:           "should stop parsing when next main arg is encountered",
			nextStartIndex: 0,
			input: []args.ArgTable{
				tAct(),
				tArg(), {QuoteTypeSignal: args.NoQuote}, tStr("arg1"),
				tSvc(), // Next main arg boundary
				tArg(), {QuoteTypeSignal: args.NoQuote}, tStr("arg-skipped"),
			},
			isNextMainArg:   func(t args.ArgTable) bool { return t.IsSvc },
			isTargetMainArg: func(t args.ArgTable) bool { return t.IsAct },
			appendFn: func(ind int, p paramType) {
			},
			wantParam: []paramType{
				{str: testutil.Ptr("arg1"), quoteType: args.NoQuote},
			},
			wantIndices:        []int{3},
			wantNextStartIndex: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotParams []paramType
			var gotIndices []int

			// Use tt.appendFn if provided, otherwise fall back to local collectors
			appendFn := tt.appendFn
			if appendFn == nil {
				appendFn = func(ind int, p paramType) {
					gotIndices = append(gotIndices, ind)
					gotParams = append(gotParams, p)
				}
			} else {
				// Wrap to capture for assertions if table-defined appendFn is used
				origAppendFn := tt.appendFn
				appendFn = func(ind int, p paramType) {
					origAppendFn(ind, p)
					gotIndices = append(gotIndices, ind)
					gotParams = append(gotParams, p)
				}
			}

			nextIndex := parseArg(
				tt.nextStartIndex,
				tt.input,
				tt.isNextMainArg,
				tt.isTargetMainArg,
				appendFn,
			)

			assert.Equal(t, tt.wantNextStartIndex, nextIndex)
			assert.Equal(t, tt.wantParam, gotParams)
			assert.Equal(t, tt.wantIndices, gotIndices)
		})
	}
}
