package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_Convert verifies that model.ControlModel and stage models are correctly converted into domain.Yomel structures.
func Test_Convert(t *testing.T) {
	tests := []struct {
		name       string
		ctrl       model.ControlModel
		stageMods  []model.StageModel
		wantResult Yomel
	}{
		{
			name: "should convert control and single stage model with arguments and options correctly",
			ctrl: model.ControlModel{
				IsLog:        testutil.Ptr(true),
				LogFilter:    "global-log-filter",
				ErrLogFilter: "global-err-filter",
			},
			stageMods: []model.StageModel{
				{
					No:   1,
					Desc: "stage1",
					Cmd:  "echo",
					CmdOps: []model.OptParam{
						{
							Index:  2,
							OptStr: "n",
							Param:  model.ParamType{},
						},
					},
					CmdArgs: []model.ArgParam{
						{
							Index: 4,
							Param: model.ParamType{
								Str:       testutil.Ptr("hello yomel"),
								QuoteType: argtabledtos.SingleQuote,
							},
						},
					},
					IsLog:        testutil.Ptr(false),
					LogFilter:    "stage-log-filter",
					ErrLogFilter: "stage-err-filter",
				},
			},
			wantResult: Yomel{
				Ctrl: Control{
					IsLog:        testutil.Ptr(true),
					LogFilter:    "global-log-filter",
					ErrLogFilter: "global-err-filter",
				},
				Stages: []Stage{
					{
						No:           1,
						Desc:         "stage1",
						Cmd:          "echo",
						CmdOpArgs:    []string{"-n", "'hello yomel'"},
						IsLog:        testutil.Ptr(false),
						LogFilter:    "stage-log-filter",
						ErrLogFilter: "stage-err-filter",
					},
				},
			},
		},
		{
			name: "should handle comprehensive command, service, and action parameters with various quotes",
			ctrl: model.ControlModel{
				IsLog: testutil.Ptr(false),
			},
			stageMods: []model.StageModel{
				{
					No:   1,
					Desc: "comprehensive-stage",
					Cmd:  "aws",
					CmdLops: []model.OptParam{
						{
							Index:  2,
							OptStr: "profile",
							Param:  model.ParamType{},
						},
					},
					Svc: testutil.Ptr("s3"),
					SvcOps: []model.OptParam{
						{
							Index:  5,
							OptStr: "r",
							Param:  model.ParamType{},
						},
					},
					Act: testutil.Ptr("cp"),
					ActArgs: []model.ArgParam{
						{
							Index: 8,
							Param: model.ParamType{
								Str:       testutil.Ptr("s3://my-bucket/path"),
								QuoteType: argtabledtos.NoQuote,
							},
						},
					},
				},
			},
			wantResult: Yomel{
				Ctrl: Control{
					IsLog: testutil.Ptr(false),
				},
				Stages: []Stage{
					{
						No:        1,
						Desc:      "comprehensive-stage",
						Cmd:       "aws",
						CmdOpArgs: []string{"--profile"},
						Svc:       "s3",
						SvcOpArgs: []string{"-r"},
						Act:       "cp",
						ActOpArgs: []string{"s3://my-bucket/path"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Convert(tt.ctrl, tt.stageMods)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}
