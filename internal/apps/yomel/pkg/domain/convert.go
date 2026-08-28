package domain

import (
	"fmt"
	"slices"
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/color"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
)

const (
	SideCommentBlank            = "SIED_COmMetN_BAlNk"
	Commentout                  = "#"
	OpArgCommentPrefixBlank     = " "
	SpaceCommentout             = Commentout + OpArgCommentPrefixBlank
	BackslashNewline            = "\\\n"
	BackslashNewlineOpArgPrefix = OpArgCommentPrefixBlank + BackslashNewline +
		OpArgCommentPrefixBlank
)
const (
	GrayStart               = "\x1b[38;5;244m"
	ColorEnd                = "\x1b[39m"
	ForegroundColorAnsiCode = 38
	BackgroundColorAnsiCode = 48
)

type ForeOrBack int

const (
	ForegroundAnsi ForeOrBack = iota
	BackgroundAnsi
)

type opArgType struct {
	Index int
	Str   string
}
type Control struct {
	IsGen                     bool
	IsLog                     *bool
	LogFilter                 string
	ErrLogFilter              string
	Title                     string
	IsVersion                 bool
	IsHelp                    bool
	IsDirect                  bool
	IsLiveStdout              bool
	IsLiveStderr              bool
	TitleColorStartStr        string
	TitleBgColorStartStr      string
	TitleCommentColorStartStr string
	CommentColorStartStr      string
}
type Stage struct {
	No                   int
	Desc                 string
	Cmd                  string
	CmdOpArgs            []string
	CmdOpArgsWithComment []string
	IsLog                *bool
	LogFilter            string
	ErrLogFilter         string
	ColorStartStr        string
	BgColorStartStr      string
	CommentColorStartStr string
}
type Yomel struct {
	Ctrl   Control
	Stages []Stage
}

func Convert(
	ctrlModel model.ControlModel,
	stModels []model.StageModel,
	yomelToml toml.LogConfig,
	isTerminal bool,
) Yomel {
	tomlColor := yomelToml.Color
	tomlStream := yomelToml.Stream
	ctrl := Control{
		IsGen: ctrlModel.IsGen,
		IsLog: ctrlModel.IsLog,
		Title: ctrlModel.Title,
		LogFilter: decideParameterOrToml(
			ctrlModel.LogFilter,
			tomlStream.LogFilterShell,
		),
		ErrLogFilter: decideParameterOrToml(
			ctrlModel.ErrLogFilter,
			tomlStream.ErrLogFilterShell,
		),
		IsVersion:    ctrlModel.IsVersion,
		IsHelp:       ctrlModel.IsHelp,
		IsDirect:     ctrlModel.IsDirect,
		IsLiveStdout: ctrlModel.IsLiveStdout,
		IsLiveStderr: ctrlModel.IsLiveStderr,
		TitleColorStartStr: hexToAnsiFg(
			decideParameterOrToml(
				ctrlModel.TitleColorStr,
				tomlColor.TitleColor,
			),
		),
		TitleBgColorStartStr: hexToAnsiBg(
			decideParameterOrToml(
				ctrlModel.TitleBgColorStr,
				tomlColor.TitleBgColor,
			),
		),
		TitleCommentColorStartStr: hexToAnsiFg(
			decideParameterOrToml(
				ctrlModel.TitleCommentColorStr,
				tomlColor.TitleCommentColor,
			),
		),
	}
	yomel := Yomel{Ctrl: ctrl}
	stages := make([]Stage, len(stModels))
	ctrlColorStartStrList := makeColorStrStartListByOrderOperator(
		decideParameterOrToml(
			ctrlModel.ColorStr,
			tomlColor.Color,
		),
	)
	ctrlColorStartStrListLen := len(ctrlColorStartStrList)
	ctrlBgColorStartStrList := makeColorStrStartListByOrderOperator(
		decideParameterOrToml(
			ctrlModel.BgColorStr,
			tomlColor.BgColor,
		),
	)
	ctrlBgColorStartStrListLen := len(ctrlBgColorStartStrList)
	ctrlCommentColorStartStrList := makeColorStrStartListByOrderOperator(
		decideParameterOrToml(
			ctrlModel.CommentColorStr,
			tomlColor.CommentColor,
		),
	)
	ctrlCommentColorStartStrListLen := len(ctrlCommentColorStartStrList)
	for i, stModel := range stModels {
		curCtrlColorIndex := i % ctrlColorStartStrListLen
		curCtrlBgColorIndex := i % ctrlBgColorStartStrListLen
		curCtrlCommentColorIndex := i % ctrlCommentColorStartStrListLen
		var stage = Stage{
			No:           stModel.No,
			Desc:         stModel.Desc,
			Cmd:          stModel.Cmd,
			IsLog:        stModel.IsLog,
			LogFilter:    stModel.LogFilter,
			ErrLogFilter: stModel.ErrLogFilter,
			ColorStartStr: hexToAnsiFgForStage(
				stModel.ColorStr,
				ctrlColorStartStrList[curCtrlColorIndex],
			),
			BgColorStartStr: hexToAnsiBgForStage(
				stModel.BgColorStr,
				ctrlBgColorStartStrList[curCtrlBgColorIndex],
			),
			CommentColorStartStr: hexToAnsiFgForStage(
				stModel.CommentColorStr,
				ctrlCommentColorStartStrList[curCtrlCommentColorIndex],
			),
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
			isTerminal,
		)
		stages[i] = stage
	}
	yomel.Stages = stages
	return yomel
}

func decideParameterOrToml(paraStr, tomlStr string) string {
	if paraStr == "" {
		return tomlStr
	}
	return paraStr
}

func makeColorStrStartListByOrderOperator(colorStr string) []string {
	return strings.Split(
		colorStr,
		color.OrderOperator,
	)
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
	isTerminal bool,
) {
	shortOpPrefix := "-"
	longOpPrefix := "--"
	var opArgStrs []string
	opTypes := makeOptListWithComment(
		ops,
		shortOpPrefix,
		isTerminal,
	)
	lOpTypes := makeOptListWithComment(
		lOps,
		longOpPrefix,
		isTerminal,
	)
	argTypes := makeArgListWithComment(
		args,
		isTerminal,
	)
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
	isTerminal bool,
) []opArgType {
	var cmdLOpTypes []opArgType
	for _, op := range optPs {
		optStr := op.OptStr
		p := op.Param
		oat := opArgType{
			Index: op.Index,
		}
		strP := p.Str
		escapeOpComment := makeEscapeComment(
			op.Comment,
			isTerminal,
		)
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
		escapeValComemnt := makeEscapeComment(
			p.Comment,
			isTerminal,
		)
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
	isTerminal bool,
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
		escapeArgComemnt := makeEscapeComment(
			p.Comment,
			isTerminal,
		)
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
