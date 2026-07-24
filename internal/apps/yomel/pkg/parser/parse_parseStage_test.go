package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseStage(t *testing.T) {
	tests := []struct {
		name      string
		inputCtrl Control
		inputSts  []stageModel
		want      Yomel
	}{
		{
			name:      "should parse single stage model with command options and arguments correctly",
			inputCtrl: Control{IsLog: testutil.Ptr(true)},
			inputSts: []stageModel{
				{
					no:   1,
					desc: "echo-stage",
					cmd:  "echo",
					cmdArgs: []argParam{
						{
							index: 3,
							param: paramType{
								str:       testutil.Ptr("hello"),
								quoteType: args.NoQuote,
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
			inputCtrl: Control{LogFilter: "global-filter"},
			inputSts: []stageModel{
				{
					no:        1,
					desc:      "stage1",
					cmd:       "curl",
					cmdOps:    []optParam{{index: 2, optStr: "s", param: paramType{}}},
					svc:       "s3",
					act:       "cp",
					actLops:   []optParam{{index: 5, optStr: "region", param: paramType{str: testutil.Ptr("us-east-1"), quoteType: args.SingleQuote}}},
					logFilter: "custom-filter",
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
			inputCtrl: Control{IsLog: testutil.Ptr(true)},
			inputSts: []stageModel{
				{
					no:           1,
					desc:         "comprehensive-stage",
					cmd:          "aws",
					cmdLops:      []optParam{{index: 1, optStr: "profile", param: paramType{str: testutil.Ptr("prod"), quoteType: args.NoQuote}}},
					cmdArgs:      []argParam{{index: 2, param: paramType{str: testutil.Ptr("cmd-arg"), quoteType: args.DoubleQuote}}},
					svc:          "s3",
					svcOps:       []optParam{{index: 3, optStr: "r", param: paramType{}}},
					svcLops:      []optParam{{index: 4, optStr: "svc-lop", param: paramType{str: testutil.Ptr("val"), quoteType: args.SingleQuote}}},
					svcArgs:      []argParam{{index: 5, param: paramType{str: testutil.Ptr("svc-arg"), quoteType: args.NoQuote}}},
					act:          "cp",
					actOps:       []optParam{{index: 6, optStr: "v", param: paramType{}}},
					actLops:      []optParam{{index: 7, optStr: "recursive", param: paramType{}}},
					actArgs:      []argParam{{index: 8, param: paramType{str: testutil.Ptr("act-arg"), quoteType: args.NoQuote}}},
					isLog:        testutil.Ptr(false),
					logFilter:    "stage-log-filter",
					errLogFilter: "stage-err-filter",
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
			inputCtrl: Control{},
			inputSts:  []stageModel{},
			want: Yomel{
				Ctrl:   Control{},
				Stages: []Stage{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseStage(tt.inputCtrl, tt.inputSts)
			assert.Equal(t, tt.want, got)
		})
	}
}
