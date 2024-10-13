package main

//go:generate go-winres make --product-version=git-tag

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/assets"
	"github.com/bin16/wooden-fish/game/i18n"
	"github.com/bin16/wooden-fish/game/page"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	app.New(
		app.Options.UIScale(2.0),
	)

	app.Load(page.MainMenu())

	ebiten.SetWindowSize(320, 320)
	ebiten.SetWindowTitle(i18n.T(i18n.APP_NAME))
	ebiten.SetWindowIcon(
		[]image.Image{
			assets.Icon_32x32,
			assets.Icon_16x16,
			assets.Icon_8x8,
		},
	)

	ebiten.RunGame(app.Get())
}