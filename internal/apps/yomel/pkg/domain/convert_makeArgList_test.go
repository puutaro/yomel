package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtable"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_makeArgList(t *testing.T) {
	tests := []struct {
		name  string
		input []model.ArgParam
		want  []opArgType
	}{
		{
			name: "normal",
			input: []model.ArgParam{
				{
					Index: 10,
					Param: model.ParamType{
						Str:       testutil.Ptr("arg10"),
						QuoteType: argtable.DoubleQuote,
					},
				},
				{
					Index: 12,
					Param: model.ParamType{
						Str:       testutil.Ptr("arg12"),
						QuoteType: argtable.SingleQuote,
					},
				},
				{
					Index: 15,
					Param: model.ParamType{
						Str:       testutil.Ptr("arg15"),
						QuoteType: argtable.NoQuote,
					},
				},
			},
			want: []opArgType{
				{
					Index: 10,
					Str:   `"arg10"`,
				},
				{
					Index: 12,
					Str:   `'arg12'`,
				},
				{
					Index: 15,
					Str:   `arg15`,
				},
			},
		},
		{
			name:  "empty input slice should return empty opArgType slice",
			input: []model.ArgParam{},
			want:  nil,
		},
		{
			name: "argument with nil string should produce empty string",
			input: []model.ArgParam{
				{
					Index: 5,
					Param: model.ParamType{
						Str:       nil,
						QuoteType: argtable.NoQuote,
					},
				},
			},
			want: []opArgType{
				{
					Index: 5,
					Str:   "",
				},
			},
		},
		{
			name: "mix of various quote types and nil string",
			input: []model.ArgParam{
				{
					Index: 1,
					Param: model.ParamType{
						Str:       nil,
						QuoteType: argtable.DoubleQuote,
					},
				},
				{
					Index: 3,
					Param: model.ParamType{
						Str:       testutil.Ptr("raw-arg"),
						QuoteType: argtable.NoQuote,
					},
				},
				{
					Index: 7,
					Param: model.ParamType{
						Str:       testutil.Ptr("single-quoted-arg"),
						QuoteType: argtable.SingleQuote,
					},
				},
				{
					Index: 9,
					Param: model.ParamType{
						Str:       testutil.Ptr("double-quoted-arg"),
						QuoteType: argtable.DoubleQuote,
					},
				},
			},
			want: []opArgType{
				{
					Index: 1,
					Str:   "",
				},
				{
					Index: 3,
					Str:   "raw-arg",
				},
				{
					Index: 7,
					Str:   "'single-quoted-arg'",
				},
				{
					Index: 9,
					Str:   `"double-quoted-arg"`,
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
