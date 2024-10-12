package main

import (
	"fmt"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	var menu = ui.NewMenu(
		ui.MenuOpts.TextItem("Start", func() {
			fmt.Println("> Start")
		}),
		ui.MenuOpts.TextItem("Continue", func() {
			fmt.Println("> Continue")
		}),
		ui.MenuOpts.TextItem("Settings", func() {
			fmt.Println("> Settings")
		}),
		ui.MenuOpts.TextItem("Quit", func() {
			fmt.Println("> Quit")
		}),
		ui.MenuOpts.OnExit(func() {
			fmt.Println("< Exit")
		}),
	)

	var box = ui.NewBox(
		ui.BoxOpts.Contents(menu),
	)

	var page = ui.NewPage(
		ui.PageOpts.Fill(app.Theme.BackgroundColor),
		ui.PageOpts.Contents(
			box,
		),
	)

	app.Load(page)

	ebiten.SetWindowPosition(200, 200)
	ebiten.SetWindowTitle("menu demo")
	ebiten.RunGame(app.Get())
}
