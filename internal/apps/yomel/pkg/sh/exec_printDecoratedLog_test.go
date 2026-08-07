// Test_printDecoratedLog verifies that printDecoratedLog correctly outputs formatted logs to the writer under various conditions.
package sh

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// stripANSI removes ANSI escape sequences (colors, underlines, bold, etc.) from the string.
func stripANSI(str string) string {
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[mK]`)
	return ansiRegex.ReplaceAllString(str, "")
}

func Test_printDecoratedLog(t *testing.T) {
	tests := []struct {
		name                string
		no                  int
		desc                string
		cmdName             string
		logFilterShell      string
		errLogFilterShell   string
		stderrBuf           string
		stdoutBuf           string
		shouldStdErr        bool
		cmdHasError         bool
		firstPipeLogNewLine rune
		wantOutputSubstr    []string
	}{
		{
			name:                "should print decorated log with stdout and normal progress stderr when shouldStdErr is true",
			no:                  1,
			desc:                "test-stage",
			cmdName:             "echo 'hello'",
			logFilterShell:      "",
			errLogFilterShell:   "",
			stderrBuf:           "some progress info\n",
			stdoutBuf:           "hello\n",
			shouldStdErr:        true,
			cmdHasError:         false,
			firstPipeLogNewLine: '\n',
			wantOutputSubstr: []string{
				"Yomel-log[1]_",
				"Stage",
				"Test-stage",
				"Cmd",
				"echo 'hello'",
				"Progress",
				"some progress info",
				"Stdout",
				"hello",
			},
		},
		{
			name:                "should print decorated log with red error label when cmdHasError is true",
			no:                  2,
			desc:                "error-stage",
			cmdName:             "exit 1",
			logFilterShell:      "",
			errLogFilterShell:   "",
			stderrBuf:           "error occurred\n",
			stdoutBuf:           "",
			shouldStdErr:        true,
			cmdHasError:         true,
			firstPipeLogNewLine: ' ',
			wantOutputSubstr: []string{
				"Yomel-log[2]_",
				"Stage",
				"Error-stage",
				"Cmd",
				"exit 1",
				"Error",
				"error occurred",
			},
		},
		{
			name:                "should print only stdout when shouldStdErr is false",
			no:                  3,
			desc:                "stdout-only-stage",
			cmdName:             "ls",
			logFilterShell:      "",
			errLogFilterShell:   "",
			stderrBuf:           "",
			stdoutBuf:           "file1\nfile2\n",
			shouldStdErr:        false,
			cmdHasError:         false,
			firstPipeLogNewLine: '\n',
			wantOutputSubstr: []string{
				"Yomel-log[3]_",
				"Stage",
				"Stdout-only-stage",
				"Cmd",
				"ls",
				"Stdout",
				"file1",
				"file2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var combinedLog bytes.Buffer
			stderrBuffer := bytes.NewBufferString(tt.stderrBuf)
			stdoutBuffer := bytes.NewBufferString(tt.stdoutBuf)

			printDecoratedLog(
				&combinedLog,
				tt.no,
				tt.desc,
				tt.cmdName,
				tt.logFilterShell,
				tt.errLogFilterShell,
				stderrBuffer,
				stdoutBuffer,
				tt.shouldStdErr,
				tt.cmdHasError,
				tt.firstPipeLogNewLine,
			)

			// ANSIエスケープシーケンスを除去してから比較する
			output := stripANSI(combinedLog.String())

			for _, substr := range tt.wantOutputSubstr {
				assert.Contains(t, output, substr)
			}
		})
	}
}
