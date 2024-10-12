package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	FreeMode = "Free Mode"
	AutoMode = "Auto Mode"
	Quit     = "Quit"
)

var (
	SC = language.SimplifiedChinese
	EN = language.English

	lang = SC
)

func init() {
	message.SetString(SC, FreeMode, "自由模式")
	message.SetString(SC, AutoMode, "自动模式")
	message.SetString(SC, Quit, "退出游戏")
}

func T(s string, args ...any) string {
	var p = message.NewPrinter(lang)
	return p.Sprintf(s, args...)
}
