package modelvalid

import "github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"

const (
	stageDescriptionIrregular = "'" + argtables.StageSignal + "'" + " description must be meaning sentence\ndesc: '%s', stageNo: %d"

	stageNoSuffix        = "\nstageNo: %d"
	valueRequireMsgAfter = " value is required" + stageNoSuffix
	cmdStrRequire        = "'" + argtables.CmdOpSignal + "'" + valueRequireMsgAfter
	svcStrRequire        = "'" + argtables.SvcOpSignal + "'" + valueRequireMsgAfter
	actStrRequire        = "'" + argtables.ActOpSignal + "'" + valueRequireMsgAfter
)
