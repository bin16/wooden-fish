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

	type p struct {
		a ui.AlignItems
		j ui.JustifyContent
	}

	for _, n := range []p{
		{ui.AlignStart, ui.JustifyStart},
		{ui.AlignCenter, ui.JustifyCenter},
		{ui.AlignEnd, ui.JustifyEnd},
		{ui.AlignStretch, ui.JustifyCenter},
		{ui.AlignEnd, ui.SpaceBetween},
		{ui.AlignStretch, ui.SpaceBetween},
		{ui.AlignStart, ui.SpaceBetween},
	} {
		var hbox = ui.NewHBox(
			ui.HBoxOpts.AlignItems(n.a),
			ui.HBoxOpts.JustifyContent(n.j),
		)

		var m = ui.NewSpace(
			ui.SpaceOpts.Space(2),
		)

		var b = ui.NewBorder(
			ui.BorderOpts.Border(1),
			ui.BorderOpts.BorderRadius(2),
			ui.BorderOpts.Color(hexcolor.New("#0c0")),
		)

		var colors = []color.Color{
			hexcolor.New("#c00"),
			hexcolor.New("#720"),
			hexcolor.New("#87f"),
			hexcolor.New("#039"),
			hexcolor.New("#9ed"),
		}

		for i := 0; i < 5; i++ {
			var dummy = ui.NewDummy(
				ui.DummyOpts.Color(colors[i%(len(colors))]),
				ui.DummyOpts.Height(
					i*8+4,
				),
			)
			hbox.AddChild(dummy)

			var s = fmt.Sprintf("[%02d]", i+1)
			var text = ui.NewText(
				ui.TextOpts.Color(colors[len(colors)-i%(len(colors))-1]),
				ui.TextOpts.Content(s),
			)
			hbox.AddChild(text)
		}

		vbox.AddChild(m(b(hbox)))
	}

	app.Load(vbox)

	ebiten.SetWindowTitle("hbox demo")
	ebiten.SetWindowSize(800, 800)
	ebiten.RunGame(app.Get())
}
