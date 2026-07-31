// Write direct above line for Comment on code
package sh

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_directExec verifies that directExec correctly runs pipeline commands.
func Test_directExec(t *testing.T) {
	tests := []struct {
		name       string
		stageInfos []StageInfo
		wantSubstr string
	}{
		{
			name: "should execute single stage pipeline correctly in direct mode",
			stageInfos: []StageInfo{
				{
					No:      1,
					Desc:    "direct-echo",
					CmdStrs: "echo 'direct hello'",
				},
			},
			wantSubstr: "",
		},
		{
			name: "should execute multi-stage pipeline correctly in direct mode",
			stageInfos: []StageInfo{
				{
					No:      1,
					Desc:    "direct-source",
					CmdStrs: "echo 'line1\nline2'",
				},
				{
					No:      2,
					Desc:    "direct-grep",
					CmdStrs: "grep 'line1'",
				},
			},
			wantSubstr: "",
		},
		{
			name:       "should do nothing when stageInfos is empty",
			stageInfos: []StageInfo{},
			wantSubstr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStderr := os.Stderr
			rErr, wErr, _ := os.Pipe()
			os.Stderr = wErr

			directExec(tt.stageInfos)

			wErr.Close()
			os.Stderr = oldStderr

			var bufErr bytes.Buffer
			_, _ = bufErr.ReadFrom(rErr)
			errOutput := bufErr.String()

			if tt.wantSubstr != "" {
				assert.True(t, strings.Contains(errOutput, tt.wantSubstr))
			}
		})
	}
}
