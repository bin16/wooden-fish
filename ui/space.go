package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
)

type Space struct {
	Box

	Space app.Space
}

func (u *Space) Layout(ow, oh int) (bw, bh int) {
	var (
		mcw    = ow - u.Space.X()
		mch    = oh - u.Space.Y()
		cw, ch int
	)

	for _, n := range u.Children() {
		var w, h = n.Layout(mcw, mch)
		cw = max(cw, w)
		ch = max(ch, h)
	}

	bw = cw + u.Space.X()
	bh = ch + u.Space.Y()
	return
}

func (u *Space) SetBounds(r image.Rectangle) {
	u.Scene.SetBounds(r)
	for _, n := range u.Children() {
		var (
			w = r.Dx() - u.Space.X()
			h = r.Dy() - u.Space.Y()
			x = r.Min.X + u.Space.Left
			y = r.Min.Y + u.Space.Top
		)

		n.SetBounds(image.Rect(0, 0, w, h).Add(image.Pt(x, y)))
	}
}

type SpaceOpt func(u *Space)
type SpaceOptions struct{}

var SpaceOpts SpaceOptions

func (SpaceOptions) Space(num ...int) SpaceOpt {
	return func(u *Space) {
		u.Space = app.NewSpace(num...)
	}
}

func (SpaceOptions) Content(n app.Scene) SpaceOpt {
	return func(u *Space) {
		u.children = []app.Scene{n}
	}
}

func NewSpace(opts ...SpaceOpt) func(n app.Scene) *Space {
	return func(n app.Scene) *Space {
		var space = &Space{}
		for _, o := range opts {
			o(space)
		}
		SpaceOpts.Content(n)(space)

		return space
	}
}
