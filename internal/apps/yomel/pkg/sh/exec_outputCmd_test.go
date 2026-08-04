// Write direct above line for Comment on code
package sh

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_outputCmd verifies that outputCmd correctly prints the generated pipeline command to stdout.
func Test_outputCmd(t *testing.T) {
	tests := []struct {
		name       string
		stageInfos []StageInfo
		want       string
	}{
		{
			name: "should print single stage command to stdout",
			stageInfos: []StageInfo{
				{
					No:      1,
					Desc:    "echo-stage",
					CmdStrs: "echo 'hello gen'",
				},
			},
			want: "echo 'hello gen'\n",
		},
		{
			name: "should print multi-stage command pipeline to stdout",
			stageInfos: []StageInfo{
				{
					No:      1,
					Desc:    "source-stage",
					CmdStrs: "echo 'line1'",
				},
				{
					No:      2,
					Desc:    "grep-stage",
					CmdStrs: "grep 'line1'",
				},
			},
			want: "echo 'line1' \\\n| grep 'line1'\n",
		},
		{
			name:       "should print empty line when stageInfos is empty",
			stageInfos: []StageInfo{},
			want:       "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			rOut, wOut, _ := os.Pipe()
			os.Stdout = wOut

			outputCmd(tt.stageInfos)

			wOut.Close()
			os.Stdout = oldStdout

			var bufOut bytes.Buffer
			_, _ = bufOut.ReadFrom(rOut)
			got := bufOut.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
