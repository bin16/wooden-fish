package page

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/assets"
	"github.com/bin16/wooden-fish/game"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/ui"
	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func NewRaythm() *ui.Page {
	var title = ui.NewUpper()

	var (
		// knocked = false
		// knocked_flag = make([]bool, len(game.Animation.Sounds))
		knocked_frame_index = -100
		checked             = map[int]bool{}
	)

	var anim = ui.NewAnim(
		ui.AnimOpts.AutoPlay(true),
		ui.AnimOpts.Loop(true),
		ui.AnimOpts.Image(game.Animation.Image()),
		ui.AnimOpts.Size(game.Animation.Size()),
		ui.AnimOpts.FPS(util.NotZero(
			game.Animation.RhythmMode.FPS,
			game.Animation.FPS,
			9,
		)),
	)
	for _, sound := range game.Animation.Sounds {
		ui.AnimOpts.OnFrame(sound.FrameIndex, func() {
			game.Tick()
		})(anim)
	}
	for _, sound := range game.Animation.Sounds {
		var k = sound.FrameIndex
		ui.AnimOpts.OnFrame(k+2, func() {
			if !checked[k] {
				title.NewText(i18n.Miss)
				checked[k] = true
			}
		})(anim)
	}
	ui.AnchorOpts.OnEnd(func() {
		var lastCheckPoint = -1
		for _, s := range game.Animation.Sounds {
			lastCheckPoint = max(s.FrameIndex)
		}

		if !checked[lastCheckPoint] {
			title.NewText(i18n.T(i18n.Miss))
		}

		knocked_frame_index = -999
		checked = make(map[int]bool)
	})(anim)

	// var anim = ui.NewAnim(
	// 	ui.AnimOpts.NewImageFromBytes(assets.DefaultAnimSheetBytes),
	// 	ui.AnimOpts.Size(48, 48),
	// 	ui.AnimOpts.FPS(6),
	// 	ui.AnimOpts.Loop(true),
	// 	ui.AnimOpts.AutoPlay(true),
	// 	ui.AnimOpts.OnFrame(activeFrame, func() {
	// 		game.Tick()
	// 	}),
	// 	ui.AnchorOpts.OnEnd(func() {
	// 		if !knocked {
	// 			title.NewText(i18n.Miss)
	// 		}

	// 		knocked = false
	// 	}),
	// )

	var playSound = func(name string) {
		var ply = assets.NewAudioPlayer(name)

		ply.Play()
	}

	var handleBeat = func() {
		var k = anim.FrameIndex()
		if k-knocked_frame_index < 2 {
			return
		}

		for _, s := range game.Animation.Sounds {
			var k1 = s.FrameIndex
			var d = util.OR(k > k1, k-k1, k1-k)

			if d < 1 {
				checked[k1] = true
				title.NewText(i18n.T(i18n.Perfect))
				playSound(s.Source)
				return
			}

			if d < 2 {
				checked[k1] = true
				title.NewText(i18n.T(i18n.Good))
				playSound(s.Source)
				return
			}
		}

		title.NewText(i18n.T(i18n.Miss))
	}

	var helpExit = NewBack(func(data ...any) bool {
		app.Load(MainMenu())
		return true
	})

	var helpEnter = NewEnter(func(data ...any) bool {
		handleBeat()
		return true
	})

	var main = ui.Center(
		ui.OnTap(ui.NewVBox(
			ui.VBoxOpts.AlignItems(ui.AlignCenter),
			ui.VBoxOpts.Contents(
				title,
				anim,
			),
		), func(data ...any) bool {
			handleBeat()
			return true
		}),
	)

	var p = ui.NewSpace(ui.SpaceOpts.Space(4))

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.OnInput(func() bool {
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				handleBeat()
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
			)),
		),
	)

	return page

}
