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

func NewRaythm() *ui.Page {
	var title = ui.NewUpper()

	var (
		activeFrame = 5
		knocked     = false
	)

	var anim = ui.NewAnim(
		ui.AnimOpts.NewImageFromBytes(assets.DefaultAnimSheetBytes),
		ui.AnimOpts.Size(32, 32),
		ui.AnimOpts.FPS(6),
		ui.AnimOpts.Loop(true),
		ui.AnimOpts.AutoPlay(true),
		ui.AnimOpts.OnFrame(activeFrame, func() {
			game.Tick()
		}),
		ui.AnchorOpts.OnEnd(func() {
			if !knocked {
				title.NewText(i18n.Miss)
			}

			knocked = false
		}),
	)

	var playSound = func() {
		if knocked {
			return
		}

		var s, _ = vorbis.DecodeF32(bytes.NewReader(assets.DefaultSoundBytes))
		var ply, _ = util.AudioContext.NewPlayerF32(s)

		var d = anim.FrameIndex() - activeFrame
		if d < 0 {
			d = -d
		}

		if d == 0 {
			title.NewText(i18n.Perfect)
		} else if d == 1 {
			title.NewText(i18n.Good)
		} else {
			title.NewText(i18n.Miss)
		}

		ply.Play()
		knocked = true
	}

	var helpExit = ui.BottomLeft(
		ui.NewHBox(
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
		),
	)

	var helpEnter = ui.BottomRight(
		ui.NewHBox(
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
		),
	)

	var main = ui.Center(
		ui.NewVBox(
			ui.VBoxOpts.AlignItems(ui.AlignCenter),
			ui.VBoxOpts.Contents(
				title,
				anim,
			),
		),
	)

	var p = ui.NewSpace(ui.SpaceOpts.Space(4))

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.OnInput(func() bool {
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				playSound()
				return true
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
				app.Load(MainMenu())
				return true
			}

			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
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
			)),
		),
	)

	return page

}
