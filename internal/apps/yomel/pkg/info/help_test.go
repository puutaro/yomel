// Write direct above line for Comment on code
package info

import (
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_GetHelpByDefault(t *testing.T) {
	var info YomelInfo
	toml.Decode(YomelInfoRaw, &info)
	description := info.Yomel.Description
	tests := []struct {
		name      string
		input     []string
		wantHelp  *string
		wantError error
	}{
		{
			name:  "should return help content when help flag is present in control section",
			input: []string{},
			wantHelp: testutil.Ptr(
				strings.Join(
					[]string{description, "", detail},
					"\n",
				),
			),
			wantError: nil,
		},
		{
			name:      "should return nil when input args are not empty",
			input:     []string{"--log"},
			wantHelp:  nil,
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHelp, gotErr := GetHelpByDefault(tt.input)
			if tt.wantError != nil {
				assert.EqualError(t, gotErr, tt.wantError.Error())
			} else {
				assert.NoError(t, gotErr)
			}
			if tt.wantHelp != nil {
				assert.Equal(t, *tt.wantHelp, *gotHelp)
			} else {
				assert.Equal(t, tt.wantHelp, gotHelp)
			}
		})
	}
}

func Test_GetHelpByOption(t *testing.T) {
	var info YomelInfo
	toml.Decode(YomelInfoRaw, &info)
	description := info.Yomel.Description
	expectedHelp := strings.Join(
		[]string{description, "", detail},
		"\n",
	)

	tests := []struct {
		name      string
		input     []argtables.ArgTableDto
		wantHelp  *string
		wantError error
	}{
		{
			name: "should return help content when IsHelp is true in control section",
			input: []argtables.ArgTableDto{
				{StageNo: 0, IsHelp: true},
			},
			wantHelp:  &expectedHelp,
			wantError: nil,
		},
		{
			name: "should return nil when help option is not triggered",
			input: []argtables.ArgTableDto{
				{StageNo: 1, IsStage: true},
			},
			wantHelp:  nil,
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHelp, gotErr := GetHelpByOption(tt.input)
			if tt.wantError != nil {
				assert.EqualError(t, gotErr, tt.wantError.Error())
			} else {
				assert.NoError(t, gotErr)
			}
			if tt.wantHelp != nil {
				assert.Equal(t, *tt.wantHelp, *gotHelp)
			} else {
				assert.Equal(t, tt.wantHelp, gotHelp)
			}
		})
	}
}
