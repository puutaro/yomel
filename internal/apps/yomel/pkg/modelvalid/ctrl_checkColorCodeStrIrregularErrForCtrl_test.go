// Test_checkColorCodeStrIrregularErrForCtrl verifies that checkColorCodeStrIrregularErrForCtrl correctly validates color code strings in ControlModel.
package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_checkColorCodeStrIrregularErrForCtrl(t *testing.T) {
	tests := []struct {
		name      string
		ctrlModel model.ControlModel
		wantError bool
	}{
		{
			name: "should pass when all color code strings are valid",
			ctrlModel: model.ControlModel{
				TitleColorStr:        "blue",
				TitleBgColorStr:      "azure",
				TitleCommentColorStr: "blue",
				ColorStr:             "red",
				BgColorStr:           "green",
				CommentColorStr:      "blue",
			},
			wantError: false,
		},
		{
			name: "should return error when title color string is irregular/invalid",
			ctrlModel: model.ControlModel{
				TitleColorStr:        "invalid-color-code-xyz",
				TitleBgColorStr:      "azure",
				TitleCommentColorStr: "blue",
				ColorStr:             "red",
				BgColorStr:           "green",
				CommentColorStr:      "blue",
			},
			wantError: true,
		},
		{
			name: "should return error when title background color string is irregular/invalid",
			ctrlModel: model.ControlModel{
				TitleColorStr:        "blue",
				TitleBgColorStr:      "invalid-color-code-xyz",
				TitleCommentColorStr: "blue",
				ColorStr:             "red",
				BgColorStr:           "green",
				CommentColorStr:      "blue",
			},
			wantError: true,
		},
		{
			name: "should return error when title comment color string is irregular/invalid",
			ctrlModel: model.ControlModel{
				TitleColorStr:        "blue",
				TitleBgColorStr:      "azure",
				TitleCommentColorStr: "invalid-color-code-xyz",
				ColorStr:             "red",
				BgColorStr:           "green",
				CommentColorStr:      "blue",
			},
			wantError: true,
		},
		{
			name: "should return error when normal color string is irregular/invalid",
			ctrlModel: model.ControlModel{
				TitleColorStr:        "blue",
				TitleBgColorStr:      "azure",
				TitleCommentColorStr: "blue",
				ColorStr:             "invalid-color-code-xyz",
				BgColorStr:           "green",
				CommentColorStr:      "blue",
			},
			wantError: true,
		},
		{
			name: "should return error when background color string is irregular/invalid",
			ctrlModel: model.ControlModel{
				TitleColorStr:        "blue",
				TitleBgColorStr:      "azure",
				TitleCommentColorStr: "blue",
				ColorStr:             "red",
				BgColorStr:           "invalid-color-code-xyz",
				CommentColorStr:      "blue",
			},
			wantError: true,
		},
		{
			name: "should return error when comment color string is irregular/invalid",
			ctrlModel: model.ControlModel{
				TitleColorStr:        "blue",
				TitleBgColorStr:      "azure",
				TitleCommentColorStr: "blue",
				ColorStr:             "red",
				BgColorStr:           "green",
				CommentColorStr:      "invalid-color-code-xyz",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkColorCodeStrIrregularErrForCtrl(tt.ctrlModel)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
