// Write direct above line for Comment on code
package sh

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_outputCmd verifies that outputCmd correctly prints the generated pipeline command string to stdout.
func Test_outputCmd(t *testing.T) {
	tests := []struct {
		name            string
		totalPipeCmdStr string
		want            string
	}{
		{
			name:            "should print single stage command string to stdout",
			totalPipeCmdStr: "echo 'hello gen'",
			want:            "echo 'hello gen'\n",
		},
		{
			name:            "should print multi-stage command pipeline string to stdout",
			totalPipeCmdStr: "echo 'line1' \\\n| grep 'line1'",
			want:            "echo 'line1' \\\n| grep 'line1'\n",
		},
		{
			name:            "should print empty line when totalPipeCmdStr is empty",
			totalPipeCmdStr: "",
			want:            "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			rOut, wOut, _ := os.Pipe()
			os.Stdout = wOut

			outputCmd(tt.totalPipeCmdStr)

			wOut.Close()
			os.Stdout = oldStdout

			var bufOut bytes.Buffer
			_, _ = bufOut.ReadFrom(rOut)
			got := bufOut.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
