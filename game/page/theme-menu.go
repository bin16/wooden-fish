package page

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func NewThemeMenu() *ui.Page {
	var menu = ui.NewMenu(
		ui.MenuOpts.OnExit(func() {
			app.Load(NewSettings())
		}),
	)

	var p = ui.NewSpace(ui.SpaceOpts.Space(2, 4))

	var b = ui.NewBorder(
		ui.BorderOpts.Border(1),
		ui.BorderOpts.Color(app.Theme.SecondaryColor),
		ui.BorderOpts.BorderRadius(3),
	)

	for i, th := range app.ThemeOptions {
		var pal = ui.NewPalette(
			ui.PaletteOpts.Width(16),
			ui.PaletteOpts.Height(16),
			ui.PaletteOpts.BorderRadius(0),
			ui.PaletteOpts.BorderColor(th.BackgroundColor),
			ui.PaletteOpts.Colors(
				th.BackgroundColor,
				th.Color,
				th.SecondaryColor,
				th.AccentColor,
			),
		)

		var text = ui.NewText(
			ui.TextOpts.Content(i18n.T(th.Name)),
			ui.TextOpts.Padding(0, 0, 0, 2),
		)

		var item = ui.NewHBox(
			ui.HBoxOpts.AlignItems(ui.AlignCenter),
			ui.HBoxOpts.Contents(
				pal,
				text,
			),
		)

		ui.MenuOpts.Item(p(item), func() {
			app.LoadTheme(th)
			game.Game.ThemeID = th.ID
			game.Save()
			app.Load(NewThemeMenu())
		})(menu)

		if th.ID == app.Theme.ID {
			menu.HandleFocus(i)
		}
	}
	// ui.MenuOpts.TextItem(i18n.T(i18n.Back), func() {
	// 	app.Load(NewSettings())
	// }, ui.TextOpts.Color(app.Theme.SecondaryColor))(menu)

	var title = ui.NewText(
		ui.TextOpts.Content(i18n.T(i18n.Theme)),
		ui.TextOpts.Color(app.Theme.SecondaryColor),
		ui.TextOpts.Padding(0, 0, 8, 0),
	)

	var box = ui.NewVBox(
		ui.VBoxOpts.AlignItems(ui.AlignCenter),
		ui.VBoxOpts.Contents(
			title,
			b(p(menu)),
		),
	)

	var helpExit = NewBack(func(data ...any) bool {
		app.Load(MainMenu())
		return true
	})

	var nav = ui.BottomRight(
		ui.NewHBox(
			ui.HBoxOpts.Contents(
				ui.NewText(ui.TextOpts.Content("[")),
				ui.OnTap(ui.NewText(
					ui.TextOpts.Content(i18n.T(i18n.PrevItem)),
					ui.TextOpts.Color(app.Theme.SecondaryColor),
				), func(data ...any) bool {
					menu.FocusUp()
					return true
				}),
				ui.NewText(ui.TextOpts.Content("|")),
				ui.OnTap(ui.NewText(
					ui.TextOpts.Content(i18n.T(i18n.NextItem)),
					ui.TextOpts.Color(app.Theme.SecondaryColor),
				), func(data ...any) bool {
					menu.FocusDown()
					return true
				}),
				ui.NewText(ui.TextOpts.Content("]")),
			),
		),
	)

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.Contents(
			p(ui.Layers(
				ui.Center(box),
				helpExit,
				nav,
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
