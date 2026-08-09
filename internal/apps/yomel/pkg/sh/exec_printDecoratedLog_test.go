// Test_printDecoratedLog verifies that printDecoratedLog correctly outputs formatted logs to the writer under various conditions.
package sh

import (
	"bytes"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// stripANSI removes ANSI escape sequences (colors, underlines, bold, etc.) from the string.
func stripANSI(str string) string {
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[mK]`)
	return ansiRegex.ReplaceAllString(str, "")
}

func Test_printDecoratedLog(t *testing.T) {
	tests := []struct {
		name             string
		yomelInfo        YomelInfo
		stageInfo        StageInfo
		index            int
		stderrBuf        string
		stdoutBuf        string
		shouldLog        bool
		wantOutputSubstr []string
	}{
		{
			name:      "should print decorated log with stdout and progress stderr when shouldLog is true",
			yomelInfo: YomelInfo{},
			stageInfo: StageInfo{
				No:                 1,
				Desc:               "test-stage",
				CmdStrsWithComment: "echo 'hello'",
				IsLog:              true,
			},
			index:     0,
			stderrBuf: "some progress info\n",
			stdoutBuf: "hello\n",
			shouldLog: true,
			wantOutputSubstr: []string{
				"Stage[1]_",
				"Test-stage",
				"Cmd",
				"echo 'hello'",
				"Progress",
				"some progress info",
				"Stdout",
				"hello",
			},
		},
		{
			name: "should print decorated log with error label when cmdHasError is true",
			yomelInfo: YomelInfo{
				StageInfos: []StageInfo{
					{
						No:                 1,
						Desc:               "error-stage",
						CmdStrsWithComment: "exit 1",
					},
				},
			},
			stageInfo: StageInfo{
				No:                 1,
				Desc:               "error-stage",
				CmdStrsWithComment: "exit 1",
			},
			index:     0,
			stderrBuf: "error occurred\n",
			stdoutBuf: "",
			shouldLog: true,
			wantOutputSubstr: []string{
				"Stage[1]_",
				"Error-stage",
				"Cmd",
				"exit 1",
				"Error",
				"error occurred",
			},
		},
		{
			name:      "should print only stdout when stderr buffer is empty and shouldLog is true",
			yomelInfo: YomelInfo{},
			stageInfo: StageInfo{
				No:                 3,
				Desc:               "stdout-only-stage",
				CmdStrsWithComment: "ls",
				IsLog:              true,
			},
			index:     0,
			stderrBuf: "",
			stdoutBuf: "file1\nfile2\n",
			shouldLog: true,
			wantOutputSubstr: []string{
				"Stage[3]_",
				"Stdout-only-stage",
				"Cmd",
				"ls",
				"Stdout",
				"file1",
				"file2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var combinedLog bytes.Buffer
			stderrBuffer := bytes.NewBufferString(tt.stderrBuf)
			stdoutBuffer := bytes.NewBufferString(tt.stdoutBuf)

			yl := &yomelLog{
				yomelInfo:      tt.yomelInfo,
				stageInfos:     []StageInfo{tt.stageInfo},
				stdoutBuffers:  []*bytes.Buffer{stdoutBuffer},
				stderrBuffers:  []*bytes.Buffer{stderrBuffer},
				cmdHasError:    tt.name == "should print decorated log with error label when cmdHasError is true",
				startTime:      time.Now(),
				stageDurations: []time.Duration{time.Second},
			}

			yl.printDecoratedLog(
				&combinedLog,
				tt.stageInfo,
				tt.index,
				stderrBuffer,
				stdoutBuffer,
				tt.shouldLog,
			)

			output := stripANSI(combinedLog.String())

			for _, substr := range tt.wantOutputSubstr {
				assert.Contains(t, output, substr)
			}
		})
	}
}
