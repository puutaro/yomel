package argtablesvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

// TestCheckStageParameterSpecifyInCtrlErr tests validation for stage parameters specified in control section.
func TestCheckStageParameterSpecifyInCtrlErr(t *testing.T) {
	tests := []struct {
		name      string
		inputArgs []string
		wantErr   bool
	}{
		{
			name:      "Valid arg tables without stage parameters in ctrl section",
			inputArgs: []string{"stage", "desc1", "-cmd", "echo 1"},
			wantErr:   false,
		},
		{
			name:      "Invalid cmd parameter specified in ctrl section",
			inputArgs: []string{"-cmd", "echo 1", "stage", "desc1", "-cmd", "echo 1"},
			wantErr:   true,
		},
		{
			name:      "Invalid arg parameter specified in ctrl section",
			inputArgs: []string{"--arg", "val1", "stage", "desc1", "-cmd", "echo 1"},
			wantErr:   true,
		},
		{
			name:      "Invalid opt parameter specified in ctrl section",
			inputArgs: []string{"--opt", "val1", "stage", "desc1", "-cmd", "echo 1"},
			wantErr:   true,
		},
		{
			name:      "Invalid lopt parameter specified in ctrl section",
			inputArgs: []string{"--lop", "val1", "stage", "desc1", "-cmd", "echo 1"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argTables := argtables.GenArgTable(tt.inputArgs)
			err := ArgTableValidate(argTables)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgTableValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
