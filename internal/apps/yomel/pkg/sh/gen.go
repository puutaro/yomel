package sh

import (
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/domain"
)

const (
	opArgPrefixBlank            = " "
	backslashNewline            = "\\\n"
	backslashNewlineOpArgPrefix = opArgPrefixBlank + backslashNewline +
		opArgPrefixBlank
	verticalbar
)

type stageCommand string

type StageInfo struct {
	No           int
	Desc         string
	IsLog        bool
	LogFilter    string
	ErrLogFilter string
	CmdStrs      string
}
type YomelInfo struct {
	IsDirect   bool
	IsGen      bool
	StageInfos []StageInfo
}

func Gen(yomel domain.Yomel) YomelInfo {
	stageInfos := GenStageInfo(yomel)
	ctrl := yomel.Ctrl
	return YomelInfo{
		IsDirect:   ctrl.IsDirect,
		IsGen:      ctrl.IsGen,
		StageInfos: stageInfos,
	}
}

func GenStageInfo(yomel domain.Yomel) []StageInfo {
	stages := yomel.Stages
	yomelInfos := make([]StageInfo, len(stages))
	globalLogFilter := yomel.Ctrl.LogFilter
	globalErrLogFilter := yomel.Ctrl.ErrLogFilter
	// default stdout log don't display
	isLog := false
	if ctrlIsLog := yomel.Ctrl.IsLog; ctrlIsLog != nil {
		isLog = *ctrlIsLog
	}
	for i, stage := range stages {
		if stageIsLog := stage.IsLog; stageIsLog != nil {
			isLog = *stageIsLog
		}
		yomelInfo := StageInfo{
			No:           stage.No,
			Desc:         stage.Desc,
			IsLog:        isLog,
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
