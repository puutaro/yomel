package modelvalid

import "github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"

const (
	stageNoSuffix             = "\nstageNo: %d"
	stageDesc                 = "\ndesc: '%s'"
	stageSignalWithQuote      = "'" + argtables.StageSignal + "'"
	stageDescriptionIrregular = stageSignalWithQuote + " description must be meaning sentence" + stageNoSuffix + stageDesc
	stageDescriptionDuplicate = stageSignalWithQuote + " description must be unique across stages" + stageNoSuffix + stageDesc

	noBlankStrRequireErrMsg = "'%s' no blank str is required" + stageNoSuffix
	// cmdStrRequire        = "'" + argtables.CmdOpSignal + "'" + noBlankStrRequireErrMsg
	// svcStrRequire        = "'" + argtables.SvcOpSignal + "'" + noBlankStrRequireErrMsg
	// actStrRequire        = "'" + argtables.ActOpSignal + "'" + noBlankStrRequireErrMsg
)
