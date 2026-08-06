package argtablevalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

const (
	optOpSignalWithQuote   = "'" + argtables.OptOpSignal + "'"
	lOptOpSignalWithQuote  = "'" + argtables.LoptOpSignal + "'"
	argOpSignalWithQuote   = "'" + argtables.ArgOpSignal + "'"
	valueOpSignalWithQuote = "'" + argtables.ValueOptSignal + "'"
	opLopArgValueWithAnd   = optOpSignalWithQuote + " and " + lOptOpSignalWithQuote + " and " + argOpSignalWithQuote + " and " + valueOpSignalWithQuote
	stageNoSuffix          = "\nstageNo: %d"

	descriptionSuffixMustBealPhanumericPascalCaseErrMsg = "Description suffix is must be alphanumeric pascalCase in " + opLopArgValueWithAnd + stageNoSuffix
)
