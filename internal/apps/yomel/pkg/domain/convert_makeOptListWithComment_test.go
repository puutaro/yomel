// Test_makeOptListWithComment verifies that makeOptListWithComment correctly generates option lists with comments and various quote types.
package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_makeOptListWithComment(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			optPs    []model.OptParam
			opPrefix string
		}
		want []opArgType
	}{
		{
			name: "should generate option list with comments and different quote types correctly",
			input: struct {
				optPs    []model.OptParam
				opPrefix string
			}{
				optPs: []model.OptParam{
					{
						Index:   1,
						OptStr:  "e",
						Comment: "OptComment",
						Param: model.ParamType{
							Str:       testutil.Ptr("val1"),
							QuoteType: argtables.SingleQuote,
							Comment:   "ValComment",
						},
					},
					{
						Index:   2,
						OptStr:  "v",
						Comment: "",
						Param: model.ParamType{
							Str:       nil,
							QuoteType: argtables.NoQuote,
						},
					},
				},
				opPrefix: "--",
			},
			want: []opArgType{
				{
					Index: 1,
					Str:   "\x1b[37m`#opt comment`\x1b[39m \\\n --e 'val1'" + SideCommentBlank + "\x1b[37m`#val comment`\x1b[39m",
				},
				{
					Index: 2,
					Str:   "--v" + SideCommentBlank,
				},
			},
		},
		{
			name: "should return nil when input slice is empty",
			input: struct {
				optPs    []model.OptParam
				opPrefix string
			}{
				optPs:    []model.OptParam{},
				opPrefix: "--",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeOptListWithComment(
				tt.input.optPs,
				tt.input.opPrefix,
			)
			assert.Equal(t, tt.want, got)
		})
	}
}
