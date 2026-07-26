package prevalid

import (
	"errors"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
)

func PreValidate(argTables []args.ArgTable) error {
	isStageErr := checkIsStage(argTables)
	if isStageErr != nil {
		return isStageErr
	}
	return nil
}

func checkIsStage(argTables []args.ArgTable) error {

	for _, argTable := range argTables {
		if argTable.IsStage {
			return nil
		}
	}
	return errors.New(stageNotFound)
}
