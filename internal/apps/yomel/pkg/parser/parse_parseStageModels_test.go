package parser

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_parseStageModels(t *testing.T) {
	tests := []struct {
		name      string
		input     []args.ArgTable
		wantCtrl  Control
		wantStMod []stageModel
	}{
		{
			name: "should parse control flags and single stage model correctly",
			input: []args.ArgTable{
				{No: 1, StageNo: 0, IsLog: true},
				{No: 2, StageNo: 1, IsStage: true},
				{No: 3, StageNo: 1, Str: testutil.Ptr("stage1")},
				{No: 4, StageNo: 1, IsCmd: true},
				{No: 5, StageNo: 1, Str: testutil.Ptr("echo")},
				{No: 6, StageNo: 1, IsArg: true},
				{No: 7, StageNo: 1, QuoteTypeSignal: args.NoQuote},
				{No: 8, StageNo: 1, Str: testutil.Ptr("hello")},
			},
			wantCtrl: Control{
				IsLog:        true,
				LogFilter:    "",
				ErrLogFilter: "",
				IsVersion:    false,
				IsHelp:       false,
			},
			wantStMod: []stageModel{
				{
					no:   1,
					desc: "stage1",
					cmd:  "echo",
					cmdArgs: []argParam{
						{
							index: 6,
							param: paramType{
								str:       testutil.Ptr("hello"),
								quoteType: args.NoQuote,
							},
						},
					},
				},
			},
		},
		{
			name: "should parse multiple stages with options, services, actions, and filters",
			input: []args.ArgTable{
				{No: 1, StageNo: 0, IsLog: true},
				{No: 2, StageNo: 0, IsLogFilter: true},
				{No: 3, StageNo: 0, Str: testutil.Ptr("global-filter")},
				{No: 4, StageNo: 1, IsStage: true},
				{No: 5, StageNo: 1, Str: testutil.Ptr("fetch")},
				{No: 6, StageNo: 1, IsLogFilter: true},
				{No: 7, StageNo: 1, Str: testutil.Ptr("stage-filter")},
				{No: 8, StageNo: 1, IsCmd: true},
				{No: 9, StageNo: 1, Str: testutil.Ptr("curl")},
				{No: 10, StageNo: 1, IsOpt: true},
				{No: 11, StageNo: 1, Str: testutil.Ptr("s")},
				{No: 12, StageNo: 1, IsSvc: true},
				{No: 13, StageNo: 1, Str: testutil.Ptr("api")},
				{No: 14, StageNo: 1, IsAct: true},
				{No: 15, StageNo: 1, Str: testutil.Ptr("get")},
				{No: 16, StageNo: 2, IsStage: true},
				{No: 17, StageNo: 2, Str: testutil.Ptr("process")},
				{No: 18, StageNo: 2, IsCmd: true},
				{No: 19, StageNo: 2, Str: testutil.Ptr("cat")},
			},
			wantCtrl: Control{
				IsLog:        true,
				LogFilter:    "global-filter",
				ErrLogFilter: "",
				IsVersion:    false,
				IsHelp:       false,
			},
			wantStMod: []stageModel{
				{
					no:        1,
					desc:      "fetch",
					cmd:       "curl",
					cmdOps:    []optParam{{index: 7, optStr: "s", param: paramType{}}},
					svc:       "api",
					act:       "get",
					logFilter: "stage-filter",
				},
				{
					no:   2,
					desc: "process",
					cmd:  "cat",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtrl, gotStModels := parseStageModels(tt.input)
			assert.Equal(t, tt.wantCtrl, gotCtrl)
			assert.Equal(t, tt.wantStMod, gotStModels)
		})
	}
}
