package argtables

import (
	"slices"
	"strings"
)

const (
	Version              = "version"
	Help                 = "help"
	DirectMode           = "direct"
	GenMode              = "gen"
	Title                = "title"
	StageArgName         = "stage"
	NoLiveStdoutFlagName = "no-live-stdout"
	NoLiveStderrFlagName = "no-live-stderr"
	LogOpName            = "log"
	NoLogOpName          = "no-log"
	LogFilter            = "log-filter"
	ErrLogFilter         = "err-log-filter"
	CmdOpName            = "cmd"
	SvcOpName            = "svc"
	ActOpName            = "act"
	OptOpName            = "opt"
	LoptOpName           = "lop"
	ArgOpName            = "arg"
	ValueOpName          = "val"
	SingleOpName         = "single"
	SingleShortOpName    = "s"
	NoQuoteOpName        = "no-quote"
	NoQuoteShortOpName   = "n"
)
const (
	VersionOpSignal        = "--" + Version
	HelpOpSignal           = "--" + Help
	DirectModeFlagSignal   = "--" + DirectMode
	GenModeFlagSignal      = "--" + GenMode
	TitleSignal            = Title
	StageSignal            = StageArgName
	CmdOpSignal            = "-" + CmdOpName
	NoLiveStdoutFlagSignal = "--" + NoLiveStdoutFlagName
	NoLiveStderrFlagSignal = "--" + NoLiveStderrFlagName
	LogFlagSignal          = "--" + LogOpName
	NoLogFlagSignal        = "--" + NoLogOpName
	LogFilterOpSignal      = "--" + LogFilter
	ErrLogFilterOpSignal   = "--" + ErrLogFilter
	SvcOpSignal            = "-" + SvcOpName
	ActOpSignal            = "-" + ActOpName
	OptOpSignal            = "--" + OptOpName
	LoptOpSignal           = "--" + LoptOpName
	ArgOpSignal            = "--" + ArgOpName
	ValueOptSignal         = "--" + ValueOpName
	SingleOpSignal         = "--" + SingleOpName
	SingleShortOpSignal    = "--" + SingleShortOpName
	NoQuoteOpSignal        = "--" + NoQuoteOpName
	NoQuoteShortOpSignal   = "--" + NoQuoteShortOpName
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
	IsNoLiveStdout  bool
	IsNoLiveStderr  bool
	IsTitle         bool
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
	Comment         string
	QuoteTypeSignal QuoteType
	UnknownOption   string
	Str             *string
}
type ExpectedNext int

const (
	ExpectNormalStr ExpectedNext = iota
	ExpectUp2PureParameterStr
	ExpectUp2HyphenStr
)

// TODO judge option or str about --s, --n
// I escape this judge for rare case considering --s, --n as str
// So this implement apply --s, --n as option always
func GenArgTable(inputArgs []string) []ArgTable {
	var argtables []ArgTable
	stageNum := 0
	expectedNext := ExpectNormalStr

	pureStrParameters := []string{
		TitleSignal,
		StageSignal,
	}
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
			argtables = append(argtables, argTable)
			expectedNext = ExpectNormalStr
			continue
		case
			expectedNext == ExpectUp2PureParameterStr &&
				!strings.HasPrefix(inputArg, "-"):
			val := inputArg
			argTable.Str = &val
			argtables = append(argtables, argTable)
			expectedNext = ExpectNormalStr
			continue
		case
			expectedNext == ExpectNormalStr &&
				!strings.HasPrefix(inputArg, "-") &&
				!slices.Contains(pureStrParameters, inputArg):
			val := inputArg
			argTable.Str = &val
			argtables = append(argtables, argTable)
			expectedNext = ExpectNormalStr
			continue
		}

		switch true {
		case inputArg == VersionOpSignal:
			argTable.IsVersion = true
		case inputArg == HelpOpSignal:
			argTable.IsHelp = true
		case inputArg == GenModeFlagSignal:
			argTable.IsGen = true
		case inputArg == DirectModeFlagSignal:
			argTable.IsDirect = true
		case inputArg == NoLiveStdoutFlagSignal:
			argTable.IsNoLiveStdout = true
		case inputArg == NoLiveStderrFlagSignal:
			argTable.IsNoLiveStderr = true
		case inputArg == TitleSignal:
			argTable.IsTitle = true
			expectedNext = ExpectUp2PureParameterStr
		case inputArg == LogFlagSignal:
			argTable.IsLog = true
		case inputArg == NoLogFlagSignal:
			argTable.IsNoLog = true
		case inputArg == LogFilterOpSignal:
			argTable.IsLogFilter = true
		case inputArg == ErrLogFilterOpSignal:
			argTable.IsErrLogFilter = true
		// stage interpret as special arg always in this line by first switch sentence
		case inputArg == StageArgName:
			stageNum++
			argTable.StageNo = stageNum
			argTable.IsStage = true
			expectedNext = ExpectUp2PureParameterStr
		case inputArg == CmdOpSignal:
			argTable.IsCmd = true
			expectedNext = ExpectUp2PureParameterStr
		case inputArg == SvcOpSignal:
			argTable.IsSvc = true
			expectedNext = ExpectUp2PureParameterStr
		case inputArg == ActOpSignal:
			argTable.IsAct = true
			expectedNext = ExpectUp2PureParameterStr
		case strings.HasPrefix(inputArg, OptOpSignal):
			argTable.IsOpt = true
			argTable.Comment = removePrefix(inputArg, OptOpSignal)
		case strings.HasPrefix(inputArg, LoptOpSignal):
			argTable.IsLopt = true
			argTable.Comment = removePrefix(inputArg, LoptOpSignal)
		case strings.HasPrefix(inputArg, ValueOptSignal):
			argTable.IsValue = true
			argTable.Comment = removePrefix(inputArg, ValueOptSignal)
			expectedNext = ExpectUp2HyphenStr
		case strings.HasPrefix(inputArg, ArgOpSignal):
			argTable.IsArg = true
			argTable.Comment = removePrefix(inputArg, ArgOpSignal)
			expectedNext = ExpectUp2HyphenStr
		case
			inputArg == SingleOpSignal,
			inputArg == SingleShortOpSignal:
			argTable.QuoteTypeSignal = SingleQuote
		case
			inputArg == NoQuoteOpSignal,
			inputArg == NoQuoteShortOpSignal:
			argTable.QuoteTypeSignal = NoQuote
		default:
			argTable.UnknownOption = inputArg
			expectedNext = ExpectNormalStr
		}

		argtables = append(argtables, argTable)
	}
	return argtables
}

func removePrefix(str, prefix string) string {
	rest, _ := strings.CutPrefix(str, prefix)
	return rest
}
