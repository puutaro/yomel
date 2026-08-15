package color

import (
	"fmt"
	"strconv"
	"strings"
)

func DetectColorStrIrregularErrForCtrl(
	colorStr string,
	colorStrIrregularErrMsgForCtrl string,
) error {
	err := DetectColorCodeStrIrregularErr(colorStr)
	if err != nil {
		return fmt.Errorf(
			colorStrIrregularErrMsgForCtrl,
			0,
			err,
		)
	}
	return nil
}

func DetectColorCodeStrIrregularErr(str string) error {
	if str == "" {
		return nil
	}
	strList := strings.Split(
		str,
		AndOperator,
	)
	for _, s := range strList {
		if s == "" {
			continue
		}
		trimed := strings.TrimSpace(s)
		hexStr := ConvertFixColorStr(trimed)
		// 先頭の '#' があればトリムする
		hex := strings.TrimPrefix(hexStr, "#")
		// 桁数のチェック（3桁または6桁に対応する場合もあるが、基本は6桁を想定）
		if len(hex) != 6 {
			return fmt.Errorf("invalid hex color length or macro: %s", hex)
		}
		// 16進数文字列をパース
		_, err := strconv.ParseUint(hex, 16, 32)
		if err != nil {
			return err
		}
	}
	return nil
}
