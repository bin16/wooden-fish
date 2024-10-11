package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/png"

	"github.com/bin16/go-hexcolor"
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var audioContext = audio.NewContext(48000)

func main() {

	var text = ui.NewText(
		ui.TextOpts.Content("\n\n功德+1"),
		ui.TextOpts.Color(hexcolor.New("#da0")),
	)

	var anim = ui.NewAnim(
		ui.AnimOpts.NewImageFromBytes(sheet_bytes),
		ui.AnimOpts.Size(32, 32),
		ui.AnimOpts.OnFrame(5, func() {
			fmt.Println("d5")

			var s, _ = vorbis.DecodeF32(bytes.NewReader(sound_bytes))
			var ply, _ = audioContext.NewPlayerF32(s)

			ply.Play()
		}),
	)

	var box = ui.NewBox()
	box.AddChild(text)
	box.AddChild(anim)

	var game = app.New(
		app.Options.OnInput(func() bool {

			if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
				app.Quit()
				return true
			}

			return false
		}),
	)

	game.Load(box)

	ebiten.SetWindowPosition(200, 200)
	ebiten.SetWindowIcon(
		[]image.Image{
			icon_img,
		},
	)
	ebiten.SetWindowTitle("功德+1")
	ebiten.SetWindowSize(160, 160)
	ebiten.SetWindowDecorated(false)
	ebiten.RunGame(game)
}

//go:embed "data/Theme/default/sound.ogg"
var sound_bytes []byte

//go:embed "data/Theme/default/wooden-fish-32x32-anim.png"
var sheet_bytes []byte

//go:embed "icon-32x32.png"
var icon_bytes []byte

var icon_png, _ = png.Decode(bytes.NewReader(icon_bytes))
var icon_img = ebiten.NewImageFromImage(icon_png)
