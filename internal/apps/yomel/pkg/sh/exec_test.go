// Write direct above line for Comment on code
package sh

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Exec verifies that Exec handles empty stage pipelines, normal mode execution, and direct mode execution correctly.
func Test_Exec(t *testing.T) {
	tests := []struct {
		name           string
		yomelInfo      YomelInfo
		expectedStatus int
	}{
		{
			name: "should return ExitSuccess when stageInfos is empty in normal mode",
			yomelInfo: YomelInfo{
				IsDirect:   false,
				StageInfos: []StageInfo{},
			},
			expectedStatus: ExitSuccess,
		},
		{
			name: "should return ExitSuccess when stageInfos is empty in direct mode",
			yomelInfo: YomelInfo{
				IsDirect:   true,
				StageInfos: []StageInfo{},
			},
			expectedStatus: ExitSuccess,
		},
		{
			name: "should execute single stage successfully and return ExitSuccess in normal mode",
			yomelInfo: YomelInfo{
				IsDirect: false,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "echo-stage",
						CmdStrs: "echo 'hello normal'",
					},
				},
			},
			expectedStatus: ExitSuccess,
		},
		{
			name: "should execute single stage successfully and return ExitSuccess in direct mode",
			yomelInfo: YomelInfo{
				IsDirect: true,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "echo-direct",
						CmdStrs: "echo 'hello direct'",
					},
				},
			},
			expectedStatus: ExitSuccess,
		},
		{
			name: "should return failure exit status when command fails in normal mode",
			yomelInfo: YomelInfo{
				IsDirect: false,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "error-stage",
						CmdStrs: "exit 42",
					},
				},
			},
			expectedStatus: 42,
		},
		{
			name: "should return failure exit status when command fails in direct mode",
			yomelInfo: YomelInfo{
				IsDirect: true,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "error-direct",
						CmdStrs: "exit 55",
					},
				},
			},
			expectedStatus: 55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			oldStderr := os.Stderr
			_, wOut, _ := os.Pipe()
			_, wErr, _ := os.Pipe()
			os.Stdout = wOut
			os.Stderr = wErr

			status := Exec(tt.yomelInfo)

			wOut.Close()
			wErr.Close()
			os.Stdout = oldStdout
			os.Stderr = oldStderr

			assert.Equal(t, tt.expectedStatus, status)
		})
	}
}
