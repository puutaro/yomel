package domain

import (
	"testing"
)

func TestHexToAnsiBg(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		want string
	}{
		{
			name: "Empty hex string",
			hex:  "",
			want: "",
		},
		{
			name: "Valid hex string (Red)",
			hex:  "#ff0000",
			want: hexToAnsi("#ff0000", BackgroundAnsi),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hexToAnsiBg(tt.hex)
			if got != tt.want {
				t.Errorf("hexToAnsiBg() = %v, want %v", got, tt.want)
			}
		})
	}
}
