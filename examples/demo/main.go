package main

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game/page"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	var page = page.NewHead(&page.HeadPageOptions{})

	app.Load(page)

	ebiten.SetWindowPosition(200, 200)
	ebiten.SetWindowTitle("menu demo")
	ebiten.RunGame(app.Get())
}
