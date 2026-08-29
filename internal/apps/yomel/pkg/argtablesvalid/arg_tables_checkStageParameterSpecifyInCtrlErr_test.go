package argtablesvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

func TestCheckStageParameterSpecifyInCtrlErr(t *testing.T) {
	tests := []struct {
		name      string
		inputArgs []string
		wantErr   bool
	}{
		{
			name:      "Valid arg tables without stage parameters in ctrl section",
			inputArgs: []string{"//", "desc1", "-c", "echo 1"},
			wantErr:   false,
		},
		{
			name:      "Invalid cmd parameter specified in ctrl section",
			inputArgs: []string{"-c", "echo 1", "//", "desc1", "-c", "echo 1"},
			wantErr:   true,
		},
		{
			name:      "Invalid arg parameter specified in ctrl section",
			inputArgs: []string{"-a", "val1", "//", "desc1", "-c", "echo 1"},
			wantErr:   true,
		},
		{
			name:      "Invalid opt parameter specified in ctrl section",
			inputArgs: []string{"-o", "val1", "//", "desc1", "-c", "echo 1"},
			wantErr:   true,
		},
		{
			name:      "Invalid lopt parameter specified in ctrl section",
			inputArgs: []string{"--o", "val1", "//", "desc1", "-c", "echo 1"},
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
