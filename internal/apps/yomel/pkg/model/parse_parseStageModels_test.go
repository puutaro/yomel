package model_test

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
)

func TestParseStageModels(t *testing.T) {
	// Define test argtables for stage parsing
	argTables := []argtables.ArgTable{
		{
			No:      1,
			StageNo: 1,
			IsStage: true,
			Str:     testutil.Ptr("first stage"),
		},
		{
			No:      2,
			StageNo: 1,
			IsCmd:   true,
		},
		{
			No:      3,
			StageNo: 1,
			Str:     testutil.Ptr("echo hello"),
		},
		{
			No:          4,
			StageNo:     1,
			IsLogFilter: true,
		},
		{
			No:      5,
			StageNo: 1,
			Str:     testutil.Ptr("grep hello"),
		},
	}

	// Parse argtables into control and stage models
	_, stModels := model.Parse(argTables)
	if len(stModels) != 1 {
		t.Fatalf("expected 1 stage model, got %d", len(stModels))
	}
	if stModels[0].Cmd != "echo hello" {
		t.Errorf("expected cmd 'echo hello', got '%s'", stModels[0].Cmd)
	}
	if stModels[0].LogFilter != "grep hello" {
		t.Errorf("expected log filter 'grep hello', got '%s'", stModels[0].LogFilter)
	}
}
