package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_makeOptList(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			optPs    []optParam
			opPrefix string
		}
		want []opArgType
	}{
		{
			name: "normal single-quoted option",
			input: struct {
				optPs    []optParam
				opPrefix string
			}{
				optPs: []optParam{
					{
						index:  5,
						optStr: "e",
						param: paramType{
							str:       testutil.Ptr("s/aa/bb/g"),
							quoteType: args.SingleQuote,
						},
					},
				},
				opPrefix: "--",
			},
			want: []opArgType{
				{
					index: 5,
					str:   `--e 's/aa/bb/g'`,
				},
			},
		},
		{
			name: "empty input slice should return nil or empty slice",
			input: struct {
				optPs    []optParam
				opPrefix string
			}{
				optPs:    []optParam{},
				opPrefix: "--",
			},
			want: nil,
		},
		{
			name: "option with nil string should produce prefix and option string only",
			input: struct {
				optPs    []optParam
				opPrefix string
			}{
				optPs: []optParam{
					{
						index:  2,
						optStr: "v",
						param: paramType{
							str:       nil,
							quoteType: args.NoQuote,
						},
					},
				},
				opPrefix: "-",
			},
			want: []opArgType{
				{
					index: 2,
					str:   "-v",
				},
			},
		},
		{
			name: "mix of double quote, single quote, no quote, and nil string options with short prefix",
			input: struct {
				optPs    []optParam
				opPrefix string
			}{
				optPs: []optParam{
					{
						index:  1,
						optStr: "f",
						param: paramType{
							str:       nil,
							quoteType: args.NoQuote,
						},
					},
					{
						index:  3,
						optStr: "n",
						param: paramType{
							str:       testutil.Ptr("100"),
							quoteType: args.NoQuote,
						},
					},
					{
						index:  7,
						optStr: "m",
						param: paramType{
							str:       testutil.Ptr("message text"),
							quoteType: args.SingleQuote,
						},
					},
					{
						index:  9,
						optStr: "c",
						param: paramType{
							str:       testutil.Ptr("config.json"),
							quoteType: args.DoubleQuote,
						},
					},
				},
				opPrefix: "-",
			},
			want: []opArgType{
				{
					index: 1,
					str:   "-f",
				},
				{
					index: 3,
					str:   "-n 100",
				},
				{
					index: 7,
					str:   `-m 'message text'`,
				},
				{
					index: 9,
					str:   `-c "config.json"`,
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
