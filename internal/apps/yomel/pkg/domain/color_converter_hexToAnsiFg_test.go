package domain

import (
	"testing"
)

func TestHexToAnsiFg(t *testing.T) {
	tests := []struct {
		name   string
		hex    string
		expect string
	}{
		{
			name:   "Empty hex string",
			hex:    "",
			expect: "",
		},
		{
			name:   "Valid hex color",
			hex:    "#ff0000",
			expect: "\x1b[38;2;255;0;0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hexToAnsiFg(tt.hex)
			if got != tt.expect {
				t.Errorf("hexToAnsiFg() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
