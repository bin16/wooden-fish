package page

import (
	"bytes"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/assets"
	"github.com/bin16/wooden-fish/game"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/ui"
	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func NewFreeMode() app.Scene {
	var title = ui.NewUpper()

	var anim = ui.NewAnim(
		ui.AnimOpts.NewImageFromBytes(assets.DefaultAnimSheetBytes),
		ui.AnimOpts.Size(48, 48),
		ui.AnimOpts.FPS(30),
		ui.AnimOpts.OnFrame(5, func() {
			var s, _ = vorbis.DecodeF32(bytes.NewReader(assets.DefaultSoundBytes))
			var ply, _ = util.AudioContext.NewPlayerF32(s)

			ply.Play()
			title.NewText(i18n.T(i18n.MeritPlusOne))

			game.Tick()
		}),
	)

	var main = ui.Center(
		ui.OnTap(
			ui.NewVBox(
				ui.VBoxOpts.AlignItems(ui.AlignCenter),
				ui.VBoxOpts.Contents(
					title,
					anim,
				),
			), func(data ...any) bool {
				anim.Play()
				return true
			},
		),
	)

	var statInfo = ui.Top(
		ui.NewText(
			ui.TextOpts.Pull(func() string {
				return i18n.T(i18n.Merits_Is, game.Game.Count)
			}),
		),
	)

	var helpExit = NewBack(func(data ...any) bool {
		app.Load(MainMenu())
		return true
	})

	var helpEnter = NewEnter(func(data ...any) bool {
		anim.Play()
		return true
	})

	var p = ui.NewSpace(ui.SpaceOpts.Space(4))

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.OnInput(func() bool {
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				anim.Play()
				return true
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
				app.Load(MainMenu())
				return true
			}

			return false
		}),
		ui.PageOpts.Contents(
			p(ui.Layers(
				main,
				helpExit,
				helpEnter,
				statInfo,
			)),
		),
	)

	return page
}

func NewBack(onTap app.EventFunc) *ui.Anchor {
	var content = ui.NewHBox(
		ui.HBoxOpts.Contents(
			ui.NewText(
				ui.TextOpts.Content("[Esc]"),
				ui.TextOpts.Color(app.Theme.SecondaryColor),
			),
			ui.NewText(
				ui.TextOpts.Content(i18n.T(i18n.Back)),
				ui.TextOpts.Color(app.Theme.SecondaryColor),
			),
		),
	)

	return ui.BottomLeft(
		ui.OnTap(content, onTap),
	)
}

func NewEnter(onTap app.EventFunc) *ui.Anchor {
	var main = ui.NewHBox(
		ui.HBoxOpts.JustifyContent(ui.JustifyEnd),
		ui.HBoxOpts.Contents(
			ui.NewText(
				ui.TextOpts.Content("[Enter]"),
				ui.TextOpts.Color(app.Theme.SecondaryColor),
			),
			ui.NewText(
				ui.TextOpts.Content(i18n.T(i18n.Beat)),
				ui.TextOpts.Color(app.Theme.SecondaryColor),
			),
		),
	)

	return ui.BottomRight(
		ui.OnTap(main, onTap),
	)
}
