// Test_yomelLog_make verifies that yomelLog.make correctly generates the decorated log output based on stages, titles, and error states.
package sh

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_yomelLog_make(t *testing.T) {
	tests := []struct {
		name             string
		yomelInfo        YomelInfo
		stageInfos       []StageInfo
		stdoutStrs       []string
		stderrStrs       []string
		cmdHasError      bool
		wantOutputSubstr []string
	}{
		{
			name: "should return empty buffer when no stage requires logging and no error occurred",
			yomelInfo: YomelInfo{
				Title: "test-title",
			},
			stageInfos: []StageInfo{
				{
					No:    1,
					IsLog: false,
				},
			},
			stdoutStrs:       []string{"hello\n"},
			stderrStrs:       []string{""},
			cmdHasError:      false,
			wantOutputSubstr: []string{},
		},
		{
			name: "should generate decorated log when single stage has IsLog true",
			yomelInfo: YomelInfo{
				Title: "single-title",
			},
			stageInfos: []StageInfo{
				{
					No:                 1,
					Desc:               "single-desc",
					CmdStrsWithComment: "echo 'hello'",
					IsLog:              true,
				},
			},
			stdoutStrs:  []string{"hello\n"},
			stderrStrs:  []string{""},
			cmdHasError: false,
			wantOutputSubstr: []string{
				"Yomel-log_",
				"Single-desc",
				"Cmd",
				"echo 'hello'",
				"Stdout",
				"hello",
			},
		},
		{
			name: "should include title and total command when multiple stages are present and should log",
			yomelInfo: YomelInfo{
				Title: "multi-title",
			},
			stageInfos: []StageInfo{
				{
					No:                 1,
					Desc:               "stage-1",
					CmdStrsWithComment: "echo 'line1'",
					IsLog:              true,
				},
				{
					No:                 2,
					Desc:               "stage-2",
					CmdStrsWithComment: "grep 'line1'",
					IsLog:              true,
				},
			},
			stdoutStrs:  []string{"line1\n", "line1\n"},
			stderrStrs:  []string{"", ""},
			cmdHasError: false,
			wantOutputSubstr: []string{
				"Yomel-log_",
				"Title",
				"Multi-title",
				"Total-cmd",
				"echo 'line1'",
				"grep 'line1'",
				"Stage[1]_",
				"Stage[2]_",
			},
		},
		{
			name: "should force logging and display error header when cmdHasError is true",
			yomelInfo: YomelInfo{
				Title: "error-title",
			},
			stageInfos: []StageInfo{
				{
					No:                 1,
					Desc:               "error-stage",
					CmdStrsWithComment: "exit 1",
					IsLog:              false,
				},
			},
			stdoutStrs:  []string{""},
			stderrStrs:  []string{"some error occurred\n"},
			cmdHasError: true,
			wantOutputSubstr: []string{
				"Yomel-log_",
				"Error-stage",
				"Error",
				"some error occurred",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdoutBuffers := make([]*bytes.Buffer, len(tt.stdoutStrs))
			for i, s := range tt.stdoutStrs {
				stdoutBuffers[i] = bytes.NewBufferString(s)
			}

			stderrBuffers := make([]*bytes.Buffer, len(tt.stderrStrs))
			for i, s := range tt.stderrStrs {
				stderrBuffers[i] = bytes.NewBufferString(s)
			}

			yl := &yomelLog{
				yomelInfo:      tt.yomelInfo,
				stageInfos:     tt.stageInfos,
				stdoutBuffers:  stdoutBuffers,
				stderrBuffers:  stderrBuffers,
				cmdHasError:    tt.cmdHasError,
				startTime:      time.Now(),
				stageDurations: []time.Duration{time.Second, time.Second},
			}

			buffer := yl.make()
			output := stripANSI(buffer.String())

			if len(tt.wantOutputSubstr) == 0 {
				assert.Empty(t, output)
				return
			}

			for _, substr := range tt.wantOutputSubstr {
				assert.Contains(t, output, substr)
			}
		})
	}
}
