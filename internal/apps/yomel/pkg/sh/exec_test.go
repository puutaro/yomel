package sh

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Exec verifies that Exec executes pipelines correctly under various configurations, failure states, and exit statuses.
func Test_Exec(t *testing.T) {
	tests := []struct {
		name           string
		yomelInfo      YomelInfo
		wantSubstr     string
		expectedStatus int
	}{
		{
			name: "should do nothing and return ExitSuccess when stageInfos is empty",
			yomelInfo: YomelInfo{
				IsDirect:   false,
				StageInfos: []StageInfo{},
			},
			wantSubstr:     "",
			expectedStatus: ExitSuccess,
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
			wantSubstr:     "",
			expectedStatus: ExitSuccess,
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
			wantSubstr:     "",
			expectedStatus: ExitSuccess,
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
			wantSubstr:     "YOMEL-LOG",
			expectedStatus: ExitSuccess,
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
			wantSubstr:     "",
			expectedStatus: ExitSuccess,
		},
		{
			name: "should print log and return failure exit status when command fails in normal mode",
			yomelInfo: YomelInfo{
				IsDirect: false,
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "error-stage",
						CmdStrs: "echo 'error message' 1>&2; exit 42",
					},
				},
			},
			wantSubstr:     "YOMEL-LOG",
			expectedStatus: 42,
		},
		{
			name: "should execute successfully with live stdout and live stderr enabled",
			yomelInfo: YomelInfo{
				IsDirect:     false,
				IsLiveStdout: true,
				IsLiveStdErr: true,
				StageInfos: []StageInfo{
					{
						No:           1,
						Desc:         "live-stage",
						CmdStrs:      "echo 'live content' 1>&2",
						IsLog:        true,
						ErrLogFilter: "grep 'live'",
					},
				},
			},
			wantSubstr:     "YOMEL-LOG",
			expectedStatus: ExitSuccess,
		},
		{
			name: "should print title log and return error status when yomel title is provided and command errors",
			yomelInfo: YomelInfo{
				IsDirect: false,
				Title:    "pipeline-test-title",
				StageInfos: []StageInfo{
					{
						No:      1,
						Desc:    "title-error-stage",
						CmdStrs: "exit 1",
					},
				},
			},
			wantSubstr:     "YOMEL-LOG-TITLE:",
			expectedStatus: ExitErrGeneral,
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

			status := Exec(tt.yomelInfo)

			wOut.Close()
			wErr.Close()
			os.Stdout = oldStdout
			os.Stderr = oldStderr

			var bufErr bytes.Buffer
			_, _ = bufErr.ReadFrom(rErr)
			errOutput := bufErr.String()

			assert.Equal(t, tt.expectedStatus, status)

			if tt.wantSubstr != "" {
				assert.True(t, strings.Contains(errOutput, tt.wantSubstr))
			}
		})
	}
}
