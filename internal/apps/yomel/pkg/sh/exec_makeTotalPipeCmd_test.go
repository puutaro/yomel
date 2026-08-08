// Test_makeTotalPipeCmd verifies that makeTotalPipeCmd correctly constructs a pipeline command string from stage information.
package sh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeTotalPipeCmd(t *testing.T) {
	tests := []struct {
		name       string
		stageInfos []StageInfo
		cmdStrFn   func(stageInfo StageInfo) string
		want       string
	}{
		{
			name: "should return single command string when single stage is provided",
			stageInfos: []StageInfo{
				{
					No:      1,
					CmdStrs: "echo 'hello'",
				},
			},
			cmdStrFn: func(stInfo StageInfo) string {
				return stInfo.CmdStrs
			},
			want: "echo 'hello'",
		},
		{
			name: "should join multiple command strings with backslash and pipe when multiple stages are provided",
			stageInfos: []StageInfo{
				{
					No:      1,
					CmdStrs: "echo 'line1'",
				},
				{
					No:      2,
					CmdStrs: "grep 'line1'",
				},
			},
			cmdStrFn: func(stInfo StageInfo) string {
				return stInfo.CmdStrs
			},
			want: "echo 'line1' \\\n| grep 'line1'",
		},
		{
			name:       "should return empty string when stageInfos is empty",
			stageInfos: []StageInfo{},
			cmdStrFn: func(stInfo StageInfo) string {
				return stInfo.CmdStrs
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeTotalPipeCmd(tt.stageInfos, tt.cmdStrFn)
			assert.Equal(t, tt.want, got)
		})
	}
}
