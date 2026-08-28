package model_test

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
)

// TestParseArg tests the argument parsing functionality.
func TestParseArg(hr *testing.T) {
	// Setup test cases for parseArg via Parse function
	strVal := "echo"
	argVal := "hello"

	argTables := []argtables.ArgTable{
		{
			No:      1,
			StageNo: 1,
			IsStage: true,
			Str:     &strVal,
		},
		{
			No:      2,
			StageNo: 1,
			IsCmd:   true,
			Str:     &strVal,
		},
		{
			No:      3,
			StageNo: 1,
			IsArg:   true,
			Comment: "Arg1",
			Str:     &argVal,
		},
	}

	_, stModels := model.Parse(argTables)
	if len(stModels) != 1 {
		hr.Errorf("Expected 1 stage model, got %d", len(stModels))
	}
}
