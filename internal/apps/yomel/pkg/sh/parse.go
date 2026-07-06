package sh

import (
	"fmt"
	"slices"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
)

type paramType struct {
	str       *string
	quoteType args.QuoteType
}

type optParam struct {
	index  int
	optStr string
	param  paramType
}
type argParam struct {
	index int
	param paramType
}
type stageModel struct {
	no           int
	desc         string
	log          string
	cmd          string
	cmdOps       []optParam
	cmdLops      []optParam
	cmdArgs      []argParam
	svc          string
	svcOps       []optParam
	svcLops      []optParam
	svcArgs      []argParam
	act          string
	actOps       []optParam
	actLops      []optParam
	actArgs      []argParam
	logFilter    string
	errLogFilter string
	// Arg    []Param
	// Opt    []Param
}

type opArgType struct {
	index int
	str   string
}
type Stage struct {
	No           int
	Desc         string
	Cmd          string
	CmdOpArgs    []string
	Svc          string
	SvcOpArgs    []string
	Act          string
	ActOpArgs    []string
	LogFilter    string
	ErrLogFilter string
}
type Control struct {
	IsLog        bool
	LogFilter    string
	ErrLogFilter string
}

type Yomel struct {
	Ctrl   Control
	Stages []Stage
}

func Parse(argTables []args.ArgTable) Yomel {
	ctrl, stModels := ParseStageModels(argTables)
	return ParseStage(ctrl, stModels)
}

func ParseStage(ctrl Control, stModels []stageModel) Yomel {
	yomel := Yomel{Ctrl: ctrl}
	stages := make([]Stage, len(stModels))
	for i, stModel := range stModels {
		var stage = Stage{
			No:           stModel.no,
			Desc:         stModel.desc,
			Cmd:          stModel.cmd,
			Svc:          stModel.svc,
			Act:          stModel.act,
			LogFilter:    stModel.logFilter,
			ErrLogFilter: stModel.errLogFilter,
		}
		pushOpArgs(
			stModel.cmdOps,
			stModel.cmdLops,
			stModel.cmdArgs,
			func(opArgList []string) {
				stage.CmdOpArgs = opArgList
			},
		)
		pushOpArgs(
			stModel.svcOps,
			stModel.svcLops,
			stModel.svcArgs,
			func(opArgList []string) {
				stage.SvcOpArgs = opArgList
			},
		)
		pushOpArgs(
			stModel.actOps,
			stModel.actLops,
			stModel.actArgs,
			func(opArgList []string) {
				stage.ActOpArgs = opArgList
			},
		)
		stages[i] = stage
	}
	yomel.Stages = stages
	return yomel
}

