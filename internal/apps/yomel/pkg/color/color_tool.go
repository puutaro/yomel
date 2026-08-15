package color

import (
	"math/rand/v2"
	"strings"
)

const (
	darkBlueCode   = "darkBlue"
	blueCode       = "blue"
	lightBlueCode  = "lightBlue"
	darkRedCode    = "darkRed"
	redCode        = "red"
	lightRedCode   = "lightRed"
	darkGreenCode  = "darkGreen"
	greenCode      = "green"
	lightGreenCode = "lightGreen"
	darkAzureCode  = "darkAzure"
	azureCode      = "azure"
	lightAzureCode = "lightAzure"
	darkBrownCode  = "darkBrown"
	brownCode      = "brown"
	lightBrownCode = "lightBrown"
	yellowCode     = "yellow"
	black          = "black"
	white          = "white"
	darkGrayCode   = "darkGray"
	gray           = "gray"
	lightGrayCode  = "lightGray"
)
const (
	LightRnd      = "light"
	DarkRnd       = "dark"
	CenterRnd     = "center"
	Rnd           = "rnd"
	AndOperator   = "&"
	OrderOperator = ">"
)

func ConvertStrListToColorStr(strSrc string) string {
	if strSrc == "" {
		return ""
	}
	strList := strings.Split(strSrc, AndOperator)
	strListRndIndex := rand.IntN(len(strList))
	str := strings.TrimSpace(strList[strListRndIndex])
	var colorStrs []string
	switch str {
	case LightRnd:
		colorStrs = []string{
			lightAzureCode,
			lightBlueCode,
			lightBrownCode,
			lightGrayCode,
			lightGreenCode,
			lightRedCode,
			yellowCode,
		}
	case DarkRnd:
		colorStrs = []string{
			darkAzureCode,
			darkBlueCode,
			darkBrownCode,
			darkGrayCode,
			darkGreenCode,
			darkRedCode,
		}
	case CenterRnd:
		colorStrs = []string{
			azureCode,
			blueCode,
			brownCode,
			gray,
			redCode,
		}
	case Rnd:
		colorStrs = []string{
			lightAzureCode,
			lightBlueCode,
			lightBrownCode,
			lightGrayCode,
			lightGreenCode,
			lightRedCode,
			darkAzureCode,
			darkBlueCode,
			darkBrownCode,
			darkGrayCode,
			darkGreenCode,
			darkRedCode,
			azureCode,
			blueCode,
			brownCode,
			gray,
			redCode,
		}
	default:
		return str
	}
	randomIndex := rand.IntN(len(colorStrs))
	return colorStrs[randomIndex]
}
func ConvertFixColorStr(strSrc string) string {
	str := ConvertStrListToColorStr(strSrc)
	switch str {
	case darkBlueCode:
		return "#001d9e"
	case blueCode:
		return "#0026ff"
	case lightBlueCode:
		return "#26c0fc"
	case darkRedCode:
		return "#850101"
	case redCode:
		return "#ff0000"
	case lightRedCode:
		return "#fa4d4d"
	case darkGreenCode:
		return "#014702"
	case greenCode:
		return "#11b812"
	case lightGreenCode:
		return "#43fa46"
	case yellowCode:
		return "#fff200"
	case darkAzureCode:
		return "#1c4d44"
	case azureCode:
		return "#21ebc6"
	case lightAzureCode:
		return "#67e8eb"
	case darkBrownCode:
		return "#572b07"
	case brownCode:
		return "#b35c15"
	case lightBrownCode:
		return "#f26e27"
	case black:
		return "#000000"
	case white:
		return "#ffffff"
	case darkGrayCode:
		return "#424242"
	case gray:
		return "#808080"
	case lightGrayCode:
		return "#dbdbdb"
	default:
		return strSrc
	}
}
