package argtablevalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

const (
	logFlagWithQuote       = "'" + argtables.LogFlagSignal + "'"
	noLogFlagWithQuote     = "'" + argtables.NoLogFlagSignal + "'"
	logNoLogSingnalWithAnd = logFlagWithQuote + " and " + noLogFlagWithQuote

	versionWithQuote = "'" + argtables.VersionOpSignal + "'"
	helpWithQuote    = "'" + argtables.HelpOpSignal + "'"
	genWithQuote     = "'" + argtables.GenModeFlagSingnal + "'"
	directWithQuote  = "'" + argtables.DirectModeFlagSignal + "'"

	ctrlParameterWithAnd = versionWithQuote + " and " + helpWithQuote + " and " + genWithQuote + " and " + directWithQuote

	logFilterWithQuote    = "'" + argtables.LogFilterOpSignal + "'"
	errLogFilterWithQuote = "'" + argtables.ErrLogFilterOpSignal + "'"

	cmdOpNameWithQuote = "'" + argtables.CmdOpSignal + "'"
	svcOpNameWithQuote = "'" + argtables.SvcOpSignal + "'"
	actOpNameWithQuote = "'" + argtables.ActOpSignal + "'"

	cmdSvcActOpNameWithAnd = cmdOpNameWithQuote + " and " + svcOpNameWithQuote + " and " + actOpNameWithQuote

	argOpSignalWithQuote     = "'" + argtables.ArgOpSignal + "'"
	valueOpSignalWithQuote   = "'" + argtables.ValueOptSignal + "'"
	argValueWithAnd          = argOpSignalWithQuote + " and " + valueOpSignalWithQuote
	optOpSignalWithQuote     = "'" + argtables.OptOpSignal + "'"
	singleOpSignalWithQuote  = "'" + argtables.SingleShortOpSignal + "/" + argtables.SingleOpSignal + "'"
	noQuoteOpSignalWithQuote = "'" + argtables.NoQuoteShortOpSignal + "/" + argtables.NoQuoteOpSignal + "'"
	quoteSingnalWithAnd      = singleOpSignalWithQuote + " and " + noQuoteOpSignalWithQuote

	stageNoSuffix   = "\nstageNo: %d"
	ctrlFieldSuffix = " in stage 0 field"

	onlyOneStr            = " are only one"
	onlyOneErrStageSuffix = onlyOneStr + " in each stage field" + stageNoSuffix

	duplicationStr        = " are duplicates"
	duplicateionErrSuffix = duplicationStr + stageNoSuffix
	argQuoteWithAnd       = argOpSignalWithQuote + " and " + optOpSignalWithQuote

	unknownParameterSpecifyedErrMsg = "'%s' is unknown option" + stageNoSuffix

	// this validateion omit, becuase I wont to soft for --version and --help judge
	// ctrParameterOnlyOneErrMsg           = ctrlParameterWithAnd + onlyOneStr + ctrlFieldSuffix
	stageParameterSpecifyedInCtrlErrMsg = "'%s' must be specfied in stage field"
	stageNotFound                       = "'" + argtables.StageSignal + "'" + " not found"
	cmdNotFound                         = cmdOpNameWithQuote + " not found" + stageNoSuffix

	ctrlParameterSpecifyInStageErrMsg = "must be specified" + ctrlParameterWithAnd + ctrlFieldSuffix + stageNoSuffix

	onlyOneErrMsg = "%s" + onlyOneStr + " in each stage field" + stageNoSuffix

	cmdSvcActOrdrerIrregularErrMsg = cmdSvcActOpNameWithAnd + " must be" + cmdOpNameWithQuote + " -> " + svcOpNameWithQuote + " -> " + actOpNameWithQuote + " order " + stageNoSuffix

	quoteOptionIrregularPositionErrMsg = quoteSingnalWithAnd + " must be immediately after " + argValueWithAnd + stageNoSuffix
)
