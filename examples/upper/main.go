package main

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func main() {
	var up = ui.NewUpper()
	up.NewText("Hello?")

	var main = ui.NewPage(
		ui.PageOpts.Contents(
			ui.Center(
				up,
			),
		),
		ui.PageOpts.OnInput(func() bool {
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				up.NewText("Hello?")

				return true
			}

			return false
		}),
	)

	app.Load(
		main,
	)

	ebiten.SetWindowPosition(200, 200)
	ebiten.SetWindowTitle("menu demo")
	ebiten.RunGame(app.Get())
}
