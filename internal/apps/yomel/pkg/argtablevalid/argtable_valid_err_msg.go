package argtablevalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtable"
)

const (
	stageNotFound = "'" + argtable.StageArgName + "'" + " not found"
	cmdNotFound   = "'" + argtable.CmdOpName + "'" + " not found\nstageNo: %d"
)
