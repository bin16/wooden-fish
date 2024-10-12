package page

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/ui"
)

func MainMenu() *ui.Page {
	var menu = ui.NewMenu(
		ui.MenuOpts.TextItem(i18n.T(i18n.AutoMode), func() {
			app.Load(NewAutoMode())
		}),
		ui.MenuOpts.TextItem(i18n.T(i18n.FreeMode), func() {
			app.Load(NewFreeMode())
		}),
		ui.MenuOpts.TextItem(i18n.T(i18n.Quit), func() {
			app.Quit()
		}),
	)

	var box = ui.NewBox(
		ui.BoxOpts.Contents(menu),
	)

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.Contents(
			box,
		),
	)

	return page
}
