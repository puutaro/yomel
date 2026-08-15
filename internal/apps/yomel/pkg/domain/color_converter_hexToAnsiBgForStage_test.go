package domain

import (
	"testing"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/color"
)

func TestHexToAnsiBgForStage(t *testing.T) {
	tests := []struct {
		name    string
		hexSrc  string
		ctrlHex string
		want    string
	}{
		{
			name:    "hexSrc is empty, uses ctrlHex converted",
			hexSrc:  "",
			ctrlHex: "red",
			want:    hexToAnsi(color.ConvertFixColorStr("red"), BackgroundAnsi),
		},
		{
			name:    "hexSrc is not empty, uses hexSrc directly",
			hexSrc:  "#ff0000",
			ctrlHex: "blue",
			want:    hexToAnsi("#ff0000", BackgroundAnsi),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hexToAnsiBgForStage(tt.hexSrc, tt.ctrlHex)
			if got != tt.want {
				t.Errorf("hexToAnsiBgForStage() = %v, want %v", got, tt.want)
			}
		})
	}
}
