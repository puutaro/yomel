package sh

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_addNewline verifies that addNewline appends a newline character to the file if the buffer does not end with one.
func Test_addNewline(t *testing.T) {
	tests := []struct {
		name       string
		buffer     string
		wantOutput string
	}{
		{
			name:       "should append newline when buffer does not end with newline",
			buffer:     "no newline",
			wantOutput: "\n\n",
		},
		{
			name:       "should not append newline when buffer already ends with newline",
			buffer:     "already has newline\n",
			wantOutput: "",
		},
		{
			name:       "should do nothing and not panic when buffer is empty",
			buffer:     "",
			wantOutput: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			buf := bytes.NewBufferString(tt.buffer)

			addNewline(&out, buf)
			assert.Equal(t, tt.wantOutput, string(out.String()))
		})
	}
}
