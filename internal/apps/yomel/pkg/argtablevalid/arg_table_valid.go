package argtablevalid

import (
	"errors"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argTable"
)

func ArgTableValidate(argTables []argTable.ArgTable) error {
	isStageErr := checkIsStage(argTables)
	if isStageErr != nil {
		return isStageErr
	}
	return nil
}

func checkIsStage(argTables []argTable.ArgTable) error {

	for _, argTable := range argTables {
		if argTable.IsStage {
			return nil
		}
	}
	return errors.New(stageNotFound)
}
