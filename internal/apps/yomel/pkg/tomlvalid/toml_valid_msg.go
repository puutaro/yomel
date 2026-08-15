package tomlvalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/color"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/toml"
)

const (
	tomlSuffix                        = "\ntoml: %s"
	tomlColorStrWithQuote             = "'" + toml.ColorColor + "'"
	tomlBgColorStrWithQuote           = "'" + toml.ColorBgColor + "'"
	tomlCommentColorStrWithQuote      = "'" + toml.ColorCommentColor + "'"
	tomlTitleColorStrWithQuote        = "'" + toml.ColorTitleColor + "'"
	tomlTitleBgColorStrWithQuote      = "'" + toml.ColorTitleBgColor + "'"
	tomlTitleCommentColorStrWithQuote = "'" + toml.ColorTitleCommentColor + "'"
	tomlColorBgColorWithAnd           = tomlColorStrWithQuote + " and " + tomlBgColorStrWithQuote +
		" and " + tomlCommentColorStrWithQuote
	tomlCtrlColorBgColorWithAnd      = tomlColorStrWithQuote + " and " + tomlBgColorStrWithQuote
	tomlCtrlTitleColorBgColorWihtAnd = tomlTitleColorStrWithQuote + " and " + tomlTitleBgColorStrWithQuote + " and " +
		tomlTitleCommentColorStrWithQuote

	colorStrErrMsgBody = " color str is Irregular"
	// colorStrIrregularErrMsgForCtrl      = tomlCtrlColorBgColorWithAnd + colorStrErrMsgBody + tomlSuffix + "\n%s"
	colorStrIrregularErrMsg             = tomlColorBgColorWithAnd + colorStrErrMsgBody + tomlSuffix + "\n%s"
	orderOperatorCompMsgForTitleColorBg = color.OrderOperator + "'" + " cannot work in " + tomlCtrlTitleColorBgColorWihtAnd
	titleColorStrIrregularErrMsg        = tomlCtrlTitleColorBgColorWihtAnd + colorStrErrMsgBody + "\n" + orderOperatorCompMsgForTitleColorBg + tomlSuffix + "\n%s"
)
