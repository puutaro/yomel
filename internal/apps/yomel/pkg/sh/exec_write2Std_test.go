package sh

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_write2Std verifies that write2Std correctly processes buffer output with or without a filter shell.
func Test_write2Std(t *testing.T) {
	tests := []struct {
		name        string
		label       string
		buffer      string
		filterShell string
		wantOutput  string
	}{
		{
			name:        "should write buffer directly when filterShell is empty and buffer ends with newline",
			label:       "stdout:\n",
			buffer:      "hello world\n",
			filterShell: "",
			wantOutput:  "stdout:\nhello world\n\n",
		},
		{
			name:        "should write buffer directly and append newline when filterShell is empty and buffer lacks newline",
			label:       "stdout:\n",
			buffer:      "hello world",
			filterShell: "",
			wantOutput:  "stdout:\nhello world\n\n",
		},
		{
			name:        "should do nothing when buffer is empty",
			label:       "stdout:\n",
			buffer:      "",
			filterShell: "",
			wantOutput:  "",
		},
		{
			name:        "should filter buffer using filterShell when filterShell is provided",
			label:       "stdout:\n",
			buffer:      "apple\nbanana\napplet\n",
			filterShell: "grep 'app'",
			wantOutput:  "stdout:\napple\napplet\n\n",
		},
		{
			name:        "should append newline to filtered output if missing",
			label:       "stdout:\n",
			buffer:      "apple\nbanana\napplet\n",
			filterShell: "grep 'app' | head -n 1 | tr -d '\n'",
			wantOutput:  "stdout:\napple\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			buf := bytes.NewBufferString(tt.buffer)

			write2Std(&out, tt.label, buf, tt.filterShell)
			assert.Equal(t, tt.wantOutput, out.String())
		})
	}
}
