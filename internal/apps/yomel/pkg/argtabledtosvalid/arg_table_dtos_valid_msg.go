package argtabledtosvalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
)

const (
	optOpSignalWithQuote   = "'" + argtabledtos.OptOpSignal + "'"
	lOptOpSignalWithQuote  = "'" + argtabledtos.LoptOpSignal + "'"
	argOpSignalWithQuote   = "'" + argtabledtos.ArgOpSignal + "'"
	valueOpSignalWithQuote = "'" + argtabledtos.ValueOptSignal + "'"
	opLopArgValueWithAnd   = optOpSignalWithQuote + " and " + lOptOpSignalWithQuote + " and " + argOpSignalWithQuote + " and " + valueOpSignalWithQuote
	stageNoSuffix          = "\nstageNo: %d"

	descriptionSuffixMustBealPhanumericPascalCaseErrMsg = opLopArgValueWithAnd + " Description suffix is must be alphanumeric pascalCase in " + opLopArgValueWithAnd + stageNoSuffix
)
