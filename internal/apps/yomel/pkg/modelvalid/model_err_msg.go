package modelvalid

import "github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"

const (
	stageNoSuffix             = "\nstageNo: %d"
	stageDesc                 = "\ndesc: '%s'"
	stageSignalWithQuote      = "'" + argtabledtos.StageSignal + "'"
	stageDescriptionIrregular = stageSignalWithQuote + " description must be meaning sentence" + stageNoSuffix + stageDesc
	stageDescriptionDuplicate = stageSignalWithQuote + " description must be unique across stages" + stageNoSuffix + stageDesc

	noBlankStrRequireErrMsg = "'%s' no blank str is required" + stageNoSuffix
	// cmdStrRequire        = "'" + argtables.CmdOpSignal + "'" + noBlankStrRequireErrMsg
	// svcStrRequire        = "'" + argtables.SvcOpSignal + "'" + noBlankStrRequireErrMsg
	// actStrRequire        = "'" + argtables.ActOpSignal + "'" + noBlankStrRequireErrMsg
)
