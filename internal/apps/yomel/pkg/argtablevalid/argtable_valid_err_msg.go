package argtablevalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
)

const (
	logFlagWithQuote       = "'" + argtabledtos.LogFlagSignal + "'"
	noLogFlagWithQuote     = "'" + argtabledtos.NoLogFlagSignal + "'"
	logNoLogSingnalWithAnd = logFlagWithQuote + " and " + noLogFlagWithQuote

	versionWithQuote = "'" + argtabledtos.VersionOpSignal + "'"
	helpWithQuote    = "'" + argtabledtos.HelpOpSignal + "'"
	genWithQuote     = "'" + argtabledtos.GenModeFlagSignal + "'"
	directWithQuote  = "'" + argtabledtos.DirectModeFlagSignal + "'"

	ctrlParameterWithAnd = versionWithQuote + " and " + helpWithQuote + " and " + genWithQuote + " and " + directWithQuote

	logFilterWithQuote    = "'" + argtabledtos.LogFilterOpSignal + "'"
	errLogFilterWithQuote = "'" + argtabledtos.ErrLogFilterOpSignal + "'"

	cmdOpNameWithQuote = "'" + argtabledtos.CmdOpSignal + "'"
	svcOpNameWithQuote = "'" + argtabledtos.SvcOpSignal + "'"
	actOpNameWithQuote = "'" + argtabledtos.ActOpSignal + "'"

	cmdSvcActOpNameWithAnd = cmdOpNameWithQuote + " and " + svcOpNameWithQuote + " and " + actOpNameWithQuote

	argOpSignalWithQuote     = "'" + argtabledtos.ArgOpSignal + "'"
	valueOpSignalWithQuote   = "'" + argtabledtos.ValueOptSignal + "'"
	argValueWithAnd          = argOpSignalWithQuote + " and " + valueOpSignalWithQuote
	optOpSignalWithQuote     = "'" + argtabledtos.OptOpSignal + "'"
	singleOpSignalWithQuote  = "'" + argtabledtos.SingleShortOpSignal + "/" + argtabledtos.SingleOpSignal + "'"
	noQuoteOpSignalWithQuote = "'" + argtabledtos.NoQuoteShortOpSignal + "/" + argtabledtos.NoQuoteOpSignal + "'"
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
	stageNotFound                       = "'" + argtabledtos.StageSignal + "'" + " not found"
	cmdNotFound                         = cmdOpNameWithQuote + " not found" + stageNoSuffix

	ctrlParameterSpecifyInStageErrMsg = "must be specified" + ctrlParameterWithAnd + ctrlFieldSuffix + stageNoSuffix

	onlyOneErrMsg = "%s" + onlyOneStr + " in each stage field" + stageNoSuffix

	cmdSvcActOrdrerIrregularErrMsg = cmdSvcActOpNameWithAnd + " must be" + cmdOpNameWithQuote + " -> " + svcOpNameWithQuote + " -> " + actOpNameWithQuote + " order " + stageNoSuffix

	quoteOptionIrregularPositionErrMsg = quoteSingnalWithAnd + " must be immediately after " + argValueWithAnd + stageNoSuffix
)
