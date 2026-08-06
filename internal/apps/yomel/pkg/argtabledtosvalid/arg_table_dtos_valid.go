package argtabledtosvalid

import (
	"fmt"
	"unicode"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtabledtos"
)

func ArgTableDtosValid(argTableDtos []argtabledtos.ArgTableDto) error {
	for _, argTableDto := range argTableDtos {
		checkParameterSuffixDescs := []*string{
			argTableDto.OptStr,
			argTableDto.LoptStr,
			argTableDto.ValueStr,
			argTableDto.ArgStr,
		}
		stageNo := argTableDto.StageNo
		for _, checkParaSuffixDesc := range checkParameterSuffixDescs {
			err := checkDescriptionSuffixMustBealPhanumericPascalCaseErrMsg(
				checkParaSuffixDesc,
				stageNo,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func checkDescriptionSuffixMustBealPhanumericPascalCaseErrMsg(
	str *string,
	stageNo int,
) error {
	if str == nil {
		return nil
	}
	if *str == "" {
		return nil
	}
	runes := []rune(*str)
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
