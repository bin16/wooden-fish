package game

import (
	"log"
	"path"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game/save"
	"golang.org/x/text/language"
)

var Version string = "v0.0.0"

type game struct {
	Count    int          `json:"count"`
	Language language.Tag `json:"lang"`
	ThemeID  string       `json:"theme"`
	AnimID   string       `json:"anim"`
}

var Game = game{
	Count:    0,
	Language: language.English,
	ThemeID:  "id-theme-default",
	AnimID:   "id-anim-default",
}

func init() {
	names, err := save.Find("data/Theme", ".theme.json")
	if err != nil {
		log.Println(err)
	}

	for _, name := range names {
		var filename = path.Join("data/Theme", name)
		var td = &app.ThemeData{
			ID:   filename,
			Name: filename,
		}

		if err := save.ReadJSON(filename, td); err != nil {
			continue
		}

		var th = app.NewTheme(app.ThemeOpts.Import(td))
		app.RegisterTheme(th)
	}

	save.ReadSave("game.save.json", &Game)
	app.SetTheme(Game.ThemeID)
	SetAnim(Game.AnimID)
}

func Save() {
	save.WriteSave("game.save.json", &Game)
}

func Tick() {
	Game.Count += 1
	Save()
}
