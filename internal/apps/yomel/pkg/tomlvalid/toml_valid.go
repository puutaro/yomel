package tomlvalid

import (
	"fmt"
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/color"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
)

func TomlValidate(
	logConfig toml.LogConfig,
	tomlPath string,
) error {
	checkers := []func(toml.LogConfig, string) error{
		checkColorCodeStrIrregularErr,
		checkColorCodeStrIrregularErrForTitleColor,
	}
	for _, checker := range checkers {
		err := checker(
			logConfig,
			tomlPath,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkColorCodeStrIrregularErr(
	logConfig toml.LogConfig,
	tomlPath string,
) error {
	for _, colorStrSrc := range []string{
		logConfig.Color.Color,
		logConfig.Color.BgColor,
		logConfig.Color.CommentColor,
	} {
		colorStr := strings.ReplaceAll(
			colorStrSrc,
			color.OrderOperator,
			color.AndOperator,
		)
		err := color.DetectColorCodeStrIrregularErr(colorStr)
		if err != nil {
			return fmt.Errorf(
				colorStrIrregularErrMsg,
				tomlPath,
				err,
			)
		}
	}
	return nil
}

func checkColorCodeStrIrregularErrForTitleColor(
	logConfig toml.LogConfig,
	tomlPath string,
) error {
	for _, colorStr := range []string{
		logConfig.Color.TitleColor,
		logConfig.Color.TitleBgColor,
		logConfig.Color.TitleCommentColor,
	} {
		err := color.DetectColorCodeStrIrregularErr(colorStr)
		if err != nil {
			return fmt.Errorf(
				titleColorStrIrregularErrMsg,
				tomlPath,
				err,
			)
		}
	}
	return nil
}
