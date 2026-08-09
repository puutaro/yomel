package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_getQuoteStr verifies that getQuoteStr correctly extracts parameter strings and quote types based on quote signals in the argument table.
func Test_getQuoteStr(t *testing.T) {
	tests := []struct {
		name     string
		input    []argtables.ArgTable
		curIndex int
		want     ParamType
		wantIdx  int
	}{
		{
			name: "should return double quoted string and next index when quote type signal is DoubleQuote",
			input: []argtables.ArgTable{
				{},                                // index 0
				{},                                // index 1 (curIndex)
				{Str: testutil.Ptr("double-val")}, // index 2
			},
			curIndex: 1,
			want: ParamType{
				Str:       testutil.Ptr("double-val"),
				QuoteType: argtables.DoubleQuote,
			},
			wantIdx: 2,
		},
		{
			name: "should return single quoted string, quote type, and updated index when quote type signal is SingleQuote",
			input: []argtables.ArgTable{
				{},                                       // index 0
				{},                                       // index 1 (curIndex)
				{QuoteTypeSignal: argtables.SingleQuote}, // index 2 (afterFirstIndex)
				{Str: testutil.Ptr("single-val"), IsValue: true}, // index 3 (afterNextIndex)
			},
			curIndex: 1,
			want: ParamType{
				Str:       testutil.Ptr("single-val"),
				QuoteType: argtables.SingleQuote,
			},
			wantIdx: 3,
		},
		{
			name: "should return no-quote string, quote type, and updated index when quote type signal is NoQuote",
			input: []argtables.ArgTable{
				{},                                   // index 0
				{},                                   // index 1 (curIndex)
				{QuoteTypeSignal: argtables.NoQuote}, // index 2 (afterFirstIndex)
				{Str: testutil.Ptr("no-quote-val"), IsValue: true}, // index 3 (afterNextIndex)
			},
			curIndex: 1,
			want: ParamType{
				Str:       testutil.Ptr("no-quote-val"),
				QuoteType: argtables.NoQuote,
			},
			wantIdx: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParam, gotIdx := getQuoteStr(tt.input, tt.curIndex)
			assert.Equal(t, tt.want.Str, gotParam.Str)
			assert.Equal(t, tt.want.QuoteType, gotParam.QuoteType)
			assert.Equal(t, tt.wantIdx, gotIdx)
		})
	}
}
