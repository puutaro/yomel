package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtable"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseArg(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) argtable.ArgTable { return argtable.ArgTable{Str: &s} }
	tArg := func() argtable.ArgTable { return argtable.ArgTable{IsArg: true} }
	tAct := func() argtable.ArgTable { return argtable.ArgTable{IsAct: true} }
	tSvc := func() argtable.ArgTable { return argtable.ArgTable{IsSvc: true} }

	tests := []struct {
		name               string
		nextStartIndex     int
		input              []argtable.ArgTable
		isNextMainArg      func(t argtable.ArgTable) bool
		isTargetMainArg    func(t argtable.ArgTable) bool
		appendFn           func(ind int, p ParamType) // Added appendFn to table test items
		wantParam          []ParamType
		wantIndices        []int
		wantNextStartIndex int
	}{
		{
			name:           "should parse positional arguments correctly when target main arg matches",
			nextStartIndex: 0,
			input: []argtable.ArgTable{
				tAct(),
				tArg(),
				{QuoteTypeSignal: argtable.SingleQuote},
				tStr("arg1"),
				tArg(),
				{QuoteTypeSignal: argtable.NoQuote},
				tStr("arg2"),
			},
			isNextMainArg:   func(t argtable.ArgTable) bool { return false },
			isTargetMainArg: func(t argtable.ArgTable) bool { return t.IsAct },
			appendFn: func(ind int, p ParamType) {
				// Default append logic can be placed here or handled dynamically
			},
			wantParam: []ParamType{
				{Str: testutil.Ptr("arg1"), QuoteType: argtable.SingleQuote},
				{Str: testutil.Ptr("arg2"), QuoteType: argtable.NoQuote},
			},
			wantIndices:        []int{3, 6},
			wantNextStartIndex: 6,
		},
		{
			name:           "should stop parsing when next main arg is encountered",
			nextStartIndex: 0,
			input: []argtable.ArgTable{
				tAct(),
				tArg(), {QuoteTypeSignal: argtable.NoQuote}, tStr("arg1"),
				tSvc(), // Next main arg boundary
				tArg(), {QuoteTypeSignal: argtable.NoQuote}, tStr("arg-skipped"),
			},
			isNextMainArg:   func(t argtable.ArgTable) bool { return t.IsSvc },
			isTargetMainArg: func(t argtable.ArgTable) bool { return t.IsAct },
			appendFn: func(ind int, p ParamType) {
			},
			wantParam: []ParamType{
				{Str: testutil.Ptr("arg1"), QuoteType: argtable.NoQuote},
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
