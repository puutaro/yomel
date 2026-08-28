package argtablesvalid

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtablecounter"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/descjudger"
)

func ArgTableValidate(argTables []argtables.ArgTable) error {
	validators := []func([]argtables.ArgTable) error{
		checkUnknownOptionSpecifyedErrMsg,
		// checkCtrlParameterOnlyOne,
		checkStageParameterSpecifyInCtrlErr,
		checkIsStage,
		checkIsCmd,
		checkCtrlParameterSpecifyInStageErr,
		checkOnlyOneOptionErr,
		checkQuoteOptionIrregularPositionErr,
		checkDescriptionSuffixErr,
	}
	for _, validate := range validators {
		if err := validate(argTables); err != nil {
			return err
		}
	}
	return nil
}

func checkUnknownOptionSpecifyedErrMsg(argTables []argtables.ArgTable) error {
	stageNo := 0
	for _, argTable := range argTables {
		stageNo += argtablecounter.IncrementStageNo(argTable.IsStage)
		if argTable.UnknownOption != "" {
			return fmt.Errorf(
				unknownParameterSpecifyedErrMsg,
				argTable.UnknownOption,
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
		if argTable.IsVersion ||
			argTable.IsHelp ||
			argTable.IsDirect ||
			argTable.IsGen ||
			argTable.IsTitleColor ||
			argTable.IsTitleBgColor ||
			argTable.IsTitleCommentColor {
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

func checkDescriptionSuffixErr(argtables []argtables.ArgTable) error {
	for _, argTable := range argtables {
		if !argTable.IsOpt &&
			!argTable.IsLopt &&
			!argTable.IsValue &&
			!argTable.IsArg {
			continue
		}
		if err := checkDescriptionSuffixMustBeAlPhanumericPascalCaseErr(
			argTable.Comment,
			argTable.StageNo,
		); err != nil {
			return err
		}
		if err := checkDescriptionSuffixIrregularErrMsg(
			argTable.Comment,
			argTable.StageNo,
		); err != nil {
			return err
		}
	}
	return nil
}

func checkDescriptionSuffixMustBeAlPhanumericPascalCaseErr(
	str string,
	stageNo int,
) error {
	if str == "" {
		return nil
	}
	runes := []rune(str)
	if !unicode.IsUpper(runes[0]) || !unicode.IsLetter(runes[0]) {
		return fmt.Errorf(descriptionSuffixMustBealPhanumericPascalCaseErrMsg, stageNo)
	}
	for _, r := range runes[1:] {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return fmt.Errorf(descriptionSuffixMustBealPhanumericPascalCaseErrMsg, stageNo)
		}
	}
	return nil
}
func checkDescriptionSuffixIrregularErrMsg(
	str string,
	stageNo int,
) error {
	if str == "" {
		return nil
	}
	if descjudger.IsBellowSingleCharRepeated(str) {
		return fmt.Errorf(descriptionSuffixIrregularErrMsg, stageNo, str)
	}
	return nil
}
