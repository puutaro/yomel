package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	tests := []struct {
		name      string
		inputCtrl model.ControlModel
		inputSts  []model.StageModel
		want      Yomel
	}{
		{
			name:      "should parse single stage model with command options and arguments correctly",
			inputCtrl: model.ControlModel{IsLog: testutil.Ptr(true)},
			inputSts: []model.StageModel{
				{
					No:   1,
					Desc: "echo-stage",
					Cmd:  "echo",
					CmdArgs: []model.ArgParam{
						{
							Index: 3,
							Param: model.ParamType{
								Str:       testutil.Ptr("hello"),
								QuoteType: args.NoQuote,
							},
						},
					},
				},
			},
			want: Yomel{
				Ctrl: Control{IsLog: testutil.Ptr(true)},
				Stages: []Stage{
					{
						No:        1,
						Desc:      "echo-stage",
						Cmd:       "echo",
						CmdOpArgs: []string{"hello"},
					},
				},
			},
		},
		{
			name:      "should parse multiple stages containing command, service, action options, and filters",
			inputCtrl: model.ControlModel{LogFilter: "global-filter"},
			inputSts: []model.StageModel{
				{
					No:        1,
					Desc:      "stage1",
					Cmd:       "curl",
					CmdOps:    []model.OptParam{{Index: 2, OptStr: "s", Param: model.ParamType{}}},
					Svc:       "s3",
					Act:       "cp",
					ActLops:   []model.OptParam{{Index: 5, OptStr: "region", Param: model.ParamType{Str: testutil.Ptr("us-east-1"), QuoteType: args.SingleQuote}}},
					LogFilter: "custom-filter",
				},
			},
			want: Yomel{
				Ctrl: Control{LogFilter: "global-filter"},
				Stages: []Stage{
					{
						No:           1,
						Desc:         "stage1",
						Cmd:          "curl",
						CmdOpArgs:    []string{"-s"},
						Svc:          "s3",
						Act:          "cp",
						ActOpArgs:    []string{"--region 'us-east-1'"},
						LogFilter:    "custom-filter",
						ErrLogFilter: "",
					},
				},
			},
		},
		{
			name:      "should parse comprehensive stage with all command, service, action options, lopts, args, and individual stage log settings",
			inputCtrl: model.ControlModel{IsLog: testutil.Ptr(true)},
			inputSts: []model.StageModel{
				{
					No:           1,
					Desc:         "comprehensive-stage",
					Cmd:          "aws",
					CmdLops:      []model.OptParam{{Index: 1, OptStr: "profile", Param: model.ParamType{Str: testutil.Ptr("prod"), QuoteType: args.NoQuote}}},
					CmdArgs:      []model.ArgParam{{Index: 2, Param: model.ParamType{Str: testutil.Ptr("cmd-arg"), QuoteType: args.DoubleQuote}}},
					Svc:          "s3",
					SvcOps:       []model.OptParam{{Index: 3, OptStr: "r", Param: model.ParamType{}}},
					SvcLops:      []model.OptParam{{Index: 4, OptStr: "svc-lop", Param: model.ParamType{Str: testutil.Ptr("val"), QuoteType: args.SingleQuote}}},
					SvcArgs:      []model.ArgParam{{Index: 5, Param: model.ParamType{Str: testutil.Ptr("svc-arg"), QuoteType: args.NoQuote}}},
					Act:          "cp",
					ActOps:       []model.OptParam{{Index: 6, OptStr: "v", Param: model.ParamType{}}},
					ActLops:      []model.OptParam{{Index: 7, OptStr: "recursive", Param: model.ParamType{}}},
					ActArgs:      []model.ArgParam{{Index: 8, Param: model.ParamType{Str: testutil.Ptr("act-arg"), QuoteType: args.NoQuote}}},
					IsLog:        testutil.Ptr(false),
					LogFilter:    "stage-log-filter",
					ErrLogFilter: "stage-err-filter",
				},
			},
			want: Yomel{
				Ctrl: Control{IsLog: testutil.Ptr(true)},
				Stages: []Stage{
					{
						No:           1,
						Desc:         "comprehensive-stage",
						Cmd:          "aws",
						CmdOpArgs:    []string{"--profile prod", `"cmd-arg"`},
						Svc:          "s3",
						SvcOpArgs:    []string{"-r", "--svc-lop 'val'", "svc-arg"},
						Act:          "cp",
						ActOpArgs:    []string{"-v", "--recursive", "act-arg"},
						IsLog:        testutil.Ptr(false),
						LogFilter:    "stage-log-filter",
						ErrLogFilter: "stage-err-filter",
					},
				},
			},
		},
		{
			name:      "should return empty stages when stage models slice is empty",
			inputCtrl: model.ControlModel{},
			inputSts:  []model.StageModel{},
			want: Yomel{
				Ctrl:   Control{},
				Stages: []Stage{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Convert(tt.inputCtrl, tt.inputSts)
			assert.Equal(t, tt.want, got)
		})
	}
}
