package sh

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_directExec(t *testing.T) {
	tests := []struct {
		name         string
		stageInfos   []StageInfo
		wantContains string
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
			wantContains: "direct hello",
		},
		{
			name: "should handle invalid command with error message",
			stageInfos: []StageInfo{
				{
					No:      1,
					Desc:    "invalid-cmd",
					CmdStrs: "invalid_command_xyz",
				},
			},
			wantContains: "pipeline failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 標準出力をキャプチャするためのパイプを作成
			oldStdout := os.Stdout
			oldStderr := os.Stderr
			rOut, wOut, _ := os.Pipe()
			rErr, wErr, _ := os.Pipe()
			os.Stdout = wOut
			os.Stderr = wErr

			// テスト対象の実行
			totalPipeCmdStr := makeTotalPipeCmd(tt.stageInfos)
			directExec(totalPipeCmdStr)

			// パイプを閉じて出力を取得
			wOut.Close()
			wErr.Close()
			os.Stdout = oldStdout
			os.Stderr = oldStderr

			var bufOut bytes.Buffer
			var bufErr bytes.Buffer
			_, _ = bufOut.ReadFrom(rOut)
			_, _ = bufErr.ReadFrom(rErr)

			output := bufOut.String() + bufErr.String()

			if tt.wantContains != "" {
				assert.Contains(t, output, tt.wantContains)
			}
		})
	}
}
