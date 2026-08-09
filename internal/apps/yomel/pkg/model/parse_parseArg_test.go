package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_parseArg verifies that parseArg correctly parses positional arguments and respects main argument boundaries.
func Test_parseArg(t *testing.T) {
	// Tiny helpers to minimize structural boilerplate
	tStr := func(s string) argtables.ArgTable { return argtables.ArgTable{Str: &s} }
	tArg := func() argtables.ArgTable { return argtables.ArgTable{IsArg: true} }
	tAct := func() argtables.ArgTable { return argtables.ArgTable{IsAct: true} }
	tSvc := func() argtables.ArgTable { return argtables.ArgTable{IsSvc: true} }

	tests := []struct {
		name               string
		nextStartIndex     int
		input              []argtables.ArgTable
		isNextMainArg      func(t argtables.ArgTable) bool
		isTargetMainArg    func(t argtables.ArgTable) bool
		wantParam          []ParamType
		wantIndices        []int
		wantNextStartIndex int
	}{
		{
			name:           "should parse positional arguments correctly when target main arg matches",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tAct(),
				tArg(),
				{QuoteTypeSignal: argtables.SingleQuote},
				tStr("arg1"),
				tArg(),
				{QuoteTypeSignal: argtables.NoQuote},
				tStr("arg2"),
			},
			isNextMainArg:   func(t argtables.ArgTable) bool { return false },
			isTargetMainArg: func(t argtables.ArgTable) bool { return t.IsAct },
			wantParam: []ParamType{
				{Str: testutil.Ptr("arg1"), QuoteType: argtables.SingleQuote},
				{Str: testutil.Ptr("arg2"), QuoteType: argtables.NoQuote},
			},
			wantIndices:        []int{3, 6},
			wantNextStartIndex: 6,
		},
		{
			name:           "should stop parsing when next main arg is encountered",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tAct(),
				tArg(), {QuoteTypeSignal: argtables.NoQuote}, tStr("arg1"),
				tSvc(), // Next main arg boundary
				tArg(), {QuoteTypeSignal: argtables.NoQuote}, tStr("arg-skipped"),
			},
			isNextMainArg:   func(t argtables.ArgTable) bool { return t.IsSvc },
			isTargetMainArg: func(t argtables.ArgTable) bool { return t.IsAct },
			wantParam: []ParamType{
				{Str: testutil.Ptr("arg1"), QuoteType: argtables.NoQuote},
			},
			wantIndices:        []int{3},
			wantNextStartIndex: 3,
		},
		{
			name:           "should do nothing when target main arg does not match at start",
			nextStartIndex: 0,
			input: []argtables.ArgTable{
				tSvc(),
				tArg(), {QuoteTypeSignal: argtables.NoQuote}, tStr("arg1"),
			},
			isNextMainArg:      func(t argtables.ArgTable) bool { return t.IsSvc },
			isTargetMainArg:    func(t argtables.ArgTable) bool { return t.IsAct },
			wantParam:          nil,
			wantIndices:        nil,
			wantNextStartIndex: 0,
		},
		{
			name:           "should parse correctly starting from a middle index",
			nextStartIndex: 4,
			input: []argtables.ArgTable{
				/* 0 */ tSvc(),
				/* 1 */ tArg(), {QuoteTypeSignal: argtables.NoQuote}, tStr("dummy"),
				/* 4 */ tAct(),
				/* 5 */ tArg(),
				/* 6 */ {QuoteTypeSignal: argtables.NoQuote},
				/* 7 */ tStr("arg-target"),
			},
			isNextMainArg:   func(t argtables.ArgTable) bool { return t.IsSvc },
			isTargetMainArg: func(t argtables.ArgTable) bool { return t.IsAct },
			wantParam: []ParamType{
				{Str: testutil.Ptr("arg-target"), QuoteType: argtables.NoQuote},
			},
			wantIndices:        []int{7},
			wantNextStartIndex: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotParams []ParamType
			var gotIndices []int

			appendFn := func(ind int, p ParamType) {
				gotIndices = append(gotIndices, ind)
				gotParams = append(gotParams, p)
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
