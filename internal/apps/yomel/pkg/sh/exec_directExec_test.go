// Write direct above line for Comment on code
package sh

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_directExec verifies that directExec executes the pipeline command directly and returns the correct exit code.
func Test_directExec(t *testing.T) {
	tests := []struct {
		name            string
		totalPipeCmdStr string
		wantExitCode    int
		wantContains    string
	}{
		{
			name:            "should execute single stage pipeline successfully and return ExitSuccess",
			totalPipeCmdStr: "echo 'direct hello'",
			wantExitCode:    ExitSuccess,
			wantContains:    "direct hello",
		},
		{
			name:            "should return ExitSuccess when totalPipeCmdStr is empty",
			totalPipeCmdStr: "",
			wantExitCode:    ExitSuccess,
			wantContains:    "",
		},
		{
			name:            "should return error code when invalid command is executed",
			totalPipeCmdStr: "invalid_command_xyz",
			wantExitCode:    127,
			wantContains:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			oldStderr := os.Stderr
			rOut, wOut, _ := os.Pipe()
			rErr, wErr, _ := os.Pipe()
			os.Stdout = wOut
			os.Stderr = wErr

			gotExitCode := directExec(tt.totalPipeCmdStr)

			wOut.Close()
			wErr.Close()
			os.Stdout = oldStdout
			os.Stderr = oldStderr

			var bufOut bytes.Buffer
			var bufErr bytes.Buffer
			_, _ = bufOut.ReadFrom(rOut)
			_, _ = bufErr.ReadFrom(rErr)

			output := bufOut.String() + bufErr.String()

			assert.Equal(t, tt.wantExitCode, gotExitCode)
			if tt.wantContains != "" {
				assert.Contains(t, output, tt.wantContains)
			}
		})
	}
}
