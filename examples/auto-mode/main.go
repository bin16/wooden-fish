package main

import (
	_ "embed"
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/assets"
	"github.com/bin16/wooden-fish/game/page"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func main() {
	var game = app.New(
		app.Options.OnInput(func() bool {

			if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
				app.Quit()
				return true
			}

			return false
		}),
	)

	game.Load(page.NewAutoMode())

	ebiten.SetWindowPosition(200, 200)
	ebiten.SetWindowIcon(
		[]image.Image{
			assets.NewImage(assets.Icon32x32),
			assets.NewImage(assets.Icon16x16),
			assets.NewImage(assets.Icon8x8),
		},
	)
	ebiten.SetWindowTitle("功德+1")
	ebiten.SetWindowSize(240, 240)
	// ebiten.SetWindowDecorated(false)
	ebiten.RunGameWithOptions(game, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	})
}
