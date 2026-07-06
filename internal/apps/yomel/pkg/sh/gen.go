package sh

import (
	"strings"
)

const (
	opArgPrefixBlank            = " "
	backslashNewline            = "\\\n"
	backslashNewlineOpArgPrefix = opArgPrefixBlank + backslashNewline +
		opArgPrefixBlank
	verticalbar
)

type stageCommand string

type YomelInfo struct {
	No           int
	Desc         string
	IsLog        bool
	LogFilter    string
	ErrLogFilter string
	CmdStrs      string
}

func Gen(chain Yomel) []YomelInfo {
	stages := chain.Stages
	chainInfos := make([]YomelInfo, len(stages))
	globalLogFilter := chain.Ctrl.LogFilter
	globalErrLogFilter := chain.Ctrl.ErrLogFilter
	for i, stage := range stages {
		chainInfo := YomelInfo{
			No:           stage.No,
			Desc:         stage.Desc,
			IsLog:        chain.Ctrl.IsLog,
			LogFilter:    insertFilterShellStr(globalLogFilter, stage.LogFilter),
			ErrLogFilter: insertFilterShellStr(globalErrLogFilter, stage.ErrLogFilter),
		}
		var stageCmd stageCommand
		stageCmd.insertStageEl(
			[]string{stage.Cmd},
			backslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			stage.CmdOpArgs,
			backslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			[]string{stage.Svc},
			backslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			stage.SvcOpArgs,
			backslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			[]string{stage.Act},
			backslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			stage.ActOpArgs,
			backslashNewlineOpArgPrefix,
		)

		chainInfo.CmdStrs =
			strings.Trim(string(stageCmd), backslashNewlineOpArgPrefix)
		chainInfos[i] = chainInfo
	}
	return chainInfos
}

func (tChainStr *stageCommand) insertStageEl(insertStrs []string, prefix string) {
	if len(insertStrs) == 0 {
		return
	}
	*tChainStr += stageCommand(
		prefix +
			strings.Join(
				insertStrs,
				backslashNewlineOpArgPrefix,
			),
	)
}

func insertFilterShellStr(globalFilter string, logFilter string) string {
	if logFilter == "" {
		return globalFilter
	}
	return logFilter
}
