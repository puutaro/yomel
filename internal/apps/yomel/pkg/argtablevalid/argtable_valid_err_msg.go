package argtablevalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

const (
	stageNotFound            = "'" + argtables.StageSignal + "'" + " not found"
	cmdOpNameWithQuote       = "'" + argtables.CmdOpSignal + "'"
	cmdNotFound              = cmdOpNameWithQuote + " not found\nstageNo: %d"
	cmdDuplidate             = cmdOpNameWithQuote + " duplication\nstageNo: %d"
	svcOpNameWithQuote       = "'" + argtables.SvcOpSignal + "'"
	actOpNameWithQuote       = "'" + argtables.ActOpSignal + "'"
	cmdSvcActOrdrerIrregular = "Must be" + cmdOpNameWithQuote + "->" + svcOpNameWithQuote + "->" + actOpNameWithQuote + "order \nstageNo: %d"

	argOpSignalWithQuote     = "'" + argtables.ArgOpSignal + "'"
	optOpSignalWithQuote     = "'" + argtables.OptOpSignal + "'"
	singleOpSignalWithQuote  = "'" + argtables.SingleOpSignal + "/" + argtables.SingleShortOpSignal + "'"
	noQuoteOpSignalWithQuote = "'" + argtables.NoQuoteOpSignal + "/" + argtables.NoQuoteShortOpSignal + "'"
	// Include duplicate quoteOption specifed
	quoteOptionIrregularPosition = singleOpSignalWithQuote + " and " + noQuoteOpSignalWithQuote + " must be immediately after " + argOpSignalWithQuote + " and " + optOpSignalWithQuote + "\nstageNo: %d"
)
