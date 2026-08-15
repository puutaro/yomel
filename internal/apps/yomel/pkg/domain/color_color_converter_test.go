package domain

import (
	"testing"
)

func TestParseHexColor(t *testing.T) {
	tests := []struct {
		name  string
		hex   string
		wantR int
		wantG int
		wantB int
	}{
		{
			name:  "normal hex with hash",
			hex:   "#ff0000",
			wantR: 255,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "normal hex without hash",
			hex:   "00ff00",
			wantR: 0,
			wantG: 255,
			wantB: 0,
		},
		{
			name:  "color name via ConvertFixColorStr",
			hex:   "red",
			wantR: 255,
			wantG: 0,
			wantB: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := parseHexColor(tt.hex)
			if gotR != tt.wantR || gotG != tt.wantG || gotB != tt.wantB {
				t.Errorf("parseHexColor() = (%v, %v, %v), want (%v, %v, %v)", gotR, gotG, gotB, tt.wantR, tt.wantG, tt.wantB)
			}
		})
	}
}
