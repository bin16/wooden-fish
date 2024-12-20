package page

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func NewSettings() *ui.Page {

	var menu = ui.NewMenu(
		ui.MenuOpts.TextItem(i18n.T(i18n.Language), func() {
			app.Load(NewLangMenu())
		}),
		ui.MenuOpts.TextItem(i18n.T(i18n.Theme), func() {
			app.Load(NewThemeMenu())
		}),
		ui.MenuOpts.TextItem(i18n.T(i18n.Animation), func() {
			app.Load(NewAnimMenu())
		}),
		ui.MenuOpts.OnExit(func() {
			app.Load(MainMenu())
		}),
	)
	// ui.MenuOpts.TextItem(i18n.T(i18n.Back), func() {
	// 	app.Load(MainMenu())
	// }, ui.TextOpts.Color(app.Theme.SecondaryColor))(menu)

	var title = ui.NewText(
		ui.TextOpts.Content(i18n.T(i18n.Settings)),
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

	var back = NewBack(func(data ...any) bool {
		app.Load(MainMenu())
		return true
	})

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.Contents(
			p(ui.Layers(
				ui.Center(box),
				back,
			)),
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
