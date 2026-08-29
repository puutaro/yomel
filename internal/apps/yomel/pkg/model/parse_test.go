package model_test

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

// TestParse verifies the parsing logic of argTables into ControlModel and StageModel.
func TestParse(t *testing.T) {
	tests := []struct {
		name          string
		argTables     []argtables.ArgTable
		expectedCtrl  model.ControlModel
		expectedStage []model.StageModel
	}{
		{
			name:      "Empty argTables returns empty models",
			argTables: []argtables.ArgTable{},
			expectedCtrl: model.ControlModel{
				IsLiveStdout: false,
				IsLiveStderr: false,
			},
			expectedStage: []model.StageModel{},
		},
		{
			name: "Control flags and single stage with command",
			argTables: []argtables.ArgTable{
				{No: 1, StageNo: 0, IsVersion: true},
				{No: 2, StageNo: 1, IsStage: true},
				{No: 4, StageNo: 1, Str: testutil.Ptr("stage1")},
				{No: 3, StageNo: 1, IsCmd: true},
				{No: 3, StageNo: 1, Str: testutil.Ptr("echo hello")},
			},
			expectedCtrl: model.ControlModel{
				IsVersion:    true,
				IsLiveStdout: true,
				IsLiveStderr: true,
				IsLog:        testutil.Ptr(true),
			},
			expectedStage: []model.StageModel{
				{
					No:    1,
					Desc:  "stage1",
					Cmd:   "echo hello",
					IsLog: testutil.Ptr(true),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, stages := model.Parse(tt.argTables)

			if ctrl.IsVersion != tt.expectedCtrl.IsVersion {
				t.Errorf("ControlModel.IsVersion = %v, want %v", ctrl.IsVersion, tt.expectedCtrl.IsVersion)
			}
			if ctrl.IsLiveStdout != tt.expectedCtrl.IsLiveStdout {
				t.Errorf("ControlModel.IsLiveStdout = %v, want %v", ctrl.IsLiveStdout, tt.expectedCtrl.IsLiveStdout)
			}
			if ctrl.IsLiveStderr != tt.expectedCtrl.IsLiveStderr {
				t.Errorf("ControlModel.IsLiveStderr = %v, want %v", ctrl.IsLiveStderr, tt.expectedCtrl.IsLiveStderr)
			}

			if len(stages) != len(tt.expectedStage) {
				t.Fatalf("Stages length = %d, want %d", len(stages), len(tt.expectedStage))
			}

			for i, st := range stages {
				exp := tt.expectedStage[i]
				if st.No != exp.No {
					t.Errorf("Stage[%d].No = %d, want %d", i, st.No, exp.No)
				}
				if st.Desc != exp.Desc {
					t.Errorf("Stage[%d].Desc = %s, want %s", i, st.Desc, exp.Desc)
				}
				if st.Cmd != exp.Cmd {
					t.Errorf("Stage[%d].Cmd = %s, want %s", i, st.Cmd, exp.Cmd)
				}
			}
		})
	}
}
