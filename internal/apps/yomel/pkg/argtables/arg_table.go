package argtables

import (
	"strings"
)

const (
	Version            = "version"
	VersionShort       = "v"
	Help               = "help"
	HelpShort          = "h"
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
	VersionShortOpSignal = "-" + VersionShort
	HelpOpSignal         = "--" + Help
	HelpShortOpSignal    = "-" + HelpShort
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
func GenArgTable(inputArgs []string) []ArgTable {
	var argTables []ArgTable
	stageNum := 0
	expectedNext := ExpectNormalStr

	for i, inputArg := range inputArgs {
		displayNum := i + 1
		argTable := ArgTable{
			No:      displayNum,
			StageNo: stageNum,
		}
		// Primary judge expectedNext about str or special option or stage
		switch true {
		case
			// Apply --s, --n as option always
			expectedNext == ExpectUp2HyphenStr &&
				(inputArg != NoQuoteOpSignal &&
					inputArg != NoQuoteShortOpSignal &&
					inputArg != SingleOpSignal &&
					inputArg != SingleShortOpSignal):
			val := inputArg
			argTable.Str = &val
			argTables = append(argTables, argTable)
			expectedNext = ExpectNormalStr
			continue
		case
			expectedNext == ExpectUp2StageStr &&
				!strings.HasPrefix(inputArg, "-"):
			val := inputArg
			argTable.Str = &val
			argTables = append(argTables, argTable)
			expectedNext = ExpectNormalStr
			continue
		case
			expectedNext == ExpectNormalStr &&
				!strings.HasPrefix(inputArg, "-") &&
				inputArg != StageSignal:
			val := inputArg
			argTable.Str = &val
			argTables = append(argTables, argTable)
			expectedNext = ExpectNormalStr
			continue
		}

		switch inputArg {
		case VersionOpSignal,
			VersionShortOpSignal:
			argTable.IsVersion = true
		case HelpOpSignal,
			HelpShortOpSignal:
			argTable.IsHelp = true
		case LogFlagSignal:
			argTable.IsLog = true
		case NoLogFlagSignal:
			argTable.IsNoLog = true
		case LogFilterOpSignal:
			argTable.IsLogFilter = true
		case ErrLogFilterOpSignal:
			argTable.IsErrLogFilter = true
		// stage interpret as special arg always in this line by first switch sentence
		case StageArgName:
			stageNum++
			argTable.StageNo = stageNum
			argTable.IsStage = true
		case CmdOpSignal:
			argTable.IsCmd = true
			expectedNext = ExpectUp2StageStr
		case SvcOpSignal:
			argTable.IsSvc = true
			expectedNext = ExpectUp2StageStr
		case ActOpSignal:
			argTable.IsAct = true
			expectedNext = ExpectUp2StageStr
		case OptOpSignal:
			argTable.IsOpt = true
		case LoptOpSignal:
			argTable.IsLopt = true
		case ValueOptSignal:
			argTable.IsValue = true
			expectedNext = ExpectUp2HyphenStr
		case ArgOpSignal:
			argTable.IsArg = true
			expectedNext = ExpectUp2HyphenStr
		case
			SingleOpSignal,
			SingleShortOpSignal:
			argTable.QuoteTypeSignal = SingleQuote
		case
			NoQuoteOpSignal,
			NoQuoteShortOpSignal:
			argTable.QuoteTypeSignal = NoQuote
		default:
			argTable.UnkownOption = inputArg
			expectedNext = ExpectNormalStr
		}

		argTables = append(argTables, argTable)
	}
	return argTables
}
