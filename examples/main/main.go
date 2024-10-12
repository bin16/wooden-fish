package main

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/game/page"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	app.Load(page.MainMenu())

	ebiten.SetWindowPosition(200, 200)
	ebiten.SetWindowTitle("demo")
	ebiten.RunGame(app.Get())
}
