package page

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var p = ui.NewSpace(ui.SpaceOpts.Space(2, 4))

func NewAnimMenu() *ui.Page {

	var b = ui.NewBorder(
		ui.BorderOpts.Border(1),
		ui.BorderOpts.Color(app.Theme.SecondaryColor),
		ui.BorderOpts.BorderRadius(3),
	)

	var menu = ui.NewMenu(
		ui.MenuOpts.OnExit(func() {
			app.Load(NewSettings())
		}),
	)
	for i, anim := range game.AnimationOptions {
		var o = ui.MenuOpts.TextItem(i18n.T(anim.Name), func() {
			game.LoadAnim(anim)
			game.Save()
			app.Load(MainMenu())
		})
		o(menu)

		if anim.ID == game.Animation.ID {
			menu.HandleFocus(i)
		}
	}
	// ui.MenuOpts.TextItem(i18n.T(i18n.Back), func() {
	// 	app.Load(NewSettings())
	// }, ui.TextOpts.Color(app.Theme.SecondaryColor))(menu)

	var title = ui.NewText(
		ui.TextOpts.Content(i18n.T(i18n.Animation)),
		ui.TextOpts.Color(app.Theme.SecondaryColor),
		ui.TextOpts.Padding(0, 0, 8, 0),
	)

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

	var box = ui.NewVBox(
		ui.VBoxOpts.AlignItems(ui.AlignCenter),
		ui.VBoxOpts.Contents(
			title,
			b(p(menu)),
		),
	)

	var back = NewBack(func(data ...any) bool {
		app.Load(NewSettings())
		return true
	})

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.Contents(
			p(ui.Layers(
				ui.Center(box),
				back,
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
