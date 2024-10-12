package i18n

import (
	"github.com/bin16/wooden-fish/game"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	FreeMode     = "Free Mode"
	AutoMode     = "Auto Mode"
	Quit         = "Quit"
	Merit        = "Merit"
	MeritPlusOne = "Merit +1"
	Merits       = "Merits: %d"
)

var (
	SC = language.SimplifiedChinese
	EN = language.English
)

func init() {
	message.SetString(SC, FreeMode, "自由模式")
	message.SetString(SC, AutoMode, "自动模式")
	message.SetString(SC, Quit, "退出游戏")
	message.SetString(SC, Merit, "功德")
	message.SetString(SC, MeritPlusOne, "功德+1")
	message.Set(EN, Merits,
		plural.Selectf(1, "%d",
			1, "Merits",
			">1", "Merits",
		),
	)
	message.SetString(SC, Merits, "功德")
}

func T(s string, args ...any) string {
	var p = message.NewPrinter(game.Game.Language)
	return p.Sprintf(s, args...)
}
