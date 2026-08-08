package domain

import (
	"fmt"
	"slices"
	"strings"
	"unicode"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
)

const (
	SideCommentBlank            = "SIED_COmMetN_BAlNk"
	Commentout                  = "#"
	OpArgCommentPrefixBlank     = " "
	SpaceCommentout             = OpArgCommentPrefixBlank + Commentout
	BackslashNewline            = "\\\n"
	BackslashNewlineOpArgPrefix = OpArgCommentPrefixBlank + BackslashNewline +
		OpArgCommentPrefixBlank
)

type opArgType struct {
	Index int
	Str   string
}
type Control struct {
	IsGen        bool
	IsLog        *bool
	LogFilter    string
	ErrLogFilter string
	Title        string
	IsVersion    bool
	IsHelp       bool
	IsDirect     bool
	IsLiveStdout bool
	IsLiveStderr bool
}
type Stage struct {
	No                   int
	Desc                 string
	Cmd                  string
	CmdOpArgs            []string
	CmdOpArgsWithComment []string
	Svc                  string
	SvcOpArgs            []string
	SvcOpArgsWithComment []string
	Act                  string
	ActOpArgs            []string
	ActOpArgsWithComment []string
	IsLog                *bool
	LogFilter            string
	ErrLogFilter         string
}
type Yomel struct {
	Ctrl   Control
	Stages []Stage
}

func Convert(ctrlModel model.ControlModel, stModels []model.StageModel) Yomel {
	stringValue := func(s *string) string {
		if s == nil || *s == "" {
			return ""
		}
		return *s
	}
	ctrl := Control{
		IsGen:        ctrlModel.IsGen,
		IsLog:        ctrlModel.IsLog,
		Title:        ctrlModel.Title,
		LogFilter:    ctrlModel.LogFilter,
		ErrLogFilter: ctrlModel.ErrLogFilter,
		IsVersion:    ctrlModel.IsVersion,
		IsHelp:       ctrlModel.IsHelp,
		IsDirect:     ctrlModel.IsDirect,
		IsLiveStdout: ctrlModel.IsLiveStdout,
		IsLiveStderr: ctrlModel.IsLiveStderr,
	}
	yomel := Yomel{Ctrl: ctrl}
	stages := make([]Stage, len(stModels))
	for i, stModel := range stModels {
		var stage = Stage{
			No:   stModel.No,
			Desc: stModel.Desc,
			Cmd:  stModel.Cmd,
			// svc and act *string is no need, because finish validation and keep culc speed
			Svc:          stringValue(stModel.Svc),
			Act:          stringValue(stModel.Act),
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
		pushOpArgsWithComment(
			stModel.CmdOps,
			stModel.CmdLops,
			stModel.CmdArgs,
			func(opArgList []string) {
				stage.CmdOpArgsWithComment = opArgList
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
		pushOpArgsWithComment(
			stModel.SvcOps,
			stModel.SvcLops,
			stModel.SvcArgs,
			func(opArgList []string) {
				stage.SvcOpArgsWithComment = opArgList
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
		pushOpArgsWithComment(
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
func pushOpArgsWithComment(
	ops []model.OptParam,
	lOps []model.OptParam,
	args []model.ArgParam,
	insertFn func([]string),
) {
	shortOpPrefix := "-"
	longOpPrefix := "--"
	var opArgStrs []string
	opTypes := makeOptListWithComment(ops, shortOpPrefix)
	lOpTypes := makeOptListWithComment(lOps, longOpPrefix)
	argTypes := makeArgListWithComment(args)
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
func makeOptListWithComment(
	optPs []model.OptParam,
	opPrefix string,
) []opArgType {
	var cmdLOpTypes []opArgType
	for _, op := range optPs {
		optStr := op.OptStr
		p := op.Param
		oat := opArgType{
			Index: op.Index,
		}
		strP := p.Str
		escapeOpComment := makeEscapeComment(op.Comment)
		if strP == nil {
			oat.Str = fmt.Sprintf(`%s%s`, opPrefix, optStr) +
				SideCommentBlank + escapeOpComment
			cmdLOpTypes = append(cmdLOpTypes, oat)
			continue
		}
		oatStr := ""
		if escapeOpComment != "" {
			oatStr = fmt.Sprintf(
				"%s%s",
				escapeOpComment,
				BackslashNewlineOpArgPrefix,
			)
		}
		str := *strP
		switch p.QuoteType {
		case argtables.DoubleQuote:
			oatStr += fmt.Sprintf(`%s%s "%s"`, opPrefix, optStr, str)
		case argtables.SingleQuote:
			oatStr += fmt.Sprintf(`%s%s '%s'`, opPrefix, optStr, str)
		case argtables.NoQuote:
			oatStr += fmt.Sprintf(`%s%s %s`, opPrefix, optStr, str)
		}
		escapeValComemnt := makeEscapeComment(p.Comment)
		if escapeValComemnt == "" {
			oat.Str = oatStr
		} else {
			oat.Str = oatStr +
				SideCommentBlank + escapeValComemnt
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
func makeArgListWithComment(
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
		oatStr := ""
		switch p.QuoteType {
		case argtables.DoubleQuote:
			oatStr = fmt.Sprintf(`"%s"`, str)
		case argtables.SingleQuote:
			oatStr = fmt.Sprintf(`'%s'`, str)
		case argtables.NoQuote:
			oatStr = fmt.Sprintf(`%s`, str)
		}
		escapeArgComemnt := makeEscapeComment(p.Comment)
		if escapeArgComemnt == "" {
			oat.Str = oatStr
		} else {
			oat.Str = oatStr +
				SideCommentBlank + escapeArgComemnt
		}
		cmdArgTypes = append(cmdArgTypes, oat)
	}
	return cmdArgTypes
}

func makeEscapeComment(comment string) string {
	if comment == "" {
		return ""
	}
	grayStart := "\x1b[38;5;244m"
	colorEnd := "\x1b[39m"
	return fmt.Sprintf(
		"%s`%s%s`%s",
		grayStart,
		Commentout,
		toLowerWithSpaces(comment),
		colorEnd,
	)
}

func toLowerWithSpaces(s string) string {
	var sb strings.Builder
	runes := []rune(s)
	runesLen := len(runes)
	for i := 0; i < runesLen; i++ {
		r := runes[i]
		nextRuneIndex := i + 1
		if !unicode.IsUpper(r) ||
			(nextRuneIndex < runesLen &&
				unicode.IsUpper(runes[nextRuneIndex])) {
			sb.WriteRune(r)
			continue
		}
		if i > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteRune(unicode.ToLower(r))
	}
	return sb.String()
}
