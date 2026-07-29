package modelvalid

import (
	"fmt"
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
)

func ModelValidate(stModels []model.StageModel) error {
	validators := []func(model.StageModel) error{
		checkByIrregularStageDesc,
		checkNoBlankStrRequireErrForCmd,
		checkNoBlankStrRequireErr,
	}
	for _, stModel := range stModels {
		for _, validate := range validators {
			if err := validate(stModel); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkByIrregularStageDesc(stModel model.StageModel) error {
	if isBellowSingleCharRepeated(stModel.Desc) {
		return fmt.Errorf(stageDescriptionIrregular, stModel.Desc, stModel.No)
	}
	return nil
}

func isBellowSingleCharRepeated(s string) bool {
	trimmed := strings.Trim(s, " 　")
	if len(trimmed) == 0 {
		return true
	}
	runes := []rune(trimmed)
	seen := make(map[rune]bool)
	for _, r := range runes {
		seen[r] = true
		if len(seen) > 1 {
			return false
		}
	}
	return len(seen) == 1
}

func checkNoBlankStrRequireErrForCmd(stModel model.StageModel) error {
	trimmed := strings.Trim(stModel.Cmd, " 　")
	if trimmed == "" {
		return fmt.Errorf(
			noBlankStrRequireErrMsg,
			argtables.CmdOpSignal,
			stModel.No,
		)
	}
	return nil
}

func checkNoBlankStrRequireErr(stModel model.StageModel) error {
	checkers := []struct {
		str     *string
		mainArg string
	}{
		{
			str:     stModel.Svc,
			mainArg: argtables.SvcOpSignal,
		},
		{
			str:     stModel.Act,
			mainArg: argtables.ActOpSignal,
		},
	}
	stageNo := stModel.No
	for _, c := range checkers {
		if err := execCheckNoBlankStrRequireErr(
			c.str,
			c.mainArg,
			stageNo,
		); err != nil {
			return err
		}
	}
	return nil
}
func execCheckNoBlankStrRequireErr(str *string, mainArg string, stageNo int) error {
	if str == nil {
		return nil
	}
	trimmed := strings.Trim(*str, " 　")
	if trimmed == "" {
		return fmt.Errorf(noBlankStrRequireErrMsg, mainArg, stageNo)
	}
	return nil
}
