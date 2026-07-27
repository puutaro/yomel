package argtablevalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

const (
	stageNotFound = "'" + argtables.StageArgName + "'" + " not found"
	cmdNotFound   = "'" + argtables.CmdOpName + "'" + " not found\nstageNo: %d"
)
