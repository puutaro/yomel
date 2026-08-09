package descjudger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_IsBellowSingleCharRepeated verifies that IsBellowSingleCharRepeated correctly detects strings composed of repeated single characters or empty strings.
func Test_IsBellowSingleCharRepeated(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "should return true when input is empty string",
			input: "",
			want:  true,
		},
		{
			name:  "should return true when input is single character",
			input: "a",
			want:  true,
		},
		{
			name:  "should return true when input is repeated single character",
			input: "aaa",
			want:  true,
		},
		{
			name:  "should return true when input is repeated multibyte single character",
			input: "あああ",
			want:  true,
		},
		{
			name:  "should return false when input is normal varied sentence",
			input: "hello world",
			want:  false,
		},
		{
			name:  "should return false when input is valid descriptive Japanese sentence",
			input: "テスト用の説明文",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBellowSingleCharRepeated(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
