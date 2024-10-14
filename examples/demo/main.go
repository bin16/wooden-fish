package main

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/assets"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/game/page"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	var page = page.NewHead(&page.HeadPageOptions{})

	app.Load(page)

	ebiten.SetWindowPosition(200, 200)
	ebiten.SetWindowTitle(i18n.T(i18n.APP_NAME))
	ebiten.SetWindowIcon(
		[]image.Image{
			assets.NewImage(assets.Icon32x32),
			assets.NewImage(assets.Icon16x16),
			assets.NewImage(assets.Icon8x8),
		},
	)
	ebiten.RunGame(app.Get())
}
