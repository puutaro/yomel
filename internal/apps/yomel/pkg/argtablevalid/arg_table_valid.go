package argtablevalid

import (
	"errors"
	"fmt"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

func ArgTableValidate(argTables []argtables.ArgTable) error {
	if versionDuplicateErr := checkVersionDuplidate(argTables); versionDuplicateErr != nil {
		return versionDuplicateErr
	}
	if isStageErr := checkIsStage(argTables); isStageErr != nil {
		return isStageErr
	}
	if isCmdErr := checkIsCmd(argTables); isCmdErr != nil {
		return isCmdErr
	}
	return nil
}

func checkIsStage(argTables []argtables.ArgTable) error {
	for _, argTable := range argTables {
		if argTable.IsStage {
			return nil
		}
	}
	return errors.New(stageNotFound)
}
func checkIsCmd(argTables []argtables.ArgTable) error {
	stageNo := 0
	for _, argTable := range argTables {
		stageNo += incrementStageNo(argTable.IsStage)
		if argTable.IsCmd {
			return nil
		}
	}
	return fmt.Errorf(cmdNotFound, stageNo)
}

func checkVersionDuplidate(argTables []argtables.ArgTable) error {
	stageNo := 0
	count := 0
	for _, argTable := range argTables {
		stageNo += incrementStageNo(argTable.IsStage)
		if stageNo > 0 {
			break
		}
		if argTable.IsVersion {
			count++
		}
	}
	if count > 1 {
		return fmt.Errorf(versionDuplicate, 0)
	}
	return nil
}

func incrementStageNo(isStage bool) int {
	if isStage {
		return 1
	}
	return 0
}
