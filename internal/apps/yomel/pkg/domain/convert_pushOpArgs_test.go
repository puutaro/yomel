package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_pushOpArgs(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			ops  []model.OptParam
			lOps []model.OptParam
			args []model.ArgParam
		}
		want []string
	}{
		{
			name: "should correctly combine and sort short options, long options, and arguments by index",
			input: struct {
				ops  []model.OptParam
				lOps []model.OptParam
				args []model.ArgParam
			}{
				ops: []model.OptParam{
					{
						Index:  2,
						OptStr: "f",
						Param: model.ParamType{
							Str:       testutil.Ptr("file.txt"),
							QuoteType: argtabledtos.NoQuote,
						},
					},
				},
				lOps: []model.OptParam{
					{
						Index:  5,
						OptStr: "region",
						Param: model.ParamType{
							Str:       testutil.Ptr("us-east-1"),
							QuoteType: argtabledtos.SingleQuote,
						},
					},
				},
				args: []model.ArgParam{
					{
						Index: 1,
						Param: model.ParamType{
							Str:       testutil.Ptr("first-arg"),
							QuoteType: argtabledtos.NoQuote,
						},
					},
					{
						Index: 10,
						Param: model.ParamType{
							Str:       testutil.Ptr("last-arg"),
							QuoteType: argtabledtos.DoubleQuote,
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
				ops  []model.OptParam
				lOps []model.OptParam
				args []model.ArgParam
			}{
				ops:  []model.OptParam{},
				lOps: []model.OptParam{},
				args: []model.ArgParam{},
			},
			want: nil,
		},
		{
			name: "should handle options and arguments without values or with nil strings properly",
			input: struct {
				ops  []model.OptParam
				lOps []model.OptParam
				args []model.ArgParam
			}{
				ops: []model.OptParam{
					{
						Index:  3,
						OptStr: "v",
						Param: model.ParamType{
							Str:       nil,
							QuoteType: argtabledtos.NoQuote,
						},
					},
				},
				lOps: []model.OptParam{},
				args: []model.ArgParam{
					{
						Index: 1,
						Param: model.ParamType{
							Str:       nil,
							QuoteType: argtabledtos.NoQuote,
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
