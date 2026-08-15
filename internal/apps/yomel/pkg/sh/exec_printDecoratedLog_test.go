package sh

import (
	"bytes"
	"testing"
	"time"
)

func Test_printDecoratedLog(t *testing.T) {
	tests := []struct {
		name             string
		isTerminal       bool
		isLightColorMode bool
		cmdHasError      bool
		shouldLog        bool
		stageInfo        StageInfo
		yomelInfo        YomelInfo
	}{
		{
			name:             "Terminal mode enabled with light color mode and error",
			isTerminal:       true,
			isLightColorMode: true,
			cmdHasError:      true,
			shouldLog:        true,
			stageInfo: StageInfo{
				No:                 1,
				Desc:               "Test stage description",
				IsLog:              true,
				LogFilter:          "",
				ErrLogFilter:       "",
				CmdStrs:            "echo 'hello'",
				CmdStrsWithComment: "echo 'hello' # comment",
				ForegroundColor:    "\x1b[31m",
				BackgroundColor:    "\x1b[42m",
				CommentColorStart:  "\x1b[34m",
			},
			yomelInfo: YomelInfo{
				Title:           "Test Title",
				ForegroundColor: "\x1b[31m",
				BackgroundColor: "\x1b[42m",
			},
		},
		{
			name:             "Non-terminal mode without error and log filter",
			isTerminal:       false,
			isLightColorMode: false,
			cmdHasError:      false,
			shouldLog:        false,
			stageInfo: StageInfo{
				No:                 2,
				Desc:               "Another stage",
				IsLog:              false,
				LogFilter:          "cat",
				ErrLogFilter:       "cat",
				CmdStrs:            "ls",
				CmdStrsWithComment: "ls",
			},
			yomelInfo: YomelInfo{
				Title: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yl := yomelLog{
				yomelInfo:        tt.yomelInfo,
				stageInfos:       []StageInfo{tt.stageInfo},
				stdoutBuffers:    []*bytes.Buffer{bytes.NewBufferString("test stdout\n")},
				stderrBuffers:    []*bytes.Buffer{bytes.NewBufferString("test stderr\n")},
				cmdHasError:      tt.cmdHasError,
				stageDurations:   []time.Duration{time.Second},
				isTerminal:       tt.isTerminal,
				isLightColorMode: tt.isLightColorMode,
			}

			var buf bytes.Buffer
			yl.printDecoratedLog(
				&buf,
				tt.stageInfo,
				0,
				yl.stderrBuffers[0],
				yl.stdoutBuffers[0],
				tt.shouldLog,
				"\x1b[41m",
			)

			if buf.Len() == 0 {
				t.Errorf("printDecoratedLog() produced empty output")
			}
		})
	}
}
