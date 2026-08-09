package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// Test_parseStageModels verifies that control flags and various stage models are parsed correctly.
func Test_parseStageModels(t *testing.T) {
	tests := []struct {
		name      string
		input     []argtables.ArgTable
		wantCtrl  ControlModel
		wantStMod []StageModel
	}{
		{
			name: "should parse control flags and single stage model correctly",
			input: []argtables.ArgTable{
				{No: 1, StageNo: 0, IsLog: true},
				{No: 2, StageNo: 1, IsStage: true},
				{No: 3, StageNo: 1, Str: testutil.Ptr("stage1")},
				{No: 4, StageNo: 1, IsCmd: true},
				{No: 5, StageNo: 1, Str: testutil.Ptr("echo")},
				{No: 6, StageNo: 1, IsArg: true},
				{No: 7, StageNo: 1, QuoteTypeSignal: argtables.NoQuote},
				{No: 8, StageNo: 1, Str: testutil.Ptr("hello")},
			},
			wantCtrl: ControlModel{
				IsLiveStdout: true,
				IsLiveStderr: true,
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
								QuoteType: argtables.NoQuote,
							},
						},
					},
				},
			},
		},
		{
			name: "should parse multiple stages with options, services, actions, and filters",
			input: []argtables.ArgTable{
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
				IsLiveStdout: true,
				IsLiveStderr: true,
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
					Svc:       testutil.Ptr("api"),
					Act:       testutil.Ptr("get"),
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
			input: []argtables.ArgTable{
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
				{No: 24, StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{No: 25, StageNo: 1, Str: testutil.Ptr("arg-val")},
			},
			wantCtrl: ControlModel{
				IsLiveStdout: true,
				IsLiveStderr: true,
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
					Svc:          testutil.Ptr("s3"),
					SvcOps:       []OptParam{{Index: 12, OptStr: "r", Param: ParamType{}}},
					Act:          testutil.Ptr("cp"),
					ActLops:      []OptParam{{Index: 16, OptStr: "recursive", Param: ParamType{}}},
					ActArgs:      []ArgParam{{Index: 19, Param: ParamType{Str: testutil.Ptr("arg-val"), QuoteType: argtables.SingleQuote}}},
					IsLog:        testutil.Ptr(false),
					ErrLogFilter: "stage-err-filter",
				},
			},
		},
		{
			name: "should handle service and action specific options, lopts, and arguments comprehensively",
			input: []argtables.ArgTable{
				{No: 1, StageNo: 1, IsStage: true},
				{No: 2, StageNo: 1, Str: testutil.Ptr("service-action-thorough")},
				{No: 3, StageNo: 1, IsCmd: true},
				{No: 4, StageNo: 1, Str: testutil.Ptr("kubectl")},
				{No: 5, StageNo: 1, IsSvc: true},
				{No: 6, StageNo: 1, Str: testutil.Ptr("pods")},
				{No: 7, StageNo: 1, IsSvc: false, IsOpt: true},
				{No: 8, StageNo: 1, Str: testutil.Ptr("n")},
				{No: 9, StageNo: 1, IsValue: true},
				{No: 10, StageNo: 1, QuoteTypeSignal: argtables.NoQuote},
				{No: 11, StageNo: 1, Str: testutil.Ptr("default")},
				{No: 12, StageNo: 1, IsAct: true},
				{No: 13, StageNo: 1, Str: testutil.Ptr("get")},
				{No: 14, StageNo: 1, IsLopt: true},
				{No: 15, StageNo: 1, Str: testutil.Ptr("watch")},
				{No: 16, StageNo: 1, IsArg: true},
				{No: 17, StageNo: 1, QuoteTypeSignal: argtables.SingleQuote},
				{No: 18, StageNo: 1, Str: testutil.Ptr("my-pod")},
			},
			wantCtrl: ControlModel{
				IsLiveStdout: true,
				IsLiveStderr: true,
				IsVersion:    false,
				IsHelp:       false,
			},
			wantStMod: []StageModel{
				{
					No:   1,
					Desc: "service-action-thorough",
					Cmd:  "kubectl",
					Svc:  testutil.Ptr("pods"),
					SvcOps: []OptParam{
						{
							Index:  7,
							OptStr: "n",
							Param: ParamType{
								Str:       testutil.Ptr("default"),
								QuoteType: argtables.NoQuote,
							},
						},
					},
					Act:     testutil.Ptr("get"),
					ActLops: []OptParam{{Index: 14, OptStr: "watch", Param: ParamType{}}},
					ActArgs: []ArgParam{
						{
							Index: 17,
							Param: ParamType{
								Str:       testutil.Ptr("my-pod"),
								QuoteType: argtables.SingleQuote,
							},
						},
					},
				},
			},
		},
		{
			name: "should handle trailing empty IsValue and IsArg safely",
			input: []argtables.ArgTable{
				{No: 1, StageNo: 1, IsStage: true},
				{No: 2, StageNo: 1, Str: testutil.Ptr("trailing-empty")},
				{No: 3, StageNo: 1, IsCmd: true},
				{No: 4, StageNo: 1, Str: testutil.Ptr("testcmd")},
				{No: 5, StageNo: 1, IsOpt: true},
				{No: 6, StageNo: 1, Str: testutil.Ptr("o")},
				{No: 7, StageNo: 1, IsValue: true}, // Interpreted as an empty argument
				{No: 8, StageNo: 1, IsArg: true},   // Interpreted as another empty argument
			},
			wantCtrl: ControlModel{
				IsLiveStdout: true,
				IsLiveStderr: true,
				IsVersion:    false,
				IsHelp:       false,
			},
			wantStMod: []StageModel{
				{
					No:   1,
					Desc: "trailing-empty",
					Cmd:  "testcmd",
					CmdOps: []OptParam{
						{
							Index:  5,
							OptStr: "o",
							Param:  ParamType{},
						},
					},
					CmdArgs: []ArgParam{
						{
							Index: 7,
							Param: ParamType{},
						},
					},
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
