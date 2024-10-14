package game

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game/save"
	"golang.org/x/text/language"
)

var Version string = "v0.0.0"

type game struct {
	Count    int          `json:"count"`
	Language language.Tag `json:"lang"`
	ThemeID  string       `json:"theme"`
}

var Game = game{
	Count:    0,
	Language: language.English,
	ThemeID:  "",
}

func init() {
	save.Read("game.save.json", &Game)
	app.LoadTheme(Game.ThemeID)
}

func Save() {
	save.Write("game.save.json", &Game)
}

func Tick() {
	Game.Count += 1
	Save()
}