func ParseStageModels(argTables []args.ArgTable) (Control, []stageModel) {
	var curCtrlArgTables []args.ArgTable
	for _, argTable := range argTables {
		if argTable.StageNo > 0 {
			break
		}
		curCtrlArgTables = append(
			curCtrlArgTables,
			argTable,
		)
	}
	ctrl := Control{}
	ctrl.IsLog = getFlag(
		0,
		curCtrlArgTables,
		func(t args.ArgTable) bool { return t.IsLog },
		false,
	)
	if strPtr := getOneStr(
		0,
		curCtrlArgTables,
		func(t args.ArgTable) bool { return t.IsLogFilter },
	); strPtr != nil {
		ctrl.LogFilter = *strPtr
	}
	if strPtr := getOneStr(
		0,
		curCtrlArgTables,
		func(t args.ArgTable) bool { return t.IsErrLogFilter },
	); strPtr != nil {
		ctrl.ErrLogFilter = *strPtr
	}

	argTablesLen := len(argTables)
	totalStageLen := argTables[argTablesLen-1].StageNo
	stModels := make([]stageModel, totalStageLen)
	for stageNo := 1; stageNo <= totalStageLen; stageNo++ {
		nextStartIndex := 0
		stModel := stageModel{
			no: stageNo,
		}
		var curStageArgTables []args.ArgTable
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
		stModel.desc = *getOneStr(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsStage },
		)
		stModel.cmd = *getOneStr(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsCmd },
		)
		if strPtr := getOneStr(
			0,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsLogFilter },
		); strPtr != nil {
			stModel.logFilter = *strPtr
		}
		if strPtr := getOneStr(
			0,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsErrLogFilter },
		); strPtr != nil {
			stModel.errLogFilter = *strPtr
		}
		parseCmdOptions(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsCmd },
			func(t args.ArgTable) bool { return t.IsSvc || t.IsAct },
			func(t args.ArgTable) bool { return t.IsOpt },
			func(p optParam) { stModel.cmdOps = append(stModel.cmdOps, p) },
		)
		parseCmdOptions(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsCmd },
			func(t args.ArgTable) bool { return t.IsSvc || t.IsAct },
			func(t args.ArgTable) bool { return t.IsLopt },
			func(p optParam) { stModel.cmdLops = append(stModel.cmdLops, p) },
		)
		nextStartIndex = parseArg(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsSvc || t.IsAct },
			func(t args.ArgTable) bool { return t.IsCmd },
			func(ind int, p paramType) {
				stModel.cmdArgs = append(
					stModel.cmdArgs,
					argParam{
						index: ind,
						param: p,
					},
				)
			},
		)
		if oneStr := getOneStr(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsSvc },
		); oneStr != nil {
			stModel.svc = *oneStr
		}
		parseCmdOptions(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsSvc },
			func(t args.ArgTable) bool { return t.IsAct },
			func(t args.ArgTable) bool { return t.IsOpt },
			func(p optParam) { stModel.svcOps = append(stModel.svcOps, p) },
		)
		parseCmdOptions(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsSvc },
			func(t args.ArgTable) bool { return t.IsAct },
			func(t args.ArgTable) bool { return t.IsLopt },
			func(p optParam) { stModel.svcLops = append(stModel.svcLops, p) },
		)
		nextStartIndex = parseArg(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsAct },
			func(t args.ArgTable) bool { return t.IsSvc },
			func(ind int, p paramType) {
				stModel.svcArgs = append(
					stModel.svcArgs,
					argParam{
						index: ind,
						param: p,
					},
				)
			},
		)
		if oneStr := getOneStr(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsAct },
		); oneStr != nil {
			stModel.act = *oneStr
		}
		parseCmdOptions(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsAct },
			func(t args.ArgTable) bool { return t.IsArg },
			func(t args.ArgTable) bool { return t.IsOpt },
			func(p optParam) { stModel.actOps = append(stModel.actOps, p) },
		)
		parseCmdOptions(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return t.IsAct },
			func(t args.ArgTable) bool { return t.IsArg },
			func(t args.ArgTable) bool { return t.IsLopt },
			func(p optParam) { stModel.actLops = append(stModel.actLops, p) },
		)
		nextStartIndex = parseArg(
			nextStartIndex,
			curStageArgTables,
			func(t args.ArgTable) bool { return false },
			func(t args.ArgTable) bool { return t.IsAct },
			func(ind int, p paramType) {
				stModel.actArgs = append(
					stModel.actArgs,
					argParam{
						index: ind,
						param: p,
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
	curStageArgTables []args.ArgTable,
	isCheckFn func(args.ArgTable) bool,
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

func getOneStr(
	nextStartIndex int,
	curStageArgTables []args.ArgTable,
	isCheckFn func(args.ArgTable) bool,
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

func parseCmdOptions(
	nextStartIndex int,
	curStageArgTables []args.ArgTable,
	isTargetMainArg func(t args.ArgTable) bool,
	isNextMainArg func(t args.ArgTable) bool,
	isTargetOpt func(args.ArgTable) bool,
	appendFn func(optParam),
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
		optParam := optParam{
			index:  optStrIndex,
			optStr: *curStageArgTables[optStrIndex].Str,
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
		optParam.param = param

		appendFn(optParam)
	}
}

func parseArg(
	nextStartIndex int,
	curStageArgTables []args.ArgTable,
	isNextMainArg func(t args.ArgTable) bool,
	isTargetMainArg func(t args.ArgTable) bool,
	appendFn func(int, paramType),
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
			j = updateIndex
			returnNextStartIndex = updateIndex
			appendFn(updateIndex, param)
		}
	}
	return returnNextStartIndex
}

func getQuoteStr(curStageArgTables []args.ArgTable, curIndex int) (paramType, int) {
	param := paramType{}
	afterFirstIndex := curIndex + 1
	if curStageArgTables[afterFirstIndex].QuoteTypeSignal == args.DoubleQuote {
		param.str = curStageArgTables[afterFirstIndex].Str
		return param, afterFirstIndex
	}
	afterNextIndex := curIndex + 2
	param.quoteType = curStageArgTables[afterFirstIndex].QuoteTypeSignal
	param.str = curStageArgTables[afterNextIndex].Str
	return param, afterNextIndex
}

func pushOpArgs(
	ops []optParam,
	lOps []optParam,
	args []argParam,
	insertFn func([]string),
) {
	shortOpPrefix := "-"
	longOpPrefix := "--"
	var opArgStrs []string
	opTypes := makeOptList(ops, shortOpPrefix)
	lOpTypes := makeOptList(lOps, longOpPrefix)
	argTypes := makeArgList(args)
	totalArgOpLen := len(opTypes) +
		len(lOpTypes) +
		len(argTypes)
	opArgTypeList := make([]opArgType, 0, totalArgOpLen)

	opArgTypeList = append(opArgTypeList, opTypes...)
	opArgTypeList = append(opArgTypeList, lOpTypes...)
	opArgTypeList = append(opArgTypeList, argTypes...)
	slices.SortFunc(opArgTypeList, func(a, b opArgType) int {
		return a.index - b.index
	})
	for _, cmdLOpArgType := range opArgTypeList {
		opArgStrs = append(opArgStrs, cmdLOpArgType.str)
	}
	insertFn(opArgStrs)
}

func makeOptList(
	optPs []optParam,
	opPrefix string,
) []opArgType {
	var cmdLOpTypes []opArgType
	for _, op := range optPs {
		optStr := op.optStr
		p := op.param
		strP := p.str
		oat := opArgType{
			index: op.index,
		}
		if strP == nil {
			oat.str = fmt.Sprintf(`%s%s`, opPrefix, optStr)
			cmdLOpTypes = append(cmdLOpTypes, oat)
			continue
		}
		str := *strP
		switch p.quoteType {
		case args.DoubleQuote:
			oat.str = fmt.Sprintf(`%s%s "%s"`, opPrefix, optStr, str)
		case args.SingleQuote:
			oat.str = fmt.Sprintf(`%s%s '%s'`, opPrefix, optStr, str)
		case args.NoQuote:
			oat.str = fmt.Sprintf(`%s%s %s`, opPrefix, optStr, str)
		}
		cmdLOpTypes = append(cmdLOpTypes, oat)
	}
	return cmdLOpTypes
}

func makeArgList(
	argPs []argParam,
) []opArgType {
	var cmdArgTypes []opArgType
	for _, arg := range argPs {
		p := arg.param
		strP := p.str
		oat := opArgType{
			index: arg.index,
		}
		if strP == nil {
			oat.str = ""
			cmdArgTypes = append(cmdArgTypes, oat)
			continue
		}
		str := *strP
		switch p.quoteType {
		case args.DoubleQuote:
			oat.str = fmt.Sprintf(`"%s"`, str)
		case args.SingleQuote:
			oat.str = fmt.Sprintf(`'%s'`, str)
		case args.NoQuote:
			oat.str = fmt.Sprintf(`%s`, str)
		}
		cmdArgTypes = append(cmdArgTypes, oat)
	}
	return cmdArgTypes
}
