package main

import (
	"fmt"
	"image/color"

	"github.com/bin16/go-hexcolor"
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	var vbox = ui.NewVBox()

	for i := 0; i < 12; i++ {
		var colors = []color.Color{
			hexcolor.New("#c00"),
			hexcolor.New("#720"),
			hexcolor.New("#87f"),
			hexcolor.New("#039"),
			hexcolor.New("#9ed"),
		}
		var dummy = ui.NewDummy(
			ui.DummyOpts.Color(colors[i%(len(colors))]),
		)
		var s = fmt.Sprintf("Row %02d", i+1)
		var text = ui.NewText(
			ui.TextOpts.Content(s),
		)

		vbox.AddChild(dummy)
		vbox.AddChild(text)
	}

	app.Load(vbox)

	ebiten.SetWindowTitle("vbox demo")
	ebiten.SetWindowSize(800, 800)
	ebiten.RunGame(app.Get())
}
