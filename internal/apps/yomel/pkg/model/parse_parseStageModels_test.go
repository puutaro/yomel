package model

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

// TestParse_ParseStageModels tests the stage parsing functionality in Parse function.
func TestParse_ParseStageModels(t *testing.T) {
	tests := []struct {
		name       string
		argTables  []argtables.ArgTable
		wantStages []StageModel
		wantCtrl   ControlModel
	}{
		{
			name:       "empty argTables",
			argTables:  []argtables.ArgTable{},
			wantStages: []StageModel{},
			wantCtrl:   ControlModel{},
		},
		{
			name: "single stage with basic command",
			argTables: []argtables.ArgTable{
				{No: 1, StageNo: 1, IsStage: true},
				{No: 2, StageNo: 1, Str: testutil.Ptr("first stage")},
				{No: 3, StageNo: 1, IsCmd: true},
				{No: 4, StageNo: 1, Str: testutil.Ptr("echo")},
				{No: 5, StageNo: 1, IsArg: true},
				{No: 6, StageNo: 1, Str: testutil.Ptr("hello")},
			},
			wantStages: []StageModel{
				{
					No:   1,
					Desc: "first stage",
					Cmd:  "echo",
					CmdArgs: []ArgParam{
						{
							Index: 6,
							Param: ParamType{
								Str: testutil.Ptr("hello"),
							},
						},
					},
					IsLog: testutil.Ptr(true),
				},
			},
			wantCtrl: ControlModel{
				IsLiveStdout: true,
				IsLiveStderr: true,
				IsLog:        testutil.Ptr(true),
			},
		},
		{
			name: "multiple stages with options and arguments",
			argTables: []argtables.ArgTable{
				{No: 1, StageNo: 0, IsTitle: true},
				{No: 2, StageNo: 0, Str: testutil.Ptr("pipeline title")},
				{No: 3, StageNo: 1, IsStage: true},
				{No: 4, StageNo: 1, Str: testutil.Ptr("stage one")},
				{No: 5, StageNo: 1, IsCmd: true},
				{No: 6, StageNo: 1, Str: testutil.Ptr("grep")},
				{No: 7, StageNo: 1, IsOpt: true, Comment: "Pattern"},
				{No: 8, StageNo: 1, IsValue: true},
				{No: 9, StageNo: 1, Str: testutil.Ptr("foo")},
				{No: 10, StageNo: 2, IsStage: true},
				{No: 11, StageNo: 2, Str: testutil.Ptr("stage two")},
				{No: 12, StageNo: 2, IsCmd: true},
				{No: 13, StageNo: 2, Str: testutil.Ptr("wc")},
				{No: 14, StageNo: 2, IsLopt: true, Comment: "Lines"},
			},
			wantStages: []StageModel{
				{
					No:   1,
					Desc: "stage one",
					Cmd:  "grep",
					CmdOps: []OptParam{
						{
							Index:   9,
							OptStr:  "foo",
							Comment: "Pattern",
							Param: ParamType{
								Str: testutil.Ptr("foo"),
							},
						},
					},
					IsLog: testutil.Ptr(true),
				},
				{
					No:   2,
					Desc: "stage two",
					Cmd:  "wc",
					CmdLops: []OptParam{
						{
							Index:   10,
							OptStr:  "",
							Comment: "Lines",
						},
					},
					IsLog: testutil.Ptr(true),
				},
			},
			wantCtrl: ControlModel{
				Title:        "pipeline title",
				IsLiveStdout: true,
				IsLiveStderr: true,
				IsLog:        testutil.Ptr(true),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, stages := Parse(tt.argTables)

			if len(stages) != len(tt.wantStages) {
				t.Fatalf("Parse() got %d stages, want %d stages", len(stages), len(tt.wantStages))
			}

			for i, gotStage := range stages {
				wantStage := tt.wantStages[i]
				if gotStage.No != wantStage.No {
					t.Errorf("stage[%d] No = %v, want %v", i, gotStage.No, wantStage.No)
				}
				if gotStage.Desc != wantStage.Desc {
					t.Errorf("stage[%d] Desc = %v, want %v", i, gotStage.Desc, wantStage.Desc)
				}
				if gotStage.Cmd != wantStage.Cmd {
					t.Errorf("stage[%d] Cmd = %v, want %v", i, gotStage.Cmd, wantStage.Cmd)
				}
			}

			if ctrl.Title != tt.wantCtrl.Title {
				t.Errorf("ControlModel.Title = %v, want %v", ctrl.Title, tt.wantCtrl.Title)
			}
			if ctrl.IsLiveStdout != tt.wantCtrl.IsLiveStdout {
				t.Errorf("ControlModel.IsLiveStdout = %v, want %v", ctrl.IsLiveStdout, tt.wantCtrl.IsLiveStdout)
			}
			if ctrl.IsLiveStderr != tt.wantCtrl.IsLiveStderr {
				t.Errorf("ControlModel.IsLiveStderr = %v, want %v", ctrl.IsLiveStderr, tt.wantCtrl.IsLiveStderr)
			}
		})
	}
}
