package modelvalid

import "github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"

const (
	stageDescriptionIrregular = "'" + argtables.StageSignal + "'" + " description must be meaning sentence\ndesc: '%s', stageNo: %d"

	cmdStrNotFound = "'" + argtables.CmdOpSignal + "'" + " value is required\nstageNo: %d"
	svcStrNotFound = "'" + argtables.SvcOpSignal + "'" + " value is required\nstageNo: %d"
	actStrNotFound = "'" + argtables.ActOpSignal + "'" + " value is required\nstageNo: %d"
)
