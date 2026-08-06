package sh

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		firstPipeLogNewLine string
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
			firstPipeLogNewLine: "\n",
			wantOutputSubstr: []string{
				"\n#### YOMEL-LOG[1]_",
				"# Stage: \nTest-stage",
				"# Cmd: \necho 'hello'",
				"# Progress:",
				"some progress info",
				"# Stdout:",
				"hello",
				"####",
			},
		},
		{
			name:              "should print decorated log with red error label when cmdHasError is true",
			no:                2,
			desc:              "error-stage",
			cmdName:           "exit 1",
			logFilterShell:    "",
			errLogFilterShell: "",
			stderrBuf:         "error occurred\n",
			stdoutBuf:         "",
			shouldStdErr:      true,
			cmdHasError:       true,
			wantOutputSubstr: []string{
				"#### YOMEL-LOG[2]_",
				"# Stage: \nError-stage",
				"# Cmd: \nexit 1",
				"#\x1b[31m Error:\x1b[0m",
				"error occurred",
				"####",
			},
		},
		{
			name:              "should print only stdout when shouldStdErr is false",
			no:                3,
			desc:              "stdout-only-stage",
			cmdName:           "ls",
			logFilterShell:    "",
			errLogFilterShell: "",
			stderrBuf:         "",
			stdoutBuf:         "file1\nfile2\n",
			shouldStdErr:      false,
			cmdHasError:       false,
			wantOutputSubstr: []string{
				"#### YOMEL-LOG[3]_",
				"# Stage: \nStdout-only-stage",
				"# Cmd: \nls",
				"# Stdout:",
				"file1",
				"file2",
				"####",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStderr := os.Stderr
			rErr, wErr, _ := os.Pipe()
			os.Stderr = wErr

			stderrBuffer := bytes.NewBufferString(tt.stderrBuf)
			stdoutBuffer := bytes.NewBufferString(tt.stdoutBuf)

			printDecoratedLog(
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

			wErr.Close()
			os.Stderr = oldStderr

			var bufErr bytes.Buffer
			_, _ = bufErr.ReadFrom(rErr)
			output := bufErr.String()

			for _, substr := range tt.wantOutputSubstr {
				assert.Contains(t, output, substr)
			}
		})
	}
}
