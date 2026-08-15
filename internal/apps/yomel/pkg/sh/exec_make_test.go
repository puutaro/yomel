package sh

import (
	"bytes"
	"testing"
	"time"
)

func TestYomelLogMake(t *testing.T) {
	startTime := time.Now()

	t.Run("Normal case with single stage and no log", func(t *testing.T) {
		yl := yomelLog{
			yomelInfo: YomelInfo{
				Title: "Test Title",
				StageInfos: []StageInfo{
					{
						No:                 1,
						Desc:               "First stage",
						IsLog:              true,
						CmdStrs:            "echo hello",
						CmdStrsWithComment: "echo hello",
					},
				},
				ForegroundColor: "",
				BackgroundColor: "",
			},
			stageInfos: []StageInfo{
				{
					No:                 1,
					Desc:               "First stage",
					IsLog:              true,
					CmdStrs:            "echo hello",
					CmdStrsWithComment: "echo hello",
				},
			},
			stdoutBuffers:    []*bytes.Buffer{bytes.NewBufferString("hello\n")},
			stderrBuffers:    []*bytes.Buffer{bytes.NewBufferString("")},
			cmdHasError:      false,
			startTime:        startTime,
			stageDurations:   []time.Duration{time.Second},
			isTerminal:       false,
			isLightColorMode: false,
		}
		buf := yl.make()
		if buf.Len() == 0 {
			t.Error("expected log output, got empty")
		}
	})

	t.Run("Multiple stages with terminal color and error", func(t *testing.T) {
		yl := yomelLog{
			yomelInfo: YomelInfo{
				Title: "Multi Stage Title",
				StageInfos: []StageInfo{
					{
						No:                 1,
						Desc:               "Stage one",
						IsLog:              true,
						CmdStrs:            "pwd",
						CmdStrsWithComment: "pwd",
					},
					{
						No:                 2,
						Desc:               "Stage two",
						IsLog:              true,
						CmdStrs:            "ls",
						CmdStrsWithComment: "ls",
					},
				},
				ForegroundColor: "\x1b[31m",
				BackgroundColor: "\x1b[41m",
			},
			stageInfos: []StageInfo{
				{
					No:                 1,
					Desc:               "Stage one",
					IsLog:              true,
					CmdStrs:            "pwd",
					CmdStrsWithComment: "pwd",
				},
				{
					No:                 2,
					Desc:               "Stage two",
					IsLog:              true,
					CmdStrs:            "ls",
					CmdStrsWithComment: "ls",
				},
			},
			stdoutBuffers: []*bytes.Buffer{
				bytes.NewBufferString("/tmp\n"),
				bytes.NewBufferString("file.txt\n"),
			},
			stderrBuffers: []*bytes.Buffer{
				bytes.NewBufferString(""),
				bytes.NewBufferString("some error\n"),
			},
			cmdHasError:      true,
			startTime:        startTime,
			stageDurations:   []time.Duration{time.Second, 2 * time.Second},
			isTerminal:       true,
			isLightColorMode: true,
		}
		buf := yl.make()
		if buf.Len() == 0 {
			t.Error("expected non-empty log output for multi stage with error")
		}
	})

	t.Run("Filter shell handling and light mode toggles", func(t *testing.T) {
		yl := yomelLog{
			yomelInfo: YomelInfo{
				Title: "",
				StageInfos: []StageInfo{
					{
						No:                 1,
						Desc:               "Filter stage",
						IsLog:              true,
						CmdStrs:            "echo test",
						CmdStrsWithComment: "echo test",
						LogFilter:          "cat",
						ErrLogFilter:       "cat",
					},
				},
			},
			stageInfos: []StageInfo{
				{
					No:                 1,
					Desc:               "Filter stage",
					IsLog:              true,
					CmdStrs:            "echo test",
					CmdStrsWithComment: "echo test",
					LogFilter:          "cat",
					ErrLogFilter:       "cat",
				},
			},
			stdoutBuffers:    []*bytes.Buffer{bytes.NewBufferString("filtered out\n")},
			stderrBuffers:    []*bytes.Buffer{bytes.NewBufferString("err filtered\n")},
			cmdHasError:      false,
			startTime:        startTime,
			stageDurations:   []time.Duration{time.Millisecond * 500},
			isTerminal:       true,
			isLightColorMode: false,
		}
		buf := yl.make()
		if buf.Len() == 0 {
			t.Error("expected output from filter test case")
		}
	})
}

func TestExecHelpers(t *testing.T) {
	t.Run("Direct exec with empty command", func(t *testing.T) {
		code := directExec("")
		if code != ExitSuccess {
			t.Errorf("expected ExitSuccess, got %d", code)
		}
	})

	t.Run("Colorize pipeline command", func(t *testing.T) {
		res := colorizePipelineCmd("echo hello \\\n| cat", "\x1b[32m", true)
		if res == "" {
			t.Error("expected colored pipeline command")
		}

		resEmpty := colorizePipelineCmd("", "\x1b[32m", true)
		if resEmpty != "" {
			t.Error("expected empty string")
		}
	})

	t.Run(ansiColorTestsIfNeeded, func(t *testing.T) {
		stripped := stripAnsi("\x1b[31mhello\x1b[39m")
		if stripped != "hello" {
			t.Errorf("expected 'hello', got '%s'", stripped)
		}
	})
}

const ansiColorTestsIfNeeded = "Strip ansi and color functions"
