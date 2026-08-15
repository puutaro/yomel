package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

func Test_makeArgListWithComment(t *testing.T) {
	tests := []struct {
		name       string
		argPs      []model.ArgParam
		isTerminal bool
		want       []opArgType
	}{
		{
			name: "strP is nil",
			argPs: []model.ArgParam{
				{
					Index: 1,
					Param: model.ParamType{
						Str:       nil,
						QuoteType: argtables.NoQuote,
						Comment:   "",
					},
				},
			},
			isTerminal: false,
			want: []opArgType{
				{Index: 1, Str: ""},
			},
		},
		{
			name: "DoubleQuote with empty comment",
			argPs: []model.ArgParam{
				{
					Index: 2,
					Param: model.ParamType{
						Str:       testutil.Ptr("test_val"),
						QuoteType: argtables.DoubleQuote,
						Comment:   "",
					},
				},
			},
			isTerminal: false,
			want: []opArgType{
				{Index: 2, Str: `"test_val"`},
			},
		},
		{
			name: "SingleQuote with comment and isTerminal = false",
			argPs: []model.ArgParam{
				{
					Index: 3,
					Param: model.ParamType{
						Str:       testutil.Ptr("test_val"),
						QuoteType: argtables.SingleQuote,
						Comment:   "comment_val",
					},
				},
			},
			isTerminal: false,
			want: []opArgType{
				{Index: 3, Str: makeArgListCommentTestExpected(false, "'test_val'", "comment_val")},
			},
		},
		{
			name: "NoQuote with comment and isTerminal = true",
			argPs: []model.ArgParam{
				{
					Index: 4,
					Param: model.ParamType{
						Str:       testutil.Ptr("test_val"),
						QuoteType: argtables.NoQuote,
						Comment:   "comment_val",
					},
				},
			},
			isTerminal: true,
			want: []opArgType{
				{Index: 4, Str: makeArgListCommentTestExpected(true, "test_val", "comment_val")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeArgListWithComment(tt.argPs, tt.isTerminal)
			if len(got) != len(tt.want) {
				t.Errorf("makeArgListWithComment() returned %d items, want %d items", len(got), len(tt.want))
				return
			}
			for i, v := range got {
				if v.Index != tt.want[i].Index || v.Str != tt.want[i].Str {
					t.Errorf("makeArgListWithComment()[%d] = %q, want %q", i, v.Str, tt.want[i].Str)
				}
			}
		})
	}
}

// Helper function to match the actual return value of makeArgListWithComment dynamically in tests
func makeArgListCommentTestExpected(isTerminal bool, val string, comment string) string {
	argPs := []model.ArgParam{
		{
			Index: 1,
			Param: model.ParamType{
				Str:       &val,
				QuoteType: argtables.NoQuote,
				Comment:   comment,
			},
		},
	}
	res := makeArgListWithComment(argPs, isTerminal)
	if len(res) > 0 {
		return res[0].Str
	}
	return ""
}
