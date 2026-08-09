package modelvalid

import (
	"fmt"
	"slices"
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
)

func ModelValidate(stModels []model.StageModel) error {
	validators := []func(model.StageModel) error{
		checkIrregularStageDesc,
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
	wholeValidators := []func([]model.StageModel) error{
		checkStageDescriptionDuplicate,
		checkOptStrBlankErrMsg,
	}
	for _, validate := range wholeValidators {
		if err := validate(stModels); err != nil {
			return err
		}
	}
	return nil
}

func checkIrregularStageDesc(stModel model.StageModel) error {
	if isBellowSingleCharRepeated(stModel.Desc) {
		return fmt.Errorf(stageDescriptionIrregular, stModel.No, stModel.Desc)
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
func checkStageDescriptionDuplicate(stModels []model.StageModel) error {
	var descList []string
	for _, stModel := range stModels {
		desc := stModel.Desc
		if !slices.Contains(descList, desc) {
			descList = append(descList, desc)
			continue
		}
		stageNo := len(descList) + 1
		return fmt.Errorf(stageDescriptionDuplicate, stageNo, desc)
	}
	return nil
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

func checkOptStrBlankErrMsg(stModels []model.StageModel) error {
	for _, stModel := range stModels {
		stageNo := stModel.No
		optPrams := [][]model.OptParam{
			stModel.CmdOps,
			stModel.CmdLops,
			stModel.SvcOps,
			stModel.SvcLops,
			stModel.ActOps,
			stModel.ActLops,
		}
		for _, optParam := range optPrams {
			err := execCheckOptStrBlankErrMsg(
				optParam,
				stageNo,
			)
			if err == nil {
				continue
			}
			return err
		}
	}
	return nil
}

func execCheckOptStrBlankErrMsg(
	opts []model.OptParam,
	stageNo int,
) error {
	if len(opts) == 0 {
		return nil
	}
	for _, opt := range opts {
		if opt.OptStr != "" {
			continue
		}
		return fmt.Errorf(optStrBlankErrMsg, stageNo)
	}
	return nil
}
