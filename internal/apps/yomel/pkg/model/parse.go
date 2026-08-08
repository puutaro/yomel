package model

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

type ParamType struct {
	Str       *string
	QuoteType argtables.QuoteType
	Comment   string
}

type OptParam struct {
	Index   int
	OptStr  string
	Comment string
	Param   ParamType
}
type ArgParam struct {
	Index int
	Param ParamType
}
type StageModel struct {
	No           int
	Desc         string
	Cmd          string
	CmdOps       []OptParam
	CmdLops      []OptParam
	CmdArgs      []ArgParam
	Svc          *string
	SvcOps       []OptParam
	SvcLops      []OptParam
	SvcArgs      []ArgParam
	Act          *string
	ActOps       []OptParam
	ActLops      []OptParam
	ActArgs      []ArgParam
	IsLog        *bool
	LogFilter    string
	ErrLogFilter string
	// Arg    []Param
	// Opt    []Param
}

type ControlModel struct {
	IsLog        *bool
	LogFilter    string
	ErrLogFilter string
	IsVersion    bool
	IsHelp       bool
	IsGen        bool
	IsDirect     bool
	IsLiveStdout bool
	IsLiveStderr bool
	Title        string
}

func Parse(argTables []argtables.ArgTable) (ControlModel, []StageModel) {
	var curCtrlArgTables []argtables.ArgTable
	for _, argTable := range argTables {
		if argTable.StageNo > 0 {
			break
		}
		curCtrlArgTables = append(
			curCtrlArgTables,
			argTable,
		)
	}
	ctrl := ControlModel{}
	ctrl.IsVersion = getFlag(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsVersion },
		false,
	)
	ctrl.IsHelp = getFlag(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsHelp },
		false,
	)
	ctrl.IsGen = getFlag(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsGen },
		false,
	)
	ctrl.IsDirect = getFlag(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsDirect },
		false,
	)
	ctrl.IsLiveStdout = !getFlag(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool {
			return t.IsNoLiveStdout
		},
		false,
	)
	ctrl.IsLiveStderr = !getFlag(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool {
			return t.IsNoLiveStderr
		},
		false,
	)
	if strPtr := getOneStr(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsTitle },
	); strPtr != nil {
		ctrl.Title = *strPtr
	}
	if flagBool := getFlagByPtr(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsLog },
		true,
	); flagBool != nil {
		ctrl.IsLog = flagBool
	}
	if flagBool := getFlagByPtr(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsNoLog },
		false,
	); flagBool != nil {
		ctrl.IsLog = flagBool
	}
	if strPtr := getOneStr(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsLogFilter },
	); strPtr != nil {
		ctrl.LogFilter = *strPtr
	}
	if strPtr := getOneStr(
		0,
		curCtrlArgTables,
		func(t argtables.ArgTable) bool { return t.IsErrLogFilter },
	); strPtr != nil {
		ctrl.ErrLogFilter = *strPtr
	}

	argTablesLen := len(argTables)
	totalStageLen := argTables[argTablesLen-1].StageNo
	stModels := make([]StageModel, totalStageLen)
	for stageNo := 1; stageNo <= totalStageLen; stageNo++ {
		nextStartIndex := 0
		stModel := StageModel{
			No: stageNo,
		}
		var curStageArgTables []argtables.ArgTable
		for i := nextStartIndex; i < argTablesLen; i++ {
			argTable := argTables[i]
			if argTable.StageNo < stageNo {
				continue
			}
			if argTable.StageNo > stageNo {
				break
			}
			curStageArgTables = append(
				curStageArgTables,
				argTable,
			)
		}
		if ptrStr := getOneStr(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsStage },
		); ptrStr != nil {
			stModel.Desc = *ptrStr
		}
		if ptrStr := getOneStr(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsCmd },
		); ptrStr != nil {
			stModel.Cmd = *ptrStr
		}
		if flagBool := getFlagByPtr(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsLog },
			true,
		); flagBool != nil {
			stModel.IsLog = flagBool
		}
		if flagBool := getFlagByPtr(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsNoLog },
			false,
		); flagBool != nil {
			stModel.IsLog = flagBool
		}
		if strPtr := getOneStr(
			0,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsLogFilter },
		); strPtr != nil {
			stModel.LogFilter = *strPtr
		}
		if strPtr := getOneStr(
			0,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsErrLogFilter },
		); strPtr != nil {
			stModel.ErrLogFilter = *strPtr
		}
		parseOptions(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsCmd },
			func(t argtables.ArgTable) bool { return t.IsSvc || t.IsAct },
			func(t argtables.ArgTable) bool { return t.IsOpt },
			func(p OptParam) { stModel.CmdOps = append(stModel.CmdOps, p) },
		)
		parseOptions(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsCmd },
			func(t argtables.ArgTable) bool { return t.IsSvc || t.IsAct },
			func(t argtables.ArgTable) bool { return t.IsLopt },
			func(p OptParam) { stModel.CmdLops = append(stModel.CmdLops, p) },
		)
		nextStartIndex = parseArg(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsSvc || t.IsAct },
			func(t argtables.ArgTable) bool { return t.IsCmd },
			func(ind int, p ParamType) {
				stModel.CmdArgs = append(
					stModel.CmdArgs,
					ArgParam{
						Index: ind,
						Param: p,
					},
				)
			},
		)
		stModel.Svc = getOneStr(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsSvc },
		)
		parseOptions(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsSvc },
			func(t argtables.ArgTable) bool { return t.IsAct },
			func(t argtables.ArgTable) bool { return t.IsOpt },
			func(p OptParam) { stModel.SvcOps = append(stModel.SvcOps, p) },
		)
		parseOptions(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsSvc },
			func(t argtables.ArgTable) bool { return t.IsAct },
			func(t argtables.ArgTable) bool { return t.IsLopt },
			func(p OptParam) { stModel.SvcLops = append(stModel.SvcLops, p) },
		)
		nextStartIndex = parseArg(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsAct },
			func(t argtables.ArgTable) bool { return t.IsSvc },
			func(ind int, p ParamType) {
				stModel.SvcArgs = append(
					stModel.SvcArgs,
					ArgParam{
						Index: ind,
						Param: p,
					},
				)
			},
		)
		stModel.Act = getOneStr(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsAct },
		)
		parseOptions(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsAct },
			func(t argtables.ArgTable) bool { return t.IsArg },
			func(t argtables.ArgTable) bool { return t.IsOpt },
			func(p OptParam) { stModel.ActOps = append(stModel.ActOps, p) },
		)
		parseOptions(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return t.IsAct },
			func(t argtables.ArgTable) bool { return t.IsArg },
			func(t argtables.ArgTable) bool { return t.IsLopt },
			func(p OptParam) { stModel.ActLops = append(stModel.ActLops, p) },
		)
		nextStartIndex = parseArg(
			nextStartIndex,
			curStageArgTables,
			func(t argtables.ArgTable) bool { return false },
			func(t argtables.ArgTable) bool { return t.IsAct },
			func(ind int, p ParamType) {
				stModel.ActArgs = append(
					stModel.ActArgs,
					ArgParam{
						Index: ind,
						Param: p,
					},
				)
			},
		)
		stModels[stageNo-1] = stModel
	}
	return ctrl, stModels
}

