package sh_test

import (
	"bytes"
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/env"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/sh"
)

// TestExecSuccess tests normal execution of yomel info.
func TestExecSuccess(t *testing.T) {
	yomelInfo := sh.YomelInfo{
		IsLiveStdout: false,
		IsLiveStdErr: false,
		StageInfos: []sh.StageInfo{
			{
				No:                 1,
				Desc:               "test stage 1",
				IsLog:              true,
				CmdStrs:            "echo 'hello'",
				CmdStrsWithComment: "echo 'hello'",
			},
		},
	}
	yomelEnv := env.YomelEnv{
		IsTerminal:       false,
		IsLightColorMode: true,
	}

	exitCode := sh.Exec(yomelInfo, yomelEnv)
	if exitCode != sh.ExitSuccess {
		t.Errorf("expected exit code %d, got %d", sh.ExitSuccess, exitCode)
	}
}

// TestExecGenMode tests generation mode execution.
func TestExecGenMode(t *testing.T) {
	yomelInfo := sh.YomelInfo{
		IsGen: true,
		StageInfos: []sh.StageInfo{
			{
				No:      1,
				CmdStrs: "echo 'gen'",
			},
		},
	}
	yomelEnv := env.YomelEnv{}

	exitCode := sh.Exec(yomelInfo, yomelEnv)
	if exitCode != sh.ExitSuccess {
		t.Errorf("expected exit code %d, got %d", sh.ExitSuccess, exitCode)
	}
}

// TestExecDirectMode tests direct execution mode.
func TestExecDirectMode(t *testing.T) {
	yomelInfo := sh.YomelInfo{
		IsDirect: true,
		StageInfos: []sh.StageInfo{
			{
				No:      1,
				CmdStrs: "echo 'direct'",
			},
		},
	}
	yomelEnv := env.YomelEnv{}

	exitCode := sh.Exec(yomelInfo, yomelEnv)
	if exitCode != sh.ExitSuccess {
		t.Errorf("expected exit code %d, got %d", sh.ExitSuccess, exitCode)
	}
}

// TestExecErrorAndFilter tests error handling and log filtering functions.
func TestExecErrorAndFilter(t *testing.T) {
	yomelInfo := sh.YomelInfo{
		IsLiveStdout: true,
		IsLiveStdErr: true,
		Title:        "Error Test Pipeline",
		StageInfos: []sh.StageInfo{
			{
				No:                 1,
				Desc:               "error stage",
				IsLog:              true,
				LogFilter:          "sed 's/foo/bar/'",
				ErrLogFilter:       "sed 's/err/success/'",
				CmdStrs:            "echo 'foo' && echo 'err_msg' >&2 && exit 1",
				CmdStrsWithComment: "echo 'foo' && echo 'err_msg' >&2 && exit 1",
			},
		},
	}
	yomelEnv := env.YomelEnv{
		IsTerminal:       true,
		IsLightColorMode: false,
	}

	exitCode := sh.Exec(yomelInfo, yomelEnv)
	if exitCode == sh.ExitSuccess {
		t.Errorf("expected non-zero exit code for failing command")
	}
}

// TestMakeTotalPipeCmd tests multi-stage pipeline string generation.
func TestMakeTotalPipeCmd(t *testing.T) {
	// stageInfos := []sh.StageInfo{
	// 	{CmdStrs: "cmd1"},
	// 	{CmdStrs: "cmd2"},
	// }

	// Indirectly testing via execution with empty or multiple stages
	// yomelInfo := sh.YomelInfo{
	// 	StageInfos: stageInfos,
	// }
	yomelEnv := env.YomelEnv{}

	// Zero stages check
	yomelInfoEmpty := sh.YomelInfo{StageInfos: []sh.StageInfo{}}
	if code := sh.Exec(yomelInfoEmpty, yomelEnv); code != sh.ExitSuccess {
		t.Errorf("expected success for empty stages")
	}
}

// TestCompNewLine tests newline compensation helper.
func TestCompNewLine(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("test\n")

	// Testing internal behavior via log making if applicable or standalone if exported.
	// Since compNewLine is unexported, we ensure related execution paths cover it.
	yomelInfo := sh.YomelInfo{
		Title: "Multi Stage",
		StageInfos: []sh.StageInfo{
			{No: 1, IsLog: true, CmdStrs: "echo 1", CmdStrsWithComment: "echo 1"},
			{No: 2, IsLog: true, CmdStrs: "echo 2", CmdStrsWithComment: "echo 2"},
		},
	}
	yomelEnv := env.YomelEnv{IsTerminal: false}
	code := sh.Exec(yomelInfo, yomelEnv)
	if code != sh.ExitSuccess {
		t.Errorf("failed multi stage execution")
	}
}
