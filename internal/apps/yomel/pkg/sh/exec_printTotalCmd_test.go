// Test_printTotalCmd verifies that printTotalCmd correctly outputs the formatted total command pipeline string.
package sh

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_printTotalCmd(t *testing.T) {
	tests := []struct {
		name            string
		title           string
		totalPipeCmdStr string
		want            string
	}{
		{
			name:            "should print total command with side comments formatted correctly",
			title:           "test-title",
			totalPipeCmdStr: "echo 'hello'    `#first comment`\necho 'world'    `#second comment`",
			want:            "\x1b[4m\x1b[1mT\x1b[22motal-cmd\x1b[24m\necho 'hello'    `#first comment`\necho 'world'    `#second comment`\n\n",
		},
		{
			name:            "should handle single line total command string correctly",
			title:           "simple-title",
			totalPipeCmdStr: "echo 'simple'",
			want:            "\x1b[4m\x1b[1mT\x1b[22motal-cmd\x1b[24m\necho 'simple'\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var combinedLog bytes.Buffer
			yl := &yomelLog{}

			yl.printTotalCmd(&combinedLog, tt.totalPipeCmdStr)

			assert.Equal(t, tt.want, combinedLog.String())
		})
	}
}
