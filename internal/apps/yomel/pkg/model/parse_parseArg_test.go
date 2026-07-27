package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argTable"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseArg(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) argTable.ArgTable { return argTable.ArgTable{Str: &s} }
	tArg := func() argTable.ArgTable { return argTable.ArgTable{IsArg: true} }
	tAct := func() argTable.ArgTable { return argTable.ArgTable{IsAct: true} }
	tSvc := func() argTable.ArgTable { return argTable.ArgTable{IsSvc: true} }

	tests := []struct {
		name               string
		nextStartIndex     int
		input              []argTable.ArgTable
		isNextMainArg      func(t argTable.ArgTable) bool
		isTargetMainArg    func(t argTable.ArgTable) bool
		appendFn           func(ind int, p ParamType) // Added appendFn to table test items
		wantParam          []ParamType
		wantIndices        []int
		wantNextStartIndex int
	}{
		{
			name:           "should parse positional arguments correctly when target main arg matches",
			nextStartIndex: 0,
			input: []argTable.ArgTable{
				tAct(),
				tArg(),
				{QuoteTypeSignal: argTable.SingleQuote},
				tStr("arg1"),
				tArg(),
				{QuoteTypeSignal: argTable.NoQuote},
				tStr("arg2"),
			},
			isNextMainArg:   func(t argTable.ArgTable) bool { return false },
			isTargetMainArg: func(t argTable.ArgTable) bool { return t.IsAct },
			appendFn: func(ind int, p ParamType) {
				// Default append logic can be placed here or handled dynamically
			},
			wantParam: []ParamType{
				{Str: testutil.Ptr("arg1"), QuoteType: argTable.SingleQuote},
				{Str: testutil.Ptr("arg2"), QuoteType: argTable.NoQuote},
			},
			wantIndices:        []int{3, 6},
			wantNextStartIndex: 6,
		},
		{
			name:           "should stop parsing when next main arg is encountered",
			nextStartIndex: 0,
			input: []argTable.ArgTable{
				tAct(),
				tArg(), {QuoteTypeSignal: argTable.NoQuote}, tStr("arg1"),
				tSvc(), // Next main arg boundary
				tArg(), {QuoteTypeSignal: argTable.NoQuote}, tStr("arg-skipped"),
			},
			isNextMainArg:   func(t argTable.ArgTable) bool { return t.IsSvc },
			isTargetMainArg: func(t argTable.ArgTable) bool { return t.IsAct },
			appendFn: func(ind int, p ParamType) {
			},
			wantParam: []ParamType{
				{Str: testutil.Ptr("arg1"), QuoteType: argTable.NoQuote},
			},
			wantIndices:        []int{3},
			wantNextStartIndex: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotParams []ParamType
			var gotIndices []int

			// Use tt.appendFn if provided, otherwise fall back to local collectors
			appendFn := tt.appendFn
			if appendFn == nil {
				appendFn = func(ind int, p ParamType) {
					gotIndices = append(gotIndices, ind)
					gotParams = append(gotParams, p)
				}
			} else {
				// Wrap to capture for assertions if table-defined appendFn is used
				origAppendFn := tt.appendFn
				appendFn = func(ind int, p ParamType) {
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
