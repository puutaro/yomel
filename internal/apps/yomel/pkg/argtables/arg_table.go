package argtables

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
)

const (
	Version            = "version"
	Help               = "help"
	DirectMode         = "direct"
	GenMode            = "gen"
	StageArgName       = "stage"
	LogOpName          = "log"
	NoLogOpName        = "no-log"
	LogFilter          = "log-filter"
	ErrLogFilter       = "err-log-filter"
	CmdOpName          = "cmd"
	SvcOpName          = "svc"
	ActOpName          = "act"
	OptOpName          = "opt"
	LoptOpName         = "lop"
	ArgOpName          = "arg"
	ValueOpName        = "val"
	SingleOpName       = "single"
	SingleShortOpName  = "s"
	NoQuoteOpName      = "no-quote"
	NoQuoteShortOpName = "n"
)
const (
	VersionOpSignal      = "--" + Version
	HelpOpSignal         = "--" + Help
	DirectModeFlagSignal = "--" + DirectMode
	GenModeFlagSingnal   = "--" + GenMode
	StageSignal          = StageArgName
	CmdOpSignal          = "-" + CmdOpName
	LogFlagSignal        = "--" + LogOpName
	NoLogFlagSignal      = "--" + NoLogOpName
	LogFilterOpSignal    = "--" + LogFilter
	ErrLogFilterOpSignal = "--" + ErrLogFilter
	SvcOpSignal          = "-" + SvcOpName
	ActOpSignal          = "-" + ActOpName
	OptOpSignal          = "--" + OptOpName
	LoptOpSignal         = "--" + LoptOpName
	ArgOpSignal          = "--" + ArgOpName
	ValueOptSignal       = "--" + ValueOpName
	SingleOpSignal       = "--" + SingleOpName
	SingleShortOpSignal  = "--" + SingleShortOpName
	NoQuoteOpSignal      = "--" + NoQuoteOpName
	NoQuoteShortOpSignal = "--" + NoQuoteShortOpName
)

type QuoteType int

const (
	DoubleQuote QuoteType = iota
	SingleQuote
	NoQuote
)

type ArgTable struct {
	No              int
	IsVersion       bool
	IsHelp          bool
	IsGen           bool
	IsDirect        bool
	IsLogFilter     bool
	IsErrLogFilter  bool
	StageNo         int
	IsStage         bool
	IsLog           bool
	IsNoLog         bool
	IsCmd           bool
	IsSvc           bool
	IsAct           bool
	IsOpt           bool
	IsLopt          bool
	IsValue         bool
	IsArg           bool
	QuoteTypeSignal QuoteType
	UnkownOption    string
	Str             *string
}
type ExpectedNext int

const (
	ExpectNormalStr ExpectedNext = iota
	ExpectUp2StageStr
	ExpectUp2HyphenStr
)

// TODO judge option or str about --s, --n
// I escape this judge for rare case considering --s, --n as str
// So this implement apply --s, --n as option always
func GenArgTable(argTableDtos []argtabledtos.ArgTableDto) []ArgTable {
	var argTables []ArgTable
	for _, argTableDto := range argTableDtos {
		argTable := ArgTable{
			No:              argTableDto.No,
			IsVersion:       argTableDto.IsVersion,
			IsHelp:          argTableDto.IsHelp,
			IsGen:           argTableDto.IsGen,
			IsDirect:        argTableDto.IsDirect,
			IsLogFilter:     argTableDto.IsLogFilter,
			IsErrLogFilter:  argTableDto.IsErrLogFilter,
			StageNo:         argTableDto.StageNo,
			IsStage:         argTableDto.IsStage,
			IsLog:           argTableDto.IsLog,
			IsNoLog:         argTableDto.IsNoLog,
			IsCmd:           argTableDto.IsCmd,
			IsSvc:           argTableDto.IsSvc,
			IsAct:           argTableDto.IsAct,
			IsOpt:           isSetStr(argTableDto.OptStr),
			IsLopt:          isSetStr(argTableDto.LoptStr),
			IsValue:         isSetStr(argTableDto.ValueStr),
			IsArg:           isSetStr(argTableDto.ArgStr),
			QuoteTypeSignal: QuoteType(argTableDto.QuoteTypeSignal),
			UnkownOption:    argTableDto.UnknownOption,
			Str:             argTableDto.Str,
		}

		argTables = append(argTables, argTable)
	}
	return argTables
}

func isSetStr(ptrStr *string) bool {
	return ptrStr != nil
}
