package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

// Test_checkColorCodeStrIrregularErr verifies that checkColorCodeStrIrregularErr correctly validates color code strings in StageModel.
func Test_checkColorCodeStrIrregularErr(t *testing.T) {
	tests := []struct {
		name    string
		stModel model.StageModel
		wantErr bool
	}{
		{
			name: "should return nil when all color codes are valid or empty",
			stModel: model.StageModel{
				No:              1,
				ColorStr:        "",
				BgColorStr:      "",
				CommentColorStr: "",
			},
			wantErr: false,
		},
		{
			name: "should return error when ColorStr is invalid",
			stModel: model.StageModel{
				No:              1,
				ColorStr:        "invalid-color",
				BgColorStr:      "",
				CommentColorStr: "",
			},
			wantErr: true,
		},
		{
			name: "should return error when BgColorStr is invalid",
			stModel: model.StageModel{
				No:              2,
				ColorStr:        "",
				BgColorStr:      "invalid-bg-color",
				CommentColorStr: "",
			},
			wantErr: true,
		},
		{
			name: "should return error when CommentColorStr is invalid",
			stModel: model.StageModel{
				No:              3,
				ColorStr:        "",
				BgColorStr:      "",
				CommentColorStr: "invalid-comment-color",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkColorCodeStrIrregularErr(tt.stModel)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
