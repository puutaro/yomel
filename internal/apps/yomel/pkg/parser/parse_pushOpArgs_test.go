package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_pushOpArgs(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			ops  []optParam
			lOps []optParam
			args []argParam
		}
		want []string
	}{
		{
			name: "should correctly combine and sort short options, long options, and arguments by index",
			input: struct {
				ops  []optParam
				lOps []optParam
				args []argParam
			}{
				ops: []optParam{
					{
						index:  2,
						optStr: "f",
						param: paramType{
							str:       testutil.Ptr("file.txt"),
							quoteType: args.NoQuote,
						},
					},
				},
				lOps: []optParam{
					{
						index:  5,
						optStr: "region",
						param: paramType{
							str:       testutil.Ptr("us-east-1"),
							quoteType: args.SingleQuote,
						},
					},
				},
				args: []argParam{
					{
						index: 1,
						param: paramType{
							str:       testutil.Ptr("first-arg"),
							quoteType: args.NoQuote,
						},
					},
					{
						index: 10,
						param: paramType{
							str:       testutil.Ptr("last-arg"),
							quoteType: args.DoubleQuote,
						},
					},
				},
			},
			want: []string{
				"first-arg",
				"-f file.txt",
				"--region 'us-east-1'",
				`"last-arg"`,
			},
		},
		{
			name: "should return empty slice when all inputs are empty",
			input: struct {
				ops  []optParam
				lOps []optParam
				args []argParam
			}{
				ops:  []optParam{},
				lOps: []optParam{},
				args: []argParam{},
			},
			want: nil,
		},
		{
			name: "should handle options and arguments without values or with nil strings properly",
			input: struct {
				ops  []optParam
				lOps []optParam
				args []argParam
			}{
				ops: []optParam{
					{
						index:  3,
						optStr: "v",
						param: paramType{
							str:       nil,
							quoteType: args.NoQuote,
						},
					},
				},
				lOps: []optParam{},
				args: []argParam{
					{
						index: 1,
						param: paramType{
							str:       nil,
							quoteType: args.NoQuote,
						},
					},
				},
			},
			want: []string{
				"",
				"-v",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			pushOpArgs(
				tt.input.ops,
				tt.input.lOps,
				tt.input.args,
				func(opArgList []string) {
					got = opArgList
				},
			)
			assert.Equal(t, tt.want, got)
		})
	}
}
