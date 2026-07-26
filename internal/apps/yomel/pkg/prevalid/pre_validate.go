package prevalid

import (
	"errors"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/args"
)

func PreValidate(argTables []args.ArgTable) error {
	isStageErr := validateByIsStage(argTables)
	if isStageErr != nil {
		return isStageErr
	}
	return nil
}

func validateByIsStage(argTables []args.ArgTable) error {

	for _, argTable := range argTables {
		if argTable.IsStage {
			return nil
		}
	}
	return errors.New(stageNotFound)
}
