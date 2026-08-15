package domain

import (
	"reflect"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

func Test_makeOptListWithComment(t *testing.T) {
	tests := []struct {
		name       string
		optPs      []model.OptParam
		opPrefix   string
		isTerminal bool
		want       []opArgType
	}{
		{
			name: "Param is nil with comment",
			optPs: []model.OptParam{
				{
					Index:   1,
					OptStr:  "--opt1",
					Comment: "OptComment",
					Param:   model.ParamType{Str: nil},
				},
			},
			opPrefix:   "",
			isTerminal: false,
			want: []opArgType{
				{
					Index: 1,
					Str:   "--opt1" + "SIED_COmMetN_BAlNk" + "`# opt comment`",
				},
			},
		},
		{
			name: "DoubleQuote with value comment",
			optPs: []model.OptParam{
				{
					Index:   2,
					OptStr:  "--opt2",
					Comment: "OpCom",
					Param: model.ParamType{
						Str:       testutil.Ptr("val2"),
						QuoteType: argtables.DoubleQuote,
						Comment:   "ValCom",
					},
				},
			},
			opPrefix:   "",
			isTerminal: false,
			want: []opArgType{
				{
					Index: 2,
					Str:   "`# op com` \\\n --opt2 \"val2\"" + "SIED_COmMetN_BAlNk" + "`# val com`",
				},
			},
		},
		{
			name: "SingleQuote without comment",
			optPs: []model.OptParam{
				{
					Index:   3,
					OptStr:  "--opt3",
					Comment: "",
					Param: model.ParamType{
						Str:       testutil.Ptr("val3"),
						QuoteType: argtables.SingleQuote,
						Comment:   "",
					},
				},
			},
			opPrefix:   "",
			isTerminal: false,
			want: []opArgType{
				{
					Index: 3,
					Str:   "--opt3 'val3'",
				},
			},
		},
		{
			name: "NoQuote with option comment only",
			optPs: []model.OptParam{
				{
					Index:   4,
					OptStr:  "--opt4",
					Comment: "OpCom",
					Param: model.ParamType{
						Str:       testutil.Ptr("val4"),
						QuoteType: argtables.NoQuote,
						Comment:   "",
					},
				},
			},
			opPrefix:   "",
			isTerminal: false,
			want: []opArgType{
				{
					Index: 4,
					Str:   "`# op com` \\\n --opt4 val4",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeOptListWithComment(tt.optPs, tt.opPrefix, tt.isTerminal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeOptListWithComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
