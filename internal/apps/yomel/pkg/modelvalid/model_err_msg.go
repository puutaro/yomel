package modelvalid

import "github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"

const (
	stageDescriptionIrregular = "'" + argtables.StageSignal + "'" + " description must be meaning sentence\ndesc: '%s', stageNo: %d"
	cmdStrNotFound            = "'" + argtables.CmdOpSignal + "'" + " value is require\ndesc: '%s', stageNo: %d"
	svcStrNotFound            = "'" + argtables.SvcOpSignal + "'" + " value is require\ndesc: '%s', stageNo: %d"
	actStrNotFound            = "'" + argtables.ActOpSignal + "'" + " value is require\ndesc: '%s', stageNo: %d"
)
