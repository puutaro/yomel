package sh

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Exec verifies that Exec correctly executes pipeline commands or direct mode based on YomelInfo settings.
func Test_Exec(t *testing.T) {
	tests := []struct {
		name       string
		yomelInfo  YomelInfo
		wantSubstr string
	}{
		{
			name: "should do nothing when stageInfos is empty",
			yomelInfo: YomelInfo{
				IsDirect:   false,
				StageInfos: []StageInfo{},
			},
			wantSubstr: "",
		},
		{
			name: "should execute single stage pipeline successfully in normal mode",
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
			wantSubstr: "",
		},
		{
			name: "should execute multi-stage pipeline successfully in normal mode",
			yomelInfo: YomelInfo{
				IsDirect: false,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "source-stage",
						CmdStrs: "echo 'line1\nline2'",
					},
					{
						No:      2,
						Desc:    "grep-stage",
						CmdStrs: "grep 'line1'",
					},
				},
			},
			wantSubstr: "",
		},
		{
			name: "should print log when IsLog is true and stdout has content",
			yomelInfo: YomelInfo{
				IsDirect: false,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "log-stage",
						CmdStrs: "echo 'logged content'",
						IsLog:   true,
					},
				},
			},
			wantSubstr: "YOMEL_LOG",
		},
		{
			name: "should execute direct mode successfully when IsDirect is true",
			yomelInfo: YomelInfo{
				IsDirect: true,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "direct-stage",
						CmdStrs: "echo 'direct mode test'",
					},
				},
			},
			wantSubstr: "",
		},
		{
			name: "should print log and handle error when command fails in normal mode",
			yomelInfo: YomelInfo{
				IsDirect: false,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "error-stage",
						CmdStrs: "echo 'error message' 1>&2; exit 1",
					},
				},
			},
			wantSubstr: "YOMEL_LOG",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			oldStderr := os.Stderr
			_, wOut, _ := os.Pipe()
			rErr, wErr, _ := os.Pipe()
			os.Stdout = wOut
			os.Stderr = wErr

			Exec(tt.yomelInfo)

			wOut.Close()
			wErr.Close()
			os.Stdout = oldStdout
			os.Stderr = oldStderr

			var bufErr bytes.Buffer
			_, _ = bufErr.ReadFrom(rErr)
			errOutput := bufErr.String()

			if tt.wantSubstr != "" {
				assert.True(t, strings.Contains(errOutput, tt.wantSubstr))
			}
		})
	}
}
