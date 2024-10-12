package i18n

import (
	"github.com/bin16/wooden-fish/game"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	FreeMode   = "Free Mode"
	AutoMode   = "Auto Mode"
	RhythmMode = "Rhythm Mode"

	Quit         = "Quit"
	Merit        = "Merit"
	MeritPlusOne = "Merit +1"
	Merits_Is    = "Merits: %d"
	Back         = "Back"
	Next         = "Next"
	Exit         = "Exit"

	Settings = "Settings"
	Language = "Language"

	Perfect = "Perfect"
	Good    = "Good"
	Miss    = "Miss"

	Knock = "Knock"

	WoodenFish = "Wooden Fish"

	APP_NAME = "__app_name__"
)

var (
	SC = language.SimplifiedChinese
	EN = language.English
)

func init() {
	message.SetString(SC, APP_NAME, "敲木鱼")
	message.SetString(EN, APP_NAME, "Wooden Fish")

	message.SetString(SC, Language, "语言")
	message.SetString(SC, Settings, "设置")
	message.SetString(SC, RhythmMode, "音游模式")
	message.SetString(SC, FreeMode, "自由模式")
	message.SetString(SC, AutoMode, "自动模式")
	message.SetString(SC, Quit, "退出游戏")
	message.SetString(SC, Merit, "功德")
	message.SetString(SC, MeritPlusOne, "功德+1")
	message.SetString(EN, Merits_Is, "Merits: %d")
	message.SetString(SC, Merits_Is, "功德：%d")
	message.SetString(SC, Exit, "退出")
	message.SetString(SC, Back, "返回")
	message.SetString(SC, Knock, "敲")
	message.SetString(SC, WoodenFish, "木鱼")
}

func T(s string, args ...any) string {
	var p = message.NewPrinter(game.Game.Language)
	return p.Sprintf(s, args...)
}
