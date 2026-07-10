package sh

import (
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/parser"
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

func Gen(yomel parser.Yomel) []YomelInfo {
	stages := yomel.Stages
	yomelInfos := make([]YomelInfo, len(stages))
	globalLogFilter := yomel.Ctrl.LogFilter
	globalErrLogFilter := yomel.Ctrl.ErrLogFilter
	for i, stage := range stages {
		yomelInfo := YomelInfo{
			No:           stage.No,
			Desc:         stage.Desc,
			IsLog:        yomel.Ctrl.IsLog,
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

		yomelInfo.CmdStrs =
			strings.Trim(string(stageCmd), backslashNewlineOpArgPrefix)
		yomelInfos[i] = yomelInfo
	}
	return yomelInfos
}

func (tYomelStr *stageCommand) insertStageEl(insertStrs []string, prefix string) {
	if len(insertStrs) == 0 {
		return
	}
	*tYomelStr += stageCommand(
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
