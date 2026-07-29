package argtablevalid

import (
	"errors"
	"fmt"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtablecounter"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
)

func ArgTableValidate(argTables []argtables.ArgTable) error {
	validators := []func([]argtables.ArgTable) error{
		checkUnkownOptionSpecifyedErrMsg,
		// checkCtrlParameterOnlyOne,
		checkStageParameterSpecifyInCtrlErr,
		checkIsStage,
		checkIsCmd,
		checkCtrlParameterSpecifyInStageErr,
		checkOnlyOneOptionErr,
		checkCmdSvcActOrderErr,
		checkQuoteOptionIrregularPositionErr,
	}
	for _, validate := range validators {
		if err := validate(argTables); err != nil {
			return err
		}
	}
	return nil
}

func checkUnkownOptionSpecifyedErrMsg(argTables []argtables.ArgTable) error {
	stageNo := 0
	for _, argTable := range argTables {
		stageNo += argtablecounter.IncrementStageNo(argTable.IsStage)
		if argTable.UnkownOption != "" {
			return fmt.Errorf(
				unknownParameterSpecifyedErrMsg,
				argTable.UnkownOption,
				stageNo,
			)
		}
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
	stageNum := 0
	for _, argTable := range argTables {
		if !argTable.IsStage {
			continue
		}
		stageNum++
	}
	cmdCountList := make([]int, stageNum)
	currentStageIdx := -1
	for _, argTable := range argTables {
		if argTable.IsStage {
			currentStageIdx++
		}
		if argTable.IsCmd && currentStageIdx > -1 {
			cmdCountList[currentStageIdx]++
		}
	}
	for stageIndex, cmdNum := range cmdCountList {
		if cmdNum > 0 {
			continue
		}
		return fmt.Errorf(cmdNotFound, stageIndex+1)
	}
	return nil
}

func checkCtrlParameterSpecifyInStageErr(argTables []argtables.ArgTable) error {
	stageNo := 0
	for _, argTable := range argTables {
		stageNo += argtablecounter.IncrementStageNo(argTable.IsStage)
		if stageNo == 0 {
			continue
		}
		if argTable.IsVersion || argTable.IsHelp {
			return fmt.Errorf(ctrlParameterSpecifyInStageErrMsg, stageNo)
		}
	}
	return nil
}

// func checkCtrlParameterOnlyOne(argTables []argtables.ArgTable) error {
// 	stageNo := 0
// 	count := 0
// 	for _, argTable := range argTables {
// 		stageNo += argtablecounter.IncrementStageNo(argTable.IsStage)
// 		if stageNo > 0 {
// 			break
// 		}
// 		if argTable.IsVersion || argTable.IsHelp {
// 			count++
// 		}
// 	}
// 	if count > 1 {
// 		return fmt.Errorf(ctrParameterOnlyOneErrMsg)
// 	}
// 	return nil
// }

func checkStageParameterSpecifyInCtrlErr(
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
		stageNo += argtablecounter.IncrementStageNo(argTable.IsStage)
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
func checkOnlyOneOptionErr(
	argTables []argtables.ArgTable,
) error {
	checkers := []struct {
		targetParameterCheckFn func(argtable argtables.ArgTable) bool
		targetParameters       string
	}{
		{
			targetParameterCheckFn: func(
				a argtables.ArgTable,
			) bool {
				return a.IsLog || a.IsNoLog
			},
			targetParameters: logNoLogSingnalWithAnd,
		},
		{
			targetParameterCheckFn: func(
				a argtables.ArgTable,
			) bool {
				return a.IsLogFilter
			},
			targetParameters: logFilterWithQuote,
		},
		{
			targetParameterCheckFn: func(
				a argtables.ArgTable,
			) bool {
				return a.IsErrLogFilter
			},
			targetParameters: errLogFilterWithQuote,
		},
		{
			targetParameterCheckFn: func(
				a argtables.ArgTable,
			) bool {
				return a.IsCmd
			},
			targetParameters: cmdOpNameWithQuote,
		},
		{
			targetParameterCheckFn: func(
				a argtables.ArgTable,
			) bool {
				return a.IsSvc
			},
			targetParameters: svcOpNameWithQuote,
		},
		{
			targetParameterCheckFn: func(
				a argtables.ArgTable,
			) bool {
				return a.IsAct
			},
			targetParameters: actOpNameWithQuote,
		},
	}

	for _, c := range checkers {
		if err := execCheckOnlyOneOptionErr(
			argTables,
			c.targetParameterCheckFn,
			c.targetParameters,
		); err != nil {
			return err
		}
	}
	return nil
}
func execCheckOnlyOneOptionErr(
	argTables []argtables.ArgTable,
	targetPrameterCheckFn func(argtables.ArgTable) bool,
	targetParametersSignal string,
) error {
	stageNo := 0
	count := 0
	for _, argTable := range argTables {
		incStageNo := argtablecounter.IncrementStageNo(argTable.IsStage)
		if incStageNo > 0 {
			count = 0
		}
		stageNo += incStageNo
		if targetPrameterCheckFn(argTable) {
			count++
		}
		if count > 1 {
			return fmt.Errorf(
				onlyOneErrMsg,
				targetParametersSignal,
				stageNo,
			)
		}
	}
	return nil
}

func checkCmdSvcActOrderErr(argTables []argtables.ArgTable) error {
	stageNo := 0
	curMainArgNum := 0
	cmdOrder := 1
	svcOrder := 2
	actOrder := 3
	for _, argTable := range argTables {
		incStageNo := argtablecounter.IncrementStageNo(argTable.IsStage)
		if incStageNo > 0 {
			curMainArgNum = 0
		}
		switch true {
		case argTable.IsCmd:
			if curMainArgNum > cmdOrder {
				return fmt.Errorf(cmdSvcActOrdrerIrregularErrMsg, stageNo)
			}
			curMainArgNum = cmdOrder
		case argTable.IsSvc:
			if curMainArgNum > svcOrder {
				return fmt.Errorf(cmdSvcActOrdrerIrregularErrMsg, stageNo)
			}
			curMainArgNum = svcOrder
		case argTable.IsAct:
			if curMainArgNum > actOrder {
				return fmt.Errorf(cmdSvcActOrdrerIrregularErrMsg, stageNo)
			}
			curMainArgNum = actOrder
		}
		stageNo += incStageNo
	}
	return nil
}
func checkQuoteOptionIrregularPositionErr(
	argTables []argtables.ArgTable,
) error {
	stageNo := 0
	for index, argTable := range argTables {
		stageNo += argtablecounter.IncrementStageNo(argTable.IsStage)
		if argTable.QuoteTypeSignal == argtables.DoubleQuote {
			continue
		}
		if index <= 0 {
			continue
		}
		prevArgTable := argTables[index-1]
		if prevArgTable.IsValue ||
			prevArgTable.IsArg {
			continue
		}
		return fmt.Errorf(
			quoteOptionIrregularPositionErrMsg,
			stageNo,
		)
	}
	return nil
}
