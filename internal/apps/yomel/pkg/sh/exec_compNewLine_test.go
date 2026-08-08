// Test_compNewLine verifies that compNewLine correctly computes the required number of newlines to append based on the existing content in the buffer or writer.
package sh

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compNewLine(t *testing.T) {
	tests := []struct {
		name       string
		initialBuf string
		newLineNum int
		want       string
	}{
		{
			name:       "should return full newlines when buffer is empty",
			initialBuf: "",
			newLineNum: 2,
			want:       "\n\n",
		},
		{
			name:       "should return remaining newlines when buffer ends with some newlines",
			initialBuf: "line1\n",
			newLineNum: 2,
			want:       "\n",
		},
		{
			name:       "should return empty string when buffer already has enough newlines",
			initialBuf: "line1\n\n",
			newLineNum: 2,
			want:       "",
		},
		{
			name:       "should return empty string when buffer has more than enough newlines",
			initialBuf: "line1\n\n\n",
			newLineNum: 2,
			want:       "",
		},
		{
			name:       "should work with generic io.Writer instead of bytes.Buffer",
			initialBuf: "",
			newLineNum: 1,
			want:       "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			if tt.initialBuf != "" {
				buf.WriteString(tt.initialBuf)
			}

			var got string
			if strings.Contains(tt.name, "generic io.Writer") {
				// bytes.Buffer 以外の一般的な io.Writer として渡す
				var w io.Writer = ioWriterOnly{buf: &buf}
				got = compNewLine(w, tt.newLineNum)
			} else {
				got = compNewLine(&buf, tt.newLineNum)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

type ioWriterOnly struct {
	buf *bytes.Buffer
}

func (w ioWriterOnly) Write(p []byte) (n int, err error) {
	return w.buf.Write(p)
}
