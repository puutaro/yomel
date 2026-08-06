package modelvalid

import "github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"

const (
	stageNoSuffix        = "\nstageNo: %d"
	curTitle             = "\ntitle: '%s'"
	stageDesc            = "\ndesc: '%s'"
	titleWithQuote       = "'" + argtables.TitleSignal + "'"
	stageSignalWithQuote = "'" + argtables.StageSignal + "'"

	titleDescriptionIrregular = titleWithQuote + " must be meaning sentence if stage > 1" + curTitle
	stageDescriptionIrregular = stageSignalWithQuote + " description must be meaning sentence" + stageNoSuffix + stageDesc
	stageDescriptionDuplicate = stageSignalWithQuote + " description must be unique across stages" + stageNoSuffix + stageDesc

	noBlankStrRequireErrMsg = "'%s' no blank str is required" + stageNoSuffix
	// cmdStrRequire        = "'" + argtables.CmdOpSignal + "'" + noBlankStrRequireErrMsg
	// svcStrRequire        = "'" + argtables.SvcOpSignal + "'" + noBlankStrRequireErrMsg
	// actStrRequire        = "'" + argtables.ActOpSignal + "'" + noBlankStrRequireErrMsg
)
