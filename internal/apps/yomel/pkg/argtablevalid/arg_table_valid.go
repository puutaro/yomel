package argtablevalid

import (
	"errors"
	"fmt"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtable"
)

func ArgTableValidate(argTables []argtable.ArgTable) error {
	if isStageErr := checkIsStage(argTables); isStageErr != nil {
		return isStageErr
	}
	if isCmdErr := checkIsCmd(argTables); isCmdErr != nil {
		return isCmdErr
	}
	return nil
}

func checkIsStage(argTables []argtable.ArgTable) error {
	for _, argTable := range argTables {
		if argTable.IsStage {
			return nil
		}
	}
	return errors.New(stageNotFound)
}
func checkIsCmd(argTables []argtable.ArgTable) error {
	stageNo := 0
	for _, argTable := range argTables {
		stageNo += incrementStageNo(argTable.IsStage)
		if argTable.IsCmd {
			return nil
		}
	}
	return fmt.Errorf(cmdNotFound, stageNo)
}

func incrementStageNo(isStage bool) int {
	if isStage {
		return 1
	}
	return 0
}
