package argtablevalid

import (
	"errors"
	"fmt"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

func ArgTableValidate(argTables []argtables.ArgTable) error {
	if ctrlParameterSpecifyInSatgeErr := checkCtrlParameterSpecifyInStageErr(
		argTables,
	); ctrlParameterSpecifyInSatgeErr != nil {
		return ctrlParameterSpecifyInSatgeErr
	}
	if stageParameterSpecifyInCtrlErr := checkStageParameterSpecifyInCtrl(
		argTables,
	); stageParameterSpecifyInCtrlErr != nil {
		return stageParameterSpecifyInCtrlErr
	}
	if versionDuplicateErr := checkCtrlParameterDuplidate(
		argTables,
	); versionDuplicateErr != nil {
		return versionDuplicateErr
	}
	if isStageErr := checkIsStage(
		argTables,
	); isStageErr != nil {
		return isStageErr
	}
	if isCmdErr := checkIsCmd(
		argTables,
	); isCmdErr != nil {
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

func checkCtrlParameterSpecifyInStageErr(argTables []argtables.ArgTable) error {
	stageNo := 0
	for _, argTable := range argTables {
		stageNo += incrementStageNo(argTable.IsStage)
		if stageNo == 0 {
			continue
		}
		if argTable.IsVersion || argTable.IsHelp {
			return fmt.Errorf(ctrlParameterSpecifyInSatgeErrMsg, stageNo)
		}
	}
	return nil
}

func checkCtrlParameterDuplidate(argTables []argtables.ArgTable) error {
	stageNo := 0
	count := 0
	for _, argTable := range argTables {
		stageNo += incrementStageNo(argTable.IsStage)
		if stageNo > 0 {
			break
		}
		if argTable.IsVersion || argTable.IsHelp {
			count++
		}
	}
	if count > 1 {
		return fmt.Errorf(ctrParameterDuplicate)
	}
	return nil
}

func checkStageParameterSpecifyInCtrl(
	argTables []argtables.ArgTable,
) error {
	checkers := []struct {
		targetParameterCheckFn func(argtable argtables.ArgTable) bool
		targetParameterSignal  string
	}{
		{
			targetParameterCheckFn: func(a argtables.ArgTable) bool { return a.IsCmd },
			targetParameterSignal:  argtables.CmdOpSignal,
		},
		{
			targetParameterCheckFn: func(a argtables.ArgTable) bool { return a.IsSvc },
			targetParameterSignal:  argtables.SvcOpSignal,
		},
		{
			targetParameterCheckFn: func(a argtables.ArgTable) bool { return a.IsAct },
			targetParameterSignal:  argtables.ActOpSignal,
		},
		{
			targetParameterCheckFn: func(a argtables.ArgTable) bool { return a.IsArg },
			targetParameterSignal:  argtables.ArgOpSignal,
		},
		{
			targetParameterCheckFn: func(a argtables.ArgTable) bool { return a.IsOpt },
			targetParameterSignal:  argtables.OptOpSignal,
		},
		{
			targetParameterCheckFn: func(a argtables.ArgTable) bool { return a.IsLopt },
			targetParameterSignal:  argtables.LoptOpSignal,
		},
	}

	for _, c := range checkers {
		if err := execCheckStageParameterSpecifyInCtrl(
			argTables,
			c.targetParameterCheckFn,
			c.targetParameterSignal,
		); err != nil {
			return err
		}
	}
	return nil
}
func execCheckStageParameterSpecifyInCtrl(
	argTables []argtables.ArgTable,
	isCheckParameter func(t argtables.ArgTable) bool,
	parameterSignal string,
) error {
	stageNo := 0
	for _, argTable := range argTables {
		stageNo += incrementStageNo(argTable.IsStage)
		if stageNo > 0 {
			break
		}
		if isCheckParameter(argTable) {
			return fmt.Errorf(
				stageParameterSpecifyedInCtrlErrMsg,
				parameterSignal,
			)
		}
	}
	return nil
}

func incrementStageNo(isStage bool) int {
	if isStage {
		return 1
	}
	return 0
}
