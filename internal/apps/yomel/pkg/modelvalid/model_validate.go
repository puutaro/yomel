package modelvalid

import (
	"fmt"
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
)

func ModelValidate(stModels []model.StageModel) error {
	for _, stModel := range stModels {
		isStageDescErr := validateByIrregularStageDesc(stModel)
		if isStageDescErr != nil {
			return isStageDescErr
		}
	}
	return nil
}

func validateByIrregularStageDesc(stModel model.StageModel) error {
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
