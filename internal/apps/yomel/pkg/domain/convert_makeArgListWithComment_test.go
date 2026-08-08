// Write direct above line for Comment on code
package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_makeArgListWithComment(t *testing.T) {
	tests := []struct {
		name  string
		input []model.ArgParam
		want  []opArgType
	}{
		{
			name: "should generate argument list with comments and different quote types correctly",
			input: []model.ArgParam{
				{
					Index: 1,
					Param: model.ParamType{
						Str:       testutil.Ptr("arg1"),
						QuoteType: argtables.SingleQuote,
						Comment:   "ArgComment",
					},
				},
				{
					Index: 2,
					Param: model.ParamType{
						Str:       testutil.Ptr("arg2"),
						QuoteType: argtables.NoQuote,
						Comment:   "",
					},
				},
			},
			want: []opArgType{
				{
					Index: 1,
					Str:   `'arg1'` + SideCommentBlank + "\x1b[38;5;244m`# arg comment`\x1b[39m",
				},
				{
					Index: 2,
					Str:   `arg2`,
				},
			},
		},
		{
			name:  "should return nil when input slice is empty",
			input: []model.ArgParam{},
			want:  nil,
		},
		{
			name: "should handle nil string parameter correctly",
			input: []model.ArgParam{
				{
					Index: 3,
					Param: model.ParamType{
						Str:       nil,
						QuoteType: argtables.DoubleQuote,
					},
				},
			},
			want: []opArgType{
				{
					Index: 3,
					Str:   "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeArgListWithComment(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
