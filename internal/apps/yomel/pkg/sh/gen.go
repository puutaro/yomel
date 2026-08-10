package sh

import (
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/domain"
)

type stageCommand string

type StageInfo struct {
	No                 int
	Desc               string
	IsLog              bool
	LogFilter          string
	ErrLogFilter       string
	CmdStrs            string
	CmdStrsWithComment string
}
type YomelInfo struct {
	IsDirect     bool
	IsGen        bool
	IsLiveStdout bool
	IsLiveStdErr bool
	Title        string
	StageInfos   []StageInfo
}

func Gen(yomel domain.Yomel) YomelInfo {
	stageInfos := genStageInfo(yomel)
	ctrl := yomel.Ctrl
	return YomelInfo{
		IsLiveStdout: ctrl.IsLiveStdout,
		IsLiveStdErr: ctrl.IsLiveStderr,
		IsDirect:     ctrl.IsDirect,
		IsGen:        ctrl.IsGen,
		Title:        ctrl.Title,
		StageInfos:   stageInfos,
	}
}

func genStageInfo(yomel domain.Yomel) []StageInfo {
	stages := yomel.Stages
	yomelInfos := make([]StageInfo, len(stages))
	globalLogFilter := yomel.Ctrl.LogFilter
	globalErrLogFilter := yomel.Ctrl.ErrLogFilter
	// default stdout log don't display
	isLogInCtrl := false
	if ctrlIsLog := yomel.Ctrl.IsLog; ctrlIsLog != nil {
		isLogInCtrl = *ctrlIsLog
	}
	for i, stage := range stages {
		isLogInStage := isLogInCtrl
		if stageIsLog := stage.IsLog; stageIsLog != nil {
			isLogInStage = *stageIsLog
		}
		yomelInfo := StageInfo{
			No:           stage.No,
			Desc:         stage.Desc,
			IsLog:        isLogInStage,
			LogFilter:    insertFilterShellStr(globalLogFilter, stage.LogFilter),
			ErrLogFilter: insertFilterShellStr(globalErrLogFilter, stage.ErrLogFilter),
		}
		var stageCmd stageCommand
		cmdStrList := makeStrListFromStr(stage.Cmd)
		svcStrList := makeStrListFromStr(stage.Svc)
		actStrList := makeStrListFromStr(stage.Act)
		stageCmd.insertStageEl(
			cmdStrList,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			stage.CmdOpArgs,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			svcStrList,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			stage.SvcOpArgs,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			actStrList,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmd.insertStageEl(
			stage.ActOpArgs,
			domain.BackslashNewlineOpArgPrefix,
		)

		yomelInfo.CmdStrs =
			strings.Trim(string(stageCmd), domain.BackslashNewlineOpArgPrefix)

		var stageCmdWithComment stageCommand
		stageCmdWithComment.insertStageEl(
			cmdStrList,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmdWithComment.insertStageEl(
			stage.CmdOpArgsWithComment,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmdWithComment.insertStageEl(
			svcStrList,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmdWithComment.insertStageEl(
			stage.SvcOpArgsWithComment,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmdWithComment.insertStageEl(
			actStrList,
			domain.BackslashNewlineOpArgPrefix,
		)
		stageCmdWithComment.insertStageEl(
			stage.ActOpArgsWithComment,
			domain.BackslashNewlineOpArgPrefix,
		)
		yomelInfo.CmdStrsWithComment =
			strings.Trim(string(stageCmdWithComment), domain.BackslashNewlineOpArgPrefix)

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
				domain.BackslashNewlineOpArgPrefix,
			),
	)
}

func insertFilterShellStr(globalFilter string, logFilter string) string {
	if logFilter == "" {
		return globalFilter
	}
	return logFilter
}

func makeStrListFromStr(str string) []string {
	if str == "" {
		return nil
	}
	return []string{str}
}
