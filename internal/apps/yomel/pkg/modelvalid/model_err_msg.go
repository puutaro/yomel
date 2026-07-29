package modelvalid

import "github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"

const (
	stageDescriptionIrregular = "'" + argtables.StageSignal + "'" + " description must be meaning sentence\ndesc: '%s', stageNo: %d"

	stageNoSuffix           = "\nstageNo: %d"
	noBlankStrRequireErrMsg = "'%s' no blank str is required" + stageNoSuffix
	// cmdStrRequire        = "'" + argtables.CmdOpSignal + "'" + noBlankStrRequireErrMsg
	// svcStrRequire        = "'" + argtables.SvcOpSignal + "'" + noBlankStrRequireErrMsg
	// actStrRequire        = "'" + argtables.ActOpSignal + "'" + noBlankStrRequireErrMsg
)
