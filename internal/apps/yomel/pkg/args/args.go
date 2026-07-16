package args

import (
	"os"
)

const (
	Version            = "version"
	Help               = "help"
	StageArgName       = "stage"
	LogOpName          = "log"
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
	versionOpSignal      = "--" + Version
	helpOpSignal         = "--" + Help
	cmdOpSignal          = "-" + CmdOpName
	logOpSignal          = "--" + LogOpName
	logFilterOpSignal    = "--" + LogFilter
	errLogFilterOpSignal = "--" + ErrLogFilter
	svcOpSignal          = "-" + SvcOpName
	actOpSignal          = "-" + ActOpName
	optOpSignal          = "--" + OptOpName
	loptOpSignal         = "--" + LoptOpName
	argOpSignal          = "--" + ArgOpName
	valueOptSignal       = "--" + ValueOpName
	singleOpSignal       = "--" + SingleOpName
	singleShortOpSignal  = "--" + SingleShortOpName
	noQuoteOpSignal      = "--" + NoQuoteOpName
	noQuoteShortOpSignal = "--" + NoQuoteShortOpName
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
	IsCmd           bool
	IsSvc           bool
	IsAct           bool
	IsOpt           bool
	IsLopt          bool
	IsValue         bool
	IsArg           bool
	QuoteTypeSignal QuoteType
	Str             *string
}

func GenArgTable() []ArgTable {

	inputArgs := os.Args[1:]

	var argTables []ArgTable
	stageNum := 0
	for i := 0; i < len(inputArgs); i++ {
		displayNum := i + 1
		inputArg := inputArgs[i]
		argTable := ArgTable{
			No:      displayNum,
			StageNo: stageNum,
		}
		switch inputArg {
		case versionOpSignal:
			argTable.IsVersion = true
		case helpOpSignal:
			argTable.IsHelp = true
		case logOpSignal:
			argTable.IsLog = true
		case logFilterOpSignal:
			argTable.IsLogFilter = true
		case errLogFilterOpSignal:
			argTable.IsErrLogFilter = true
		case StageArgName:
			stageNum++
			argTable.StageNo = stageNum
			argTable.IsStage = true
		case cmdOpSignal:
			argTable.IsCmd = true
		case svcOpSignal:
			argTable.IsSvc = true
		case actOpSignal:
			argTable.IsAct = true
		case optOpSignal:
			argTable.IsOpt = true
		case loptOpSignal:
			argTable.IsLopt = true
		case valueOptSignal:
			argTable.IsValue = true
		case argOpSignal:
			argTable.IsArg = true
		case
			singleOpSignal,
			singleShortOpSignal:
			argTable.QuoteTypeSignal = SingleQuote
		case
			noQuoteOpSignal,
			noQuoteShortOpSignal:
			argTable.QuoteTypeSignal = NoQuote
		default:
			argTable.Str = &inputArg
		}
		argTables = append(
			argTables,
			argTable,
		)
	}
	return argTables
}
