package domain

import (
	"testing"
)

func TestHexToAnsi(format *testing.T) {
	tests := []struct {
		name   string
		hex    string
		fOrB   ForeOrBack
		expect string
	}{
		{
			name:   "empty hex returns empty string",
			hex:    "",
			fOrB:   ForegroundAnsi,
			expect: "",
		},
		{
			name:   "foreground red color",
			hex:    "#ff0000",
			fOrB:   ForegroundAnsi,
			expect: "\x1b[38;2;255;0;0m",
		},
		{
			name:   "background blue color without hash prefix",
			hex:    "0026ff",
			fOrB:   BackgroundAnsi,
			expect: "\x1b[48;2;0;38;255m",
		},
		{
			name:   "foreground fallback or default case",
			hex:    "123456",
			fOrB:   ForeOrBack(99), // 不明な値の場合のフォールバック (ForegroundColorAnsiCodeになる)
			expect: "\x1b[38;2;18;52;86m",
		},
	}

	for _, tt := range tests {
		format.Run(tt.name, func(t *testing.T) {
			actual := hexToAnsi(tt.hex, tt.fOrB)
			if actual != tt.expect {
				t.Errorf("hexToAnsi() = %v, expect %v", actual, tt.expect)
			}
		})
	}
}
