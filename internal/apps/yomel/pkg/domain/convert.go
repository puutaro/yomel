package domain

import (
	"fmt"
	"slices"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
)

type opArgType struct {
	Index int
	Str   string
}
type Control struct {
	IsLog        *bool
	LogFilter    string
	ErrLogFilter string
	IsVersion    bool
	IsHelp       bool
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
	IsLog        *bool
	LogFilter    string
	ErrLogFilter string
}
type Yomel struct {
	Ctrl   Control
	Stages []Stage
}

func Convert(ctrlModel model.ControlModel, stModels []model.StageModel) Yomel {
	ctrl := Control{
		IsLog:        ctrlModel.IsLog,
		LogFilter:    ctrlModel.LogFilter,
		ErrLogFilter: ctrlModel.ErrLogFilter,
		IsVersion:    ctrlModel.IsVersion,
		IsHelp:       ctrlModel.IsHelp,
	}
	yomel := Yomel{Ctrl: ctrl}
	stages := make([]Stage, len(stModels))
	for i, stModel := range stModels {
		var stage = Stage{
			No:           stModel.No,
			Desc:         stModel.Desc,
			Cmd:          stModel.Cmd,
			Svc:          stModel.Svc,
			Act:          stModel.Act,
			IsLog:        stModel.IsLog,
			LogFilter:    stModel.LogFilter,
			ErrLogFilter: stModel.ErrLogFilter,
		}
		pushOpArgs(
			stModel.CmdOps,
			stModel.CmdLops,
			stModel.CmdArgs,
			func(opArgList []string) {
				stage.CmdOpArgs = opArgList
			},
		)
		pushOpArgs(
			stModel.SvcOps,
			stModel.SvcLops,
			stModel.SvcArgs,
			func(opArgList []string) {
				stage.SvcOpArgs = opArgList
			},
		)
		pushOpArgs(
			stModel.ActOps,
			stModel.ActLops,
			stModel.ActArgs,
			func(opArgList []string) {
				stage.ActOpArgs = opArgList
			},
		)
		stages[i] = stage
	}
	yomel.Stages = stages
	return yomel
}

func pushOpArgs(
	ops []model.OptParam,
	lOps []model.OptParam,
	args []model.ArgParam,
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
		return a.Index - b.Index
	})
	for _, cmdLOpArgType := range opArgTypeList {
		opArgStrs = append(opArgStrs, cmdLOpArgType.Str)
	}
	insertFn(opArgStrs)
}

func makeOptList(
	optPs []model.OptParam,
	opPrefix string,
) []opArgType {
	var cmdLOpTypes []opArgType
	for _, op := range optPs {
		optStr := op.OptStr
		p := op.Param
		strP := p.Str
		oat := opArgType{
			Index: op.Index,
		}
		if strP == nil {
			oat.Str = fmt.Sprintf(`%s%s`, opPrefix, optStr)
			cmdLOpTypes = append(cmdLOpTypes, oat)
			continue
		}
		str := *strP
		switch p.QuoteType {
		case argtables.DoubleQuote:
			oat.Str = fmt.Sprintf(`%s%s "%s"`, opPrefix, optStr, str)
		case argtables.SingleQuote:
			oat.Str = fmt.Sprintf(`%s%s '%s'`, opPrefix, optStr, str)
		case argtables.NoQuote:
			oat.Str = fmt.Sprintf(`%s%s %s`, opPrefix, optStr, str)
		}
		cmdLOpTypes = append(cmdLOpTypes, oat)
	}
	return cmdLOpTypes
}

func makeArgList(
	argPs []model.ArgParam,
) []opArgType {
	var cmdArgTypes []opArgType
	for _, arg := range argPs {
		p := arg.Param
		strP := p.Str
		oat := opArgType{
			Index: arg.Index,
		}
		if strP == nil {
			oat.Str = ""
			cmdArgTypes = append(cmdArgTypes, oat)
			continue
		}
		str := *strP
		switch p.QuoteType {
		case argtables.DoubleQuote:
			oat.Str = fmt.Sprintf(`"%s"`, str)
		case argtables.SingleQuote:
			oat.Str = fmt.Sprintf(`'%s'`, str)
		case argtables.NoQuote:
			oat.Str = fmt.Sprintf(`%s`, str)
		}
		cmdArgTypes = append(cmdArgTypes, oat)
	}
	return cmdArgTypes
}
