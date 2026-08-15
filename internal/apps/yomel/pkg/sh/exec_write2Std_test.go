package sh

import (
	"bytes"
	"testing"
)

func Test_write2Std(t *testing.T) {
	yl := &yomelLog{
		isTerminal: false,
	}

	t.Run("Empty buffer should not write anything", func(t *testing.T) {
		var buf bytes.Buffer
		stdBuf := new(bytes.Buffer)

		yl.write2Std(&buf, "Label", stdBuf, "", "\x1b[31m", false)

		if buf.Len() > 0 {
			t.Errorf("expected empty output, got %s", buf.String())
		}
	})

	t.Run("Non-empty buffer without filter shell", func(t *testing.T) {
		var buf bytes.Buffer
		stdBuf := bytes.NewBufferString("test output content")

		yl.write2Std(&buf, "Label", stdBuf, "", "\x1b[31m", false)

		output := buf.String()
		if output == "" {
			t.Errorf("expected some output, got empty")
		}
	})

	t.Run("Non-empty buffer with successful filter shell", func(t *testing.T) {
		var buf bytes.Buffer
		stdBuf := bytes.NewBufferString("line1\nline2")

		yl.write2Std(&buf, "Label", stdBuf, "cat", "\x1b[31m", false)

		output := buf.String()
		if output == "" {
			t.Errorf("expected filtered output, got empty")
		}
	})

	t.Run("Non-empty buffer with failing filter shell", func(t *testing.T) {
		var buf bytes.Buffer
		stdBuf := bytes.NewBufferString("line1\nline2")

		// Exit with error and write to stderr in filter shell
		yl.write2Std(&buf, "Label", stdBuf, "echo 'error message' >&2; exit 1", "\x1b[31m", true)

		output := buf.String()
		if output == "" {
			t.Errorf("expected filter shell error output, got empty")
		}
	})
}
