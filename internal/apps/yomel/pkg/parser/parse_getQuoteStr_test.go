package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_getQuoteStr(t *testing.T) {
	tests := []struct {
		name     string
		input    []args.ArgTable
		curIndex int
		want     paramType
		wantIdx  int
	}{
		{
			name: "should return double quoted string and next index when quote type signal is DoubleQuote",
			input: []args.ArgTable{
				{}, // index 0
				{}, // index 1 (curIndex)
				{QuoteTypeSignal: args.DoubleQuote, Str: testutil.Ptr("double-val")}, // index 2
			},
			curIndex: 1,
			want: paramType{
				str:       testutil.Ptr("double-val"),
				quoteType: args.DoubleQuote, // note: getQuoteStr implementation doesn't explicitly set quoteType for DoubleQuote, it stays zero-value or whatever is in struct, let's verify exact fields
			},
			wantIdx: 2,
		},
		{
			name: "should return single quoted string, quote type, and updated index when quote type signal is SingleQuote",
			input: []args.ArgTable{
				{},                                  // index 0
				{},                                  // index 1 (curIndex)
				{QuoteTypeSignal: args.SingleQuote}, // index 2 (afterFirstIndex)
				{Str: testutil.Ptr("single-val"), IsValue: true}, // index 3 (afterNextIndex)
			},
			curIndex: 1,
			want: paramType{
				str:       testutil.Ptr("single-val"),
				quoteType: args.SingleQuote,
			},
			wantIdx: 3,
		},
		{
			name: "should return no-quote string, quote type, and updated index when quote type signal is NoQuote",
			input: []args.ArgTable{
				{},                              // index 0
				{},                              // index 1 (curIndex)
				{QuoteTypeSignal: args.NoQuote}, // index 2 (afterFirstIndex)
				{Str: testutil.Ptr("no-quote-val"), IsValue: true}, // index 3 (afterNextIndex)
			},
			curIndex: 1,
			want: paramType{
				str:       testutil.Ptr("no-quote-val"),
				quoteType: args.NoQuote,
			},
			wantIdx: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParam, gotIdx := getQuoteStr(tt.input, tt.curIndex)
			assert.Equal(t, tt.want.str, gotParam.str)
			assert.Equal(t, tt.want.quoteType, gotParam.quoteType)
			assert.Equal(t, tt.wantIdx, gotIdx)
		})
	}
}
