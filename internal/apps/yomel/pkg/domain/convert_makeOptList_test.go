package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_makeOptList(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			optPs    []model.OptParam
			opPrefix string
		}
		want []opArgType
	}{
		{
			name: "normal single-quoted option",
			input: struct {
				optPs    []model.OptParam
				opPrefix string
			}{
				optPs: []model.OptParam{
					{
						Index:  5,
						OptStr: "e",
						Param: model.ParamType{
							Str:       testutil.Ptr("s/aa/bb/g"),
							QuoteType: argtables.SingleQuote,
						},
					},
				},
				opPrefix: "--",
			},
			want: []opArgType{
				{
					Index: 5,
					Str:   `--e 's/aa/bb/g'`,
				},
			},
		},
		{
			name: "empty input slice should return nil or empty slice",
			input: struct {
				optPs    []model.OptParam
				opPrefix string
			}{
				optPs:    []model.OptParam{},
				opPrefix: "--",
			},
			want: nil,
		},
		{
			name: "option with nil string should produce prefix and option string only",
			input: struct {
				optPs    []model.OptParam
				opPrefix string
			}{
				optPs: []model.OptParam{
					{
						Index:  2,
						OptStr: "v",
						Param: model.ParamType{
							Str:       nil,
							QuoteType: argtables.NoQuote,
						},
					},
				},
				opPrefix: "-",
			},
			want: []opArgType{
				{
					Index: 2,
					Str:   "-v",
				},
			},
		},
		{
			name: "mix of double quote, single quote, no quote, and nil string options with short prefix",
			input: struct {
				optPs    []model.OptParam
				opPrefix string
			}{
				optPs: []model.OptParam{
					{
						Index:  1,
						OptStr: "f",
						Param: model.ParamType{
							Str:       nil,
							QuoteType: argtables.NoQuote,
						},
					},
					{
						Index:  3,
						OptStr: "n",
						Param: model.ParamType{
							Str:       testutil.Ptr("100"),
							QuoteType: argtables.NoQuote,
						},
					},
					{
						Index:  7,
						OptStr: "m",
						Param: model.ParamType{
							Str:       testutil.Ptr("message text"),
							QuoteType: argtables.SingleQuote,
						},
					},
					{
						Index:  9,
						OptStr: "c",
						Param: model.ParamType{
							Str:       testutil.Ptr("config.json"),
							QuoteType: argtables.DoubleQuote,
						},
					},
				},
				opPrefix: "-",
			},
			want: []opArgType{
				{
					Index: 1,
					Str:   "-f",
				},
				{
					Index: 3,
					Str:   "-n 100",
				},
				{
					Index: 7,
					Str:   `-m 'message text'`,
				},
				{
					Index: 9,
					Str:   `-c "config.json"`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := makeOptList(
				input.optPs,
				input.opPrefix,
			)
			assert.Equal(t, tt.want, got)
		})
	}
}
