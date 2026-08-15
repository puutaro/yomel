package modelvalid

import (
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/argtables"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/color"
)

const (
	stageNoSuffix                 = "\nstageNo: %d"
	curTitle                      = "\ntitle: '%s'"
	stageDesc                     = "\ndesc: '%s'"
	titleWithQuote                = "'" + argtables.TitleSignal + "'"
	colorStrWithQuote             = "'" + argtables.ColorOpSignal + "'"
	bgColorStrWithQuote           = "'" + argtables.BgColorOpSignal + "'"
	commentColorStrWithQuote      = "'" + argtables.CommentColorOpSignal + "'"
	titleColorStrWithQuote        = "'" + argtables.TitleColorOpSignal + "'"
	titleBgColorStrWithQuote      = "'" + argtables.TitleBgColorOpSignal + "'"
	titleCommentColorStrWithQuote = "'" + argtables.TitleCommentColorOpSignal + "'"
	colorBgColorWithAnd           = colorStrWithQuote + " and " + bgColorStrWithQuote +
		" and " + commentColorStrWithQuote
	ctrlColorBgColorWithAnd      = colorStrWithQuote + " and " + bgColorStrWithQuote
	ctrlTitleColorBgColorWihtAnd = titleColorStrWithQuote + " and " + titleBgColorStrWithQuote + " and " +
		titleCommentColorStrWithQuote
	stageSignalWithQuote  = "'" + argtables.StageSignal + "'"
	optOpSignalWithQuote  = "'" + argtables.OptOpSignal + "'"
	lOptOpSignalWithQuote = "'" + argtables.LoptOpSignal + "'"
	opLopWithAnd          = optOpSignalWithQuote + " and " + lOptOpSignalWithQuote

	titleDescriptionIrregular = titleWithQuote + " must be meaning sentence if stage > 1" + curTitle
	stageDescriptionIrregular = stageSignalWithQuote + " description must be meaning sentence" + stageNoSuffix + stageDesc
	stageDescriptionDuplicate = stageSignalWithQuote + " description must be unique across stages" + stageNoSuffix + stageDesc

	noBlankStrRequireErrMsg = "'%s' no blank str is required" + stageNoSuffix
	optStrBlankErrMsg       = opLopWithAnd + " str must not be blank" + stageNoSuffix

	colorStrErrMsgBody                  = " color str is Irregular"
	colorStrIrregularErrMsgForCtrl      = ctrlColorBgColorWithAnd + colorStrErrMsgBody + stageNoSuffix + "\n%s"
	orderOperatorCompMsg                = color.OrderOperator + "'" + " cannot work only in ctrl stage[0]"
	colorStrIrregularErrMsg             = colorBgColorWithAnd + colorStrErrMsgBody + "\n" + orderOperatorCompMsg + stageNoSuffix + "\n%s"
	orderOperatorCompMsgForTitleColorBg = color.OrderOperator + "'" + " cannot work in these options"
	titleColorStrIrregularErrMsg        = ctrlTitleColorBgColorWihtAnd + colorStrErrMsgBody + "\n" + orderOperatorCompMsgForTitleColorBg + stageNoSuffix + "\n%s"

	// cmdStrRequire        = "'" + argtables.CmdOpSignal + "'" + noBlankStrRequireErrMsg
	// svcStrRequire        = "'" + argtables.SvcOpSignal + "'" + noBlankStrRequireErrMsg
	// actStrRequire        = "'" + argtables.ActOpSignal + "'" + noBlankStrRequireErrMsg
)
