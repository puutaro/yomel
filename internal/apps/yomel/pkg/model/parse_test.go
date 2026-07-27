package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argTable"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_Parse verifies that control flags and various stage models are parsed correctly.
func Test_Parse(t *testing.T) {
	tests := []struct {
		name      string
		input     []argTable.ArgTable
		wantCtrl  ControlModel
		wantStMod []StageModel
	}{
		{
			name: "should parse control flags and single stage model correctly",
			input: []argTable.ArgTable{
				{No: 1, StageNo: 0, IsLog: true},
				{No: 2, StageNo: 1, IsStage: true},
				{No: 3, StageNo: 1, Str: testutil.Ptr("stage1")},
				{No: 4, StageNo: 1, IsCmd: true},
				{No: 5, StageNo: 1, Str: testutil.Ptr("echo")},
				{No: 6, StageNo: 1, IsArg: true},
				{No: 7, StageNo: 1, QuoteTypeSignal: argTable.NoQuote},
				{No: 8, StageNo: 1, Str: testutil.Ptr("hello")},
			},
			wantCtrl: ControlModel{
				IsLog:        testutil.Ptr(true),
				LogFilter:    "",
				ErrLogFilter: "",
				IsVersion:    false,
				IsHelp:       false,
			},
			wantStMod: []StageModel{
				{
					No:   1,
					Desc: "stage1",
					Cmd:  "echo",
					CmdArgs: []ArgParam{
						{
							Index: 6,
							Param: ParamType{
								Str:       testutil.Ptr("hello"),
								QuoteType: argTable.NoQuote,
							},
						},
					},
				},
			},
		},
		{
			name: "should parse multiple stages with options, services, actions, and filters",
			input: []argTable.ArgTable{
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
			wantCtrl: ControlModel{
				IsLog:        testutil.Ptr(true),
				LogFilter:    "global-filter",
				ErrLogFilter: "",
				IsVersion:    false,
				IsHelp:       false,
			},
			wantStMod: []StageModel{
				{
					No:        1,
					Desc:      "fetch",
					Cmd:       "curl",
					CmdOps:    []OptParam{{Index: 7, OptStr: "s", Param: ParamType{}}},
					Svc:       "api",
					Act:       "get",
					LogFilter: "stage-filter",
				},
				{
					No:   2,
					Desc: "process",
					Cmd:  "cat",
				},
			},
		},
		{
			name: "should handle control version, help, no-log, err-log-filter, and comprehensive stage parameters with service/action options and lopts",
			input: []argTable.ArgTable{
				{No: 1, StageNo: 0, IsVersion: true},
				{No: 2, StageNo: 0, IsHelp: true},
				{No: 3, StageNo: 0, IsNoLog: true},
				{No: 4, StageNo: 0, IsErrLogFilter: true},
				{No: 5, StageNo: 0, Str: testutil.Ptr("global-err-filter")},
				{No: 6, StageNo: 1, IsStage: true},
				{No: 7, StageNo: 1, Str: testutil.Ptr("stage-comprehensive")},
				{No: 8, StageNo: 1, IsNoLog: true},
				{No: 9, StageNo: 1, IsErrLogFilter: true},
				{No: 10, StageNo: 1, Str: testutil.Ptr("stage-err-filter")},
				{No: 11, StageNo: 1, IsCmd: true},
				{No: 12, StageNo: 1, Str: testutil.Ptr("aws")},
				{No: 13, StageNo: 1, IsLopt: true},
				{No: 14, StageNo: 1, Str: testutil.Ptr("profile")},
				{No: 15, StageNo: 1, IsSvc: true},
				{No: 16, StageNo: 1, Str: testutil.Ptr("s3")},
				{No: 17, StageNo: 1, IsOpt: true},
				{No: 18, StageNo: 1, Str: testutil.Ptr("r")},
				{No: 19, StageNo: 1, IsAct: true},
				{No: 20, StageNo: 1, Str: testutil.Ptr("cp")},
				{No: 21, StageNo: 1, IsLopt: true},
				{No: 22, StageNo: 1, Str: testutil.Ptr("recursive")},
				{No: 23, StageNo: 1, IsArg: true},
				{No: 24, StageNo: 1, QuoteTypeSignal: argTable.SingleQuote},
				{No: 25, StageNo: 1, Str: testutil.Ptr("arg-val")},
			},
			wantCtrl: ControlModel{
				IsVersion:    true,
				IsHelp:       true,
				IsLog:        testutil.Ptr(false),
				LogFilter:    "",
				ErrLogFilter: "global-err-filter",
			},
			wantStMod: []StageModel{
				{
					No:           1,
					Desc:         "stage-comprehensive",
					Cmd:          "aws",
					CmdLops:      []OptParam{{Index: 8, OptStr: "profile", Param: ParamType{}}},
					Svc:          "s3",
					SvcOps:       []OptParam{{Index: 12, OptStr: "r", Param: ParamType{}}},
					Act:          "cp",
					ActLops:      []OptParam{{Index: 16, OptStr: "recursive", Param: ParamType{}}},
					ActArgs:      []ArgParam{{Index: 19, Param: ParamType{Str: testutil.Ptr("arg-val"), QuoteType: argTable.SingleQuote}}},
					IsLog:        testutil.Ptr(false),
					ErrLogFilter: "stage-err-filter",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtrl, gotStModels := Parse(tt.input)
			assert.Equal(t, tt.wantCtrl, gotCtrl)
			assert.Equal(t, tt.wantStMod, gotStModels)
		})
	}
}
