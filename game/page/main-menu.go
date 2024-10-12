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
		ui.MenuOpts.TextItem(i18n.T(i18n.RhythmMode), func() {
			app.Load(NewRaythm())
		}),
		ui.MenuOpts.TextItem(i18n.T(i18n.Settings), func() {
			app.Load(NewSettings())
		}),
		// ui.MenuOpts.TextItem(i18n.T(i18n.Quit), func() {
		// 	app.Quit()
		// }),
	)

	var version = NewVersionInfo()

	var title = ui.NewText(
		ui.TextOpts.Color(app.Theme.SecondaryColor),
		ui.TextOpts.Content(i18n.T(i18n.APP_NAME)),
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
			ui.Layers(
				ui.Center(box),
				version,
			),
		),
	)

	return page
}

func NewVersionInfo() *ui.Anchor {
	return ui.BottomLeft(ui.NewText(
		ui.TextOpts.Color(app.Theme.SecondaryColor),
		ui.TextOpts.Content("v0.5.0"),
		ui.TextOpts.Padding(4),
	))
}
