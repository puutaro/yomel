package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_makeArgList(t *testing.T) {
	tests := []struct {
		name  string
		input []argParam
		want  []opArgType
	}{
		{
			name: "normal",
			input: []argParam{
				{
					index: 10,
					param: paramType{
						str:       testutil.Ptr("arg10"),
						quoteType: args.DoubleQuote,
					},
				},
				{
					index: 12,
					param: paramType{
						str:       testutil.Ptr("arg12"),
						quoteType: args.SingleQuote,
					},
				},
				{
					index: 15,
					param: paramType{
						str:       testutil.Ptr("arg15"),
						quoteType: args.NoQuote,
					},
				},
			},
			want: []opArgType{
				{
					index: 10,
					str:   `"arg10"`,
				},
				{
					index: 12,
					str:   `'arg12'`,
				},
				{
					index: 15,
					str:   `arg15`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeArgList(
				tt.input,
			)
			assert.Equal(t, tt.want, got)
		})
	}
}
