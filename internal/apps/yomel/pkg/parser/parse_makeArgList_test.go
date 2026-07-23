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
		{
			name:  "empty input slice should return empty opArgType slice",
			input: []argParam{},
			want:  nil,
		},
		{
			name: "argument with nil string should produce empty string",
			input: []argParam{
				{
					index: 5,
					param: paramType{
						str:       nil,
						quoteType: args.NoQuote,
					},
				},
			},
			want: []opArgType{
				{
					index: 5,
					str:   "",
				},
			},
		},
		{
			name: "mix of various quote types and nil string",
			input: []argParam{
				{
					index: 1,
					param: paramType{
						str:       nil,
						quoteType: args.DoubleQuote,
					},
				},
				{
					index: 3,
					param: paramType{
						str:       testutil.Ptr("raw-arg"),
						quoteType: args.NoQuote,
					},
				},
				{
					index: 7,
					param: paramType{
						str:       testutil.Ptr("single-quoted-arg"),
						quoteType: args.SingleQuote,
					},
				},
				{
					index: 9,
					param: paramType{
						str:       testutil.Ptr("double-quoted-arg"),
						quoteType: args.DoubleQuote,
					},
				},
			},
			want: []opArgType{
				{
					index: 1,
					str:   "",
				},
				{
					index: 3,
					str:   "raw-arg",
				},
				{
					index: 7,
					str:   "'single-quoted-arg'",
				},
				{
					index: 9,
					str:   `"double-quoted-arg"`,
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
