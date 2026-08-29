package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/stretchr/testify/assert"
)

// Test_execCheckOptStrBlankErr verifies that execCheckOptStrBlankErr correctly returns an error when option strings are blank.
func Test_execCheckOptStrBlankErr(t *testing.T) {
	tests := []struct {
		name      string
		opts      []model.OptParam
		stageNo   int
		wantError string
	}{
		{
			name:      "should return nil when options slice is empty",
			opts:      []model.OptParam{},
			stageNo:   1,
			wantError: "",
		},
		{
			name: "should return nil when option strings are not blank",
			opts: []model.OptParam{
				{OptStr: "f"},
				{OptStr: "verbose"},
			},
			stageNo:   1,
			wantError: "",
		},
		{
			name: "should return error when option string is blank",
			opts: []model.OptParam{
				{OptStr: ""},
			},
			stageNo:   1,
			wantError: "'-o' and '--o' str must not be blank\nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := execCheckOptStrBlankErr(tt.opts, tt.stageNo)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
