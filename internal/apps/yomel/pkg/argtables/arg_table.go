package argtables

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
	Str             *string
}

func GenArgTable(inputArgs []string) []ArgTable {

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
		case StageArgName:
			stageNum++
			argTable.StageNo = stageNum
			argTable.IsStage = true
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
		case ArgOpSignal:
			argTable.IsArg = true
		case
			SingleOpSignal,
			SingleShortOpSignal:
			argTable.QuoteTypeSignal = SingleQuote
		case
			NoQuoteOpSignal,
			NoQuoteShortOpSignal:
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
