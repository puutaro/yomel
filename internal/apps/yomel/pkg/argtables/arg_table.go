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
	for i, inputArg := range inputArgs {
		displayNum := i + 1
		argTable := ArgTable{
			No:      displayNum,
			StageNo: stageNum,
		}
		// flag don't infer enableHypenPrefix
		// no hyphen prefix and no-quote or singl-quote option
		switch inputArg {
		case
			StageArgName:
			stageNum++
			argTable.StageNo = stageNum
			argTable.IsStage = true
			argTables = append(argTables, argTable)
			continue
		case
			SingleOpSignal,
			SingleShortOpSignal:
			argTable.QuoteTypeSignal = SingleQuote
			argTables = append(argTables, argTable)
			continue
		case
			NoQuoteOpSignal,
			NoQuoteShortOpSignal:
			argTable.QuoteTypeSignal = NoQuote
			argTables = append(argTables, argTable)
			continue
		}
		if enableHypenPrefix || !strings.HasPrefix(inputArg, "-") {
			argTable.Str = &inputArg
			enableHypenPrefix = false
			argTables = append(argTables, argTable)
			continue
		}

		switch inputArg {
		case VersionOpSignal:
			argTable.IsVersion = true
		case HelpOpSignal:
			argTable.IsHelp = true
		case LogFlagSignal:
			argTable.IsLog = true
		case NoLogFlagSignal:
			argTable.IsNoLog = true
		case LogFilterOpSignal:
			argTable.IsLogFilter = true
		case ErrLogFilterOpSignal:
			argTable.IsErrLogFilter = true
		case CmdOpSignal:
			argTable.IsCmd = true
		case SvcOpSignal:
			argTable.IsSvc = true
		case ActOpSignal:
			argTable.IsAct = true
		case OptOpSignal:
			argTable.IsOpt = true
		case LoptOpSignal:
			argTable.IsLopt = true
		case ValueOptSignal:
			argTable.IsValue = true
			enableHypenPrefix = true
		case ArgOpSignal:
			argTable.IsArg = true
			enableHypenPrefix = true
		default:
			argTable.UnkownOption = inputArg
			enableHypenPrefix = false
		}
		argTables = append(
			argTables,
			argTable,
		)
	}
	return argTables
}
