package ui

import (
	"image"
	"image/color"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Border struct {
	Box

	Border int
	Color  color.Color
	Radius int
}

func (u *Border) Layout(ow, oh int) (bw, bh int) {
	var (
		mcw    = ow - u.Border*2
		mch    = oh - u.Border*2
		cw, ch int
	)

	for _, n := range u.Children() {
		var w, h = n.Layout(mcw, mch)
		cw = max(cw, w)
		ch = max(ch, h)
	}

	bw = cw + u.Border*2
	bh = ch + u.Border*2
	return
}

func (u *Border) SetBounds(r image.Rectangle) {
	u.Scene.SetBounds(r)
	for _, n := range u.Children() {
		n.SetBounds(r.Inset(u.Border))
	}
}

func (u *Border) Draw(screen *ebiten.Image) {
	util.StrokeRect(screen, u.Bounds(), u.Color, u.Radius)

	u.Box.Draw(screen)
}

type BorderOpt func(u *Border)
type BorderOptions struct{}

var BorderOpts BorderOptions

func (BorderOptions) Border(d int) BorderOpt {
	return func(u *Border) {
		u.Border = d
	}
}

func (u *BorderOptions) Color(clr color.Color) BorderOpt {
	return func(u *Border) {
		u.Color = clr
	}
}

func (u *BorderOptions) BorderRadius(d int) BorderOpt {
	return func(u *Border) {
		u.Radius = d
	}
}

func (u *BorderOptions) Content(n app.Scene) BorderOpt {
	return func(u *Border) {
		u.children = []app.Scene{n}
	}
}

func NewBorder(opts ...BorderOpt) func(n app.Scene) *Border {
	return func(n app.Scene) *Border {
		var border = &Border{}
		for _, o := range opts {
			o(border)
		}
		BorderOpts.Content(n)(border)

		return border
	}
}
