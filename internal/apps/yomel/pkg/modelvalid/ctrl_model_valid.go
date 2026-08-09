package modelvalid

import (
	"fmt"

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
	return nil
}
