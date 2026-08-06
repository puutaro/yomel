package modelvalid

import (
	"fmt"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
)

func CtrlModeValidate(ctrlModel model.ControlModel, stageLen int) error {
	fmt.Println("stageLne", stageLen)
	if stageLen <= 1 {
		return nil
	}
	title := ctrlModel.Title
	if isBellowSingleCharRepeated(title) {
		return fmt.Errorf(titleDescriptionIrregular, title)
	}
	return nil
}
