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
			inputCtrl: Control{IsLog: true},
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
				Ctrl: Control{IsLog: true},
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
