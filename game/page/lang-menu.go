package page

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/text/language"
)

func NewLangMenu() *ui.Page {
	var menu = ui.NewMenu(
		ui.MenuOpts.TextItem("中文", func() {
			game.Game.Language = language.SimplifiedChinese
			game.Save()
			app.Load(NewSettings())
		}),
		ui.MenuOpts.TextItem("English", func() {
			game.Game.Language = language.English
			game.Save()
			app.Load(NewSettings())
		}),
		ui.MenuOpts.OnExit(func() {
			app.Load(NewSettings())
		}),
	)

	var title = ui.NewText(
		ui.TextOpts.Content(i18n.T(i18n.Language)),
		ui.TextOpts.Color(app.Theme.SecondaryColor),
		ui.TextOpts.Padding(0, 0, 8, 0),
	)

	var box = ui.NewVBox(
		ui.VBoxOpts.AlignItems(ui.AlignCenter),
		ui.VBoxOpts.Contents(
			title,
			menu,
		),
	)

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.Contents(
			ui.Center(box),
		),
		ui.PageOpts.OnInput(func() bool {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
				app.Load(MainMenu())
				return true
			}

			return false
		}),
	)

	return page
}
