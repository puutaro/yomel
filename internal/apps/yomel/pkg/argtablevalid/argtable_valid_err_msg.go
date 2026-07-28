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

	ctrlParameterWithAnd = versionWithQuote + " and " + helpWithQuote

	logFilterWithQuote    = "'" + argtables.LogFilterOpSignal + "'"
	errLogFilterWithQuote = "'" + argtables.ErrLogFilterOpSignal + "'"

	cmdOpNameWithQuote = "'" + argtables.CmdOpSignal + "'"
	svcOpNameWithQuote = "'" + argtables.SvcOpSignal + "'"
	actOpNameWithQuote = "'" + argtables.ActOpSignal + "'"

	argOpSignalWithQuote     = "'" + argtables.ArgOpSignal + "'"
	valueOpSignalWithQuote   = "'" + argtables.ValueOptSignal + "'"
	argValueWithAnd          = argOpSignalWithQuote + " and " + valueOpSignalWithQuote
	optOpSignalWithQuote     = "'" + argtables.OptOpSignal + "'"
	singleOpSignalWithQuote  = "'" + argtables.SingleOpSignal + "/" + argtables.SingleShortOpSignal + "'"
	noQuoteOpSignalWithQuote = "'" + argtables.NoQuoteOpSignal + "/" + argtables.NoQuoteShortOpSignal + "'"
	quoteSingnalWithAnd      = singleOpSignalWithQuote + " and " + noQuoteOpSignalWithQuote

	stageNoSuffix   = "\nstageNo: %d"
	ctrlFieldSuffix = " in stage 0 field"

	onlyOneStr            = " are only one"
	onlyOneErrStageSuffix = onlyOneStr + " in each stage field" + stageNoSuffix

	duplicationStr        = " are duplicates"
	duplicateionErrSuffix = duplicationStr + stageNoSuffix
	argQuoteWithAnd       = argOpSignalWithQuote + " and " + optOpSignalWithQuote

	unknownParameterSpecifyedErrMsg     = "'%s' is unknown option" + stageNoSuffix
	ctrParameterOnlyOneErrMsg           = ctrlParameterWithAnd + onlyOneStr + ctrlFieldSuffix
	stageParameterSpecifyedInCtrlErrMsg = "'%s' must be specfied in stage field"
	stageNotFound                       = "'" + argtables.StageSignal + "'" + " not found"
	cmdNotFound                         = cmdOpNameWithQuote + " not found" + stageNoSuffix

	ctrlParameterSpecifyInStageErrMsg = "Must be specified" + ctrlParameterWithAnd + ctrlFieldSuffix + stageNoSuffix

	onlyOneErrMsg = "%s" + onlyOneStr + " in each stage field" + stageNoSuffix

	cmdSvcActOrdrerIrregular = "Must be" + cmdOpNameWithQuote + "->" + svcOpNameWithQuote + "->" + actOpNameWithQuote + "order " + stageNoSuffix

	quoteOptionIrregularPosition = argValueWithAnd + " must be immediately after " + argQuoteWithAnd + stageNoSuffix
)
