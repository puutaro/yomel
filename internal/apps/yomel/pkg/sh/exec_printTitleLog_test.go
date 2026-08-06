package sh

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_printTitleLog(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		wantOutput string
	}{
		{
			name:       "should print title when title is not empty",
			title:      "my-pipeline-title",
			wantOutput: "\x1b[1m#### YOMEL-TITLE:\x1b[0m\n\x1b[1mmy-pipeline-title\x1b[0m\n",
		},
		{
			name:       "should do nothing when title is empty",
			title:      "",
			wantOutput: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStderr := os.Stderr
			rErr, wErr, _ := os.Pipe()
			os.Stderr = wErr

			printTitleLog(tt.title)

			wErr.Close()
			os.Stderr = oldStderr

			var bufErr bytes.Buffer
			_, _ = bufErr.ReadFrom(rErr)
			gotOutput := bufErr.String()

			assert.Equal(t, tt.wantOutput, gotOutput)
		})
	}
}
