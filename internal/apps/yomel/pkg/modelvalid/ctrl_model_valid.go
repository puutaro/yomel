package modelvalid

import (
	"fmt"
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/color"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/descjudger"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
)

func CtrlModeValidate(ctrlModel model.ControlModel, stageLen int) error {
	if stageLen <= 1 {
		return nil
	}
	title := ctrlModel.Title
	if descjudger.IsBellowSingleCharRepeated(title) {
		return fmt.Errorf(titleDescriptionIrregular, title)
	}
	if err := checkColorCodeStrIrregularErrForCtrl(ctrlModel); err != nil {
		return err
	}
	return nil
}
func checkColorCodeStrIrregularErrForCtrl(ctrlModel model.ControlModel) error {
	for _, colorStr := range []string{
		ctrlModel.TitleColorStr,
		ctrlModel.TitleBgColorStr,
		ctrlModel.TitleCommentColorStr,
	} {
		err := color.DetectColorStrIrregularErrForCtrl(
			colorStr,
			titleColorStrIrregularErrMsg,
		)
		if err != nil {
			return err
		}
	}
	for _, colorStrSrc := range []string{
		ctrlModel.ColorStr,
		ctrlModel.BgColorStr,
		ctrlModel.CommentColorStr,
	} {
		colorStr := strings.ReplaceAll(
			colorStrSrc,
			color.OrderOperator,
			color.AndOperator,
		)
		err := color.DetectColorStrIrregularErrForCtrl(
			colorStr,
			colorStrIrregularErrMsgForCtrl,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