func getFlag(
	nextStartIndex int,
	curStageArgTables []argtables.ArgTable,
	isCheckFn func(argtables.ArgTable) bool,
	defaultBool bool,
) bool {
	argTablesLen := len(curStageArgTables)
	for i := nextStartIndex; i < argTablesLen; i++ {
		argTable := curStageArgTables[i]
		if !isCheckFn(argTable) {
			continue
		}
		return !defaultBool
	}
	return defaultBool
}
func getFlagByPtr(
	nextStartIndex int,
	curStageArgTables []argtables.ArgTable,
	isCheckFn func(argtables.ArgTable) bool,
	returnBool bool,
) *bool {
	argTablesLen := len(curStageArgTables)
	for i := nextStartIndex; i < argTablesLen; i++ {
		argTable := curStageArgTables[i]
		if !isCheckFn(argTable) {
			continue
		}
		return &returnBool
	}
	return nil
}

func getOneStr(
	nextStartIndex int,
	curStageArgTables []argtables.ArgTable,
	isCheckFn func(argtables.ArgTable) bool,
) *string {
	argTablesLen := len(curStageArgTables)
	for i := nextStartIndex; i < argTablesLen; i++ {
		argTable := curStageArgTables[i]
		if !isCheckFn(argTable) {
			continue
		}
		if i+1 >= argTablesLen {
			continue
		}
		return curStageArgTables[i+1].Str
	}
	return nil
}

