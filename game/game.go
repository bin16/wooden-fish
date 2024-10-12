package game

import (
	"github.com/bin16/wooden-fish/game/save"
	"golang.org/x/text/language"
)

type game struct {
	Count    int          `json:"count"`
	Language language.Tag `json:"lang"`
}

var Game = game{
	Count:    0,
	Language: language.English,
}

func init() {
	save.Read("game.save.json", &Game)
}

func Save() {
	save.Write("game.save.json", &Game)
}

func Tick() {
	Game.Count += 1
	Save()
}
