package page

import (
	"bytes"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/assets"
	"github.com/bin16/wooden-fish/ui"
	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func NewFree() app.Scene {
	var title = ui.NewText(
		ui.TextOpts.Content("功德+1"),
	)

	var anim = ui.NewAnim(
		ui.AnimOpts.NewImageFromBytes(assets.DefaultAnimSheetBytes),
		ui.AnimOpts.Size(32, 32),
		ui.AnimOpts.FPS(30),
		ui.AnimOpts.OnFrame(5, func() {
			var s, _ = vorbis.DecodeF32(bytes.NewReader(assets.DefaultSoundBytes))
			var ply, _ = util.AudioContext.NewPlayerF32(s)

			ply.Play()
		}),
	)

	var box = ui.NewVBox(
		ui.VBoxOpts.Contents(
			title,
			anim,
		),
	)

	var main = ui.NewWindow(
		ui.WinOpts.OnInput(func() bool {
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				anim.Play()
				return true
			}

			return false
		}),
		ui.WinOpts.Contents(
			box,
		),
	)

	return main
}
