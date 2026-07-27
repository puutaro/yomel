package argtablevalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

const (
	logFlagWithQuote       = "'" + argtables.LogFlagSignal + "'"
	noLogFlagWithQuote     = "'" + argtables.NoLogFlagSignal + "'"
	logNoLogSingnalWithAnd = logFlagWithQuote + " and " + noLogFlagWithQuote

	versionWithQuote      = "'" + argtables.VersionOpSignal + "'"
	helpWithQuote         = "'" + argtables.HelpOpSignal + "'"
	logFilterWithQuote    = "'" + argtables.LogFilterOpSignal + "'"
	errLogFilterWithQuote = "'" + argtables.ErrLogFilterOpSignal + "'"

	cmdOpNameWithQuote = "'" + argtables.CmdOpSignal + "'"
	svcOpNameWithQuote = "'" + argtables.SvcOpSignal + "'"
	actOpNameWithQuote = "'" + argtables.ActOpSignal + "'"

	argOpSignalWithQuote     = "'" + argtables.ArgOpSignal + "'"
	optOpSignalWithQuote     = "'" + argtables.OptOpSignal + "'"
	singleOpSignalWithQuote  = "'" + argtables.SingleOpSignal + "/" + argtables.SingleShortOpSignal + "'"
	noQuoteOpSignalWithQuote = "'" + argtables.NoQuoteOpSignal + "/" + argtables.NoQuoteShortOpSignal + "'"
	quoteSingnalWithAnd      = singleOpSignalWithQuote + " and " + noQuoteOpSignalWithQuote

	stageNoSuffix = "\nstageNo: %d"

	// noLogSpecifyNotMeaning = noLogFlagWithQuote + "not meaning\n stageNo: 0"

	stageNotFound = "'" + argtables.StageSignal + "'" + " not found"
	cmdNotFound   = cmdOpNameWithQuote + " not found" + stageNoSuffix

	duplicateionErrSuffix  = " duplication" + stageNoSuffix
	versionDuplicate       = versionWithQuote + duplicateionErrSuffix
	helpDuplicate          = helpWithQuote + duplicateionErrSuffix
	logFilterDuplicate     = logFilterWithQuote + duplicateionErrSuffix
	errLogFilterDuplicate  = errLogFilterWithQuote + duplicateionErrSuffix
	noLogLogFlagDuplidate  = logNoLogSingnalWithAnd + duplicateionErrSuffix
	svcDuplidate           = svcOpNameWithQuote + duplicateionErrSuffix
	cmdDuplidate           = cmdOpNameWithQuote + duplicateionErrSuffix
	actDuplidate           = actOpNameWithQuote + duplicateionErrSuffix
	quoteOpSignalDuplicate = quoteSingnalWithAnd + duplicateionErrSuffix
	argQuoteWithAnd        = argOpSignalWithQuote + " and " + optOpSignalWithQuote

	cmdSvcActOrdrerIrregular = "Must be" + cmdOpNameWithQuote + "->" + svcOpNameWithQuote + "->" + actOpNameWithQuote + "order " + stageNoSuffix

	quoteOptionIrregularPosition = quoteSingnalWithAnd + " must be immediately after " + argQuoteWithAnd + stageNoSuffix
)
