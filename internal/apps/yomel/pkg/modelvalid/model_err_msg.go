package modelvalid

import "github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"

const (
	stageDescriptionIrregular = "'" + argtables.StageArgName + "'" + " description must be meaning sentence\ndesc: '%s', stageNo: %d"
)
