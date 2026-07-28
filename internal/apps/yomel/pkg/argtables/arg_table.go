package argtables

import (
	"strings"
)

const (
	Version            = "version"
	Help               = "help"
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

func GenArgTable(inputArgs []string) []ArgTable {

	var argTables []ArgTable
	stageNum := 0
	enableHypenPrefix := false
	for i := 0; i < len(inputArgs); i++ {
		displayNum := i + 1
		inputArg := inputArgs[i]
		argTable := ArgTable{
			No:      displayNum,
			StageNo: stageNum,
		}
		switch true {
		case inputArg == VersionOpSignal &&
			!enableHypenPrefix:
			argTable.IsVersion = true
		case inputArg == HelpOpSignal &&
			!enableHypenPrefix:
			argTable.IsHelp = true
		case inputArg == LogFlagSignal &&
			!enableHypenPrefix:
			argTable.IsLog = true
		case inputArg == NoLogFlagSignal &&
			!enableHypenPrefix:
			argTable.IsNoLog = true
		case inputArg == LogFilterOpSignal &&
			!enableHypenPrefix:
			argTable.IsLogFilter = true
		case inputArg == ErrLogFilterOpSignal &&
			!enableHypenPrefix:
			argTable.IsErrLogFilter = true
		case inputArg == StageArgName:
			stageNum++
			argTable.StageNo = stageNum
			argTable.IsStage = true
		case inputArg == CmdOpSignal &&
			!enableHypenPrefix:
			argTable.IsCmd = true
		case inputArg == SvcOpSignal &&
			!enableHypenPrefix:
			argTable.IsSvc = true
		case inputArg == ActOpSignal &&
			!enableHypenPrefix:
			argTable.IsAct = true
		case inputArg == OptOpSignal &&
			!enableHypenPrefix:
			argTable.IsOpt = true
		case inputArg == LoptOpSignal &&
			!enableHypenPrefix:
			argTable.IsLopt = true
		case inputArg == ValueOptSignal &&
			!enableHypenPrefix:
			argTable.IsValue = true
			enableHypenPrefix = true
		case inputArg == ArgOpSignal &&
			!enableHypenPrefix:
			argTable.IsArg = true
			enableHypenPrefix = true
		case
			inputArg == SingleOpSignal,
			inputArg == SingleShortOpSignal:
			argTable.QuoteTypeSignal = SingleQuote
		case
			inputArg == NoQuoteOpSignal &&
				!enableHypenPrefix,
			inputArg == NoQuoteShortOpSignal &&
				!enableHypenPrefix:
			argTable.QuoteTypeSignal = NoQuote
		default:
			if !strings.HasPrefix(inputArg, "-") ||
				enableHypenPrefix {
				argTable.Str = &inputArg
			} else {
				argTable.UnkownOption = inputArg
			}
			enableHypenPrefix = false
		}
		argTables = append(
			argTables,
			argTable,
		)
	}
	return argTables
}