func parseOptions(
	nextStartIndex int,
	curStageArgTables []argtables.ArgTable,
	isTargetMainArg func(t argtables.ArgTable) bool,
	isNextMainArg func(t argtables.ArgTable) bool,
	isTargetOpt func(argtables.ArgTable) bool,
	appendFn func(OptParam),
) {
	curStageArgTablesLen := len(curStageArgTables)
	curStageArgTablesLastIndex := curStageArgTablesLen - 1
	seekStartIndex := nextStartIndex
	for ; seekStartIndex < curStageArgTablesLen; seekStartIndex++ {
		argTable := curStageArgTables[seekStartIndex]
		if seekStartIndex == curStageArgTablesLastIndex {
			return
		}
		if isNextMainArg(argTable) {
			return
		}
		if isTargetMainArg(argTable) {
			break
		}
	}
	for j := seekStartIndex + 1; j < curStageArgTablesLen; j++ {
		innerArgTable := curStageArgTables[j]
		if isNextMainArg(innerArgTable) {
			return
		}
		if !isTargetOpt(innerArgTable) {
			continue
		}
		optStrIndex := j + 1
		optStrIndexStageArgTable := curStageArgTables[optStrIndex]
		optParam := OptParam{
			Index:   optStrIndex,
			OptStr:  *optStrIndexStageArgTable.Str,
			Comment: innerArgTable.Comment,
		}
		valueOpIndex := j + 2
		if valueOpIndex >= curStageArgTablesLen ||
			!curStageArgTables[valueOpIndex].IsValue {
			appendFn(optParam)
			continue
		}
		param, updateIndex :=
			getQuoteStr(curStageArgTables, valueOpIndex)
		j = updateIndex
		param.Comment = curStageArgTables[valueOpIndex].Comment
		optParam.Param = param

		appendFn(optParam)
	}
}

func parseArg(
	nextStartIndex int,
	curStageArgTables []argtables.ArgTable,
	isNextMainArg func(t argtables.ArgTable) bool,
	isTargetMainArg func(t argtables.ArgTable) bool,
	appendFn func(int, ParamType),
) int {
	curStageArgTablesLen := len(curStageArgTables)
	returnNextStartIndex := nextStartIndex
	for i := nextStartIndex; i < curStageArgTablesLen; i++ {
		argTable := curStageArgTables[i]
		if isNextMainArg(argTable) {
			break
		}
		if !isTargetMainArg(argTable) {
			continue
		}
		for j := i + 1; j < curStageArgTablesLen; j++ {
			innerArgTable := curStageArgTables[j]

			if isNextMainArg(innerArgTable) {
				break
			}
			if !innerArgTable.IsArg {
				continue
			}
			param, updateIndex := getQuoteStr(curStageArgTables, j)
			param.Comment = innerArgTable.Comment
			j = updateIndex
			returnNextStartIndex = updateIndex
			appendFn(updateIndex, param)
		}
	}
	return returnNextStartIndex
}

func getQuoteStr(curStageArgTables []argtables.ArgTable, curIndex int) (ParamType, int) {
	param := ParamType{}
	afterFirstIndex := curIndex + 1
	afterFirstIndexStageArgTable := curStageArgTables[afterFirstIndex]
	if afterFirstIndexStageArgTable.QuoteTypeSignal == argtables.DoubleQuote {
		param.Str = afterFirstIndexStageArgTable.Str
		return param, afterFirstIndex
	}
	afterNextIndex := curIndex + 2
	afterNextIndexStageArgTable := curStageArgTables[afterNextIndex]
	param.QuoteType = afterFirstIndexStageArgTable.QuoteTypeSignal
	param.Str = afterNextIndexStageArgTable.Str
	return param, afterNextIndex
}
