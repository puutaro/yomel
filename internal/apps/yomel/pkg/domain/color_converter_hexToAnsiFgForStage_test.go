package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/color"
)

func TestHexToAnsiFgForStage(t *testing.T) {
	tests := []struct {
		name     string
		hexSrc   string
		ctrlHex  string
		expected string
	}{
		{
			name:     "hexSrc is not empty",
			hexSrc:   "#ff0000",
			ctrlHex:  "blue",
			expected: hexToAnsi("#ff0000", ForegroundAnsi),
		},
		{
			name:     "hexSrc is empty and ctrlHex uses fix color string",
			hexSrc:   "",
			ctrlHex:  "red",
			expected: hexToAnsi(color.ConvertFixColorStr("red"), ForegroundAnsi),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hexToAnsiFgForStage(tt.hexSrc, tt.ctrlHex)
			if result != tt.expected {
				t.Errorf("hexToAnsiFgForStage() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
