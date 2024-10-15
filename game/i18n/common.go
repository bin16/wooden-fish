package i18n

import (
	"github.com/bin16/wooden-fish/game"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type LangOptions struct {
	Name string
	Tag  language.Tag
}

var Options = []LangOptions{
	{"中文", SC},
	{"English", EN},
}

var (
	FreeMode   = "Free Mode"
	AutoMode   = "Auto Mode"
	RhythmMode = "Rhythm Mode"

	Quit         = "Quit"
	Merit        = "Merit"
	MeritPlusOne = "Merit+1"
	Merits_Is    = "Merits: %d"
	Back         = "Back"
	Next         = "Next"
	Exit         = "Exit"

	Settings  = "Settings"
	Language  = "Language"
	Theme     = "Theme"
	Animation = "Animation"

	Perfect = "Perfect"
	Good    = "Good"
	Miss    = "Miss"

	Beat = "Beat"

	WoodenFish = "Wooden Fish"

	APP_NAME = "__app_name__"

	PrevItem = "Previous"
	Previous = "Previous"
	NextItem = "Next"
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
	message.SetString(SC, Theme, "主题")
	message.SetString(SC, Animation, "动画")
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
	message.SetString(SC, Beat, "敲")
	message.SetString(SC, WoodenFish, "木鱼")

	message.SetString(SC, "Default Theme", "默认主题")
	message.SetString(SC, "High Contrast Theme", "高对比度主题")
	message.SetString(SC, "High Contrast Theme (Dark)", "高对比度主题（暗黑）")
	message.SetString(SC, "Default Animation", "默认动画（木鱼）")
	message.SetString(SC, "Draft Animation", "草图动画")
	message.SetString(SC, "Test Animation", "测试动画")

	message.SetString(SC, "Ding~", "叮~")
	message.SetString(SC, "Dong~", "咚~")
	message.SetString(SC, "Scores: %d", "分数：%d")

	message.SetString(SC, Previous, "上一个")
	message.SetString(SC, PrevItem, "上一个")
	message.SetString(SC, NextItem, "下一个")
}

func T(s string, args ...any) string {
	var p = message.NewPrinter(game.Game.Language)
	return p.Sprintf(s, args...)
}
