package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

// Test_CtrlModeValidate verifies that CtrlModeValidate correctly validates control model parameters based on stage length.
func Test_CtrlModeValidate(t *testing.T) {
	tests := []struct {
		name      string
		ctrlModel model.ControlModel
		stageLen  int
		wantError string
	}{
		{
			name: "should return nil when stageLen is 1 or less regardless of title",
			ctrlModel: model.ControlModel{
				Title: "aaa",
			},
			stageLen:  1,
			wantError: "",
		},
		{
			name: "should return nil when stageLen is greater than 1 and title is valid",
			ctrlModel: model.ControlModel{
				Title: "valid title",
			},
			stageLen:  2,
			wantError: "",
		},
		{
			name: "should return error when stageLen is greater than 1 and title is repeated single characters",
			ctrlModel: model.ControlModel{
				Title: "aaa",
			},
			stageLen:  2,
			wantError: "'///' must be meaning sentence if stage > 1\ntitle: 'aaa'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CtrlModeValidate(tt.ctrlModel, tt.stageLen)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
