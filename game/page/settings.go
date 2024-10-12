package page

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/ui"
)

func NewSettings() *ui.Page {

	var menu = ui.NewMenu(
		ui.MenuOpts.TextItem(i18n.T(i18n.Language), func() {
			app.Load(NewLangMenu())
		}),
		ui.MenuOpts.TextItem(i18n.T(i18n.Back), func() {
			app.Load(MainMenu())
		}),
		ui.MenuOpts.OnExit(func() {
			app.Load(MainMenu())
		}),
	)

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

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.Contents(
			ui.Center(box),
		),
	)

	return page
}
