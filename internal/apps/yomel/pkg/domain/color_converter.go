package domain

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/color"
)

func hexToAnsiFg(hex string) string {
	if hex == "" {
		return ""
	}
	return hexToAnsi(hex, ForegroundAnsi)
}
func hexToAnsiBg(hex string) string {
	if hex == "" {
		return ""
	}
	return hexToAnsi(hex, BackgroundAnsi)
}
func hexToAnsiFgForStage(hexSrc string, ctrlHex string) string {
	var hex string
	if hexSrc == "" {
		hex = color.ConvertFixColorStr(ctrlHex)
	} else {
		hex = hexSrc
	}
	return hexToAnsi(hex, ForegroundAnsi)
}
func hexToAnsiBgForStage(hexSrc string, ctrlHex string) string {
	var hex string
	if hexSrc == "" {
		hex = color.ConvertFixColorStr(ctrlHex)
	} else {
		hex = hexSrc
	}

	return hexToAnsi(hex, BackgroundAnsi)
}
func hexToAnsi(hex string, fOrB ForeOrBack) string {
	if hex == "" {
		return ""
	}
	r, g, b := parseHexColor(hex)
	// 38;2;R;G;B が 24bit文字色指定のANSIコード
	ansiForeOrBackCode := ForegroundColorAnsiCode
	switch fOrB {
	case BackgroundAnsi:
		ansiForeOrBackCode = BackgroundColorAnsiCode
	case ForegroundAnsi:
		ansiForeOrBackCode = ForegroundColorAnsiCode
	}
	return fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", ansiForeOrBackCode, r, g, b)
}

func parseHexColor(hex string) (int, int, int) {
	// 先頭の '#' があればトリムする
	hexStr := color.ConvertFixColorStr(hex)
	hex = strings.TrimPrefix(hexStr, "#")
	// 16進数文字列をパース
	rgb, _ := strconv.ParseUint(hex, 16, 32)
	r := int((rgb >> 16) & 0xFF)
	g := int((rgb >> 8) & 0xFF)
	b := int(rgb & 0xFF)
	return r, g, b
}
