package argtabledtosvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/stretchr/testify/assert"
)

func Test_checkQuoteOptionIrregularPositionErr(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtabledtos.ArgTableDto
		wantError string
	}{
		{
			name: "should return nil when quote option is immediately after arg or value",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, IsArg: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
			},
			wantError: "",
		},
		{
			name: "should return error when quote option is in irregular position",
			input: []argtabledtos.ArgTableDto{
				{StageNo: 1, IsStage: true},
				{StageNo: 1, IsCmd: true},
				{StageNo: 1, IsOpt: true},
				{StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
			},
			wantError: "'--s/--single' and '--n/--no-quote' must be immediately after '--arg' and '--val'\nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkQuoteOptionIrregularPositionErr(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
