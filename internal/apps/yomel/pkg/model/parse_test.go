package model_test

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		argTables []argtables.ArgTable
		check     func(t *testing.T, ctrl model.ControlModel, stModels []model.StageModel)
	}{
		{
			name:      "Parse empty argtables",
			argTables: []argtables.ArgTable{},
			check: func(t *testing.T, ctrl model.ControlModel, stModels []model.StageModel) {
				if len(stModels) != 0 {
					t.Errorf("expected 0 stages, got %d", len(stModels))
				}
			},
		},
		{
			name: "Parse basic stage",
			argTables: []argtables.ArgTable{
				{
					No:      1,
					StageNo: 1,
					IsStage: true,
				},
				{
					No:      1,
					StageNo: 1,
					Str:     testutil.Ptr("test-stage"),
				},
				{
					No:      2,
					StageNo: 1,
					IsCmd:   true,
				},
				{
					No:      2,
					StageNo: 1,
					Str:     testutil.Ptr("echo"),
				},
			},
			check: func(t *testing.T, ctrl model.ControlModel, stModels []model.StageModel) {
				if len(stModels) != 1 {
					t.Fatalf("expected 1 stage, got %d", len(stModels))
				}
				if stModels[0].Cmd != "echo" {
					t.Errorf("expected cmd 'echo', got '%s'", stModels[0].Cmd)
				}
				if stModels[0].Desc != "test-stage" {
					t.Errorf("expected desc 'test-stage', got '%s'", stModels[0].Desc)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, stModels := model.Parse(tt.argTables)
			tt.check(t, ctrl, stModels)
		})
	}
}
