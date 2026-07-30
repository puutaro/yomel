package modelvalid

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/model"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_checkNoBlankStrRequireErr(t *testing.T) {
	tests := []struct {
		name      string
		input     model.StageModel
		wantError string
	}{
		{
			name: "should return nil when required strings are not blank",
			input: model.StageModel{
				No:  1,
				Cmd: "echo",
			},
			wantError: "",
		},
		{
			name: "should return error when svc string is blank",
			input: model.StageModel{
				No:  1,
				Svc: testutil.Ptr(""),
			},
			wantError: "'-svc' no blank str is required\nstageNo: 1",
		},
		{
			name: "should return error when act string is blank",
			input: model.StageModel{
				No:  1,
				Act: testutil.Ptr(""),
			},
			wantError: "'-act' no blank str is required\nstageNo: 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkNoBlankStrRequireErr(tt.input)
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}
