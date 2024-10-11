package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
)

type AlignItems uint8

const (
	AlignStart AlignItems = iota
	AlignCenter
	AlignEnd
)

type VBox struct {
	Box

	alignItems AlignItems
}

func (u *VBox) Layout(ow, oh int) (bw, bh int) {
	if d := len(u.children); len(u.sr) != d {
		u.sr = make([]image.Rectangle, d)
	}
	for i, n := range u.children {
		var (
			w, h = n.Layout(ow, oh)
			x    = 0
			y    = bh
		)
		bw = max(bw, w)
		bh = bh + h

		u.sr[i] = image.Rect(0, 0, w, h).Add(image.Pt(x, y))
	}

	return
}

func (u *VBox) SetBounds(r image.Rectangle) {
	u.Scene.SetBounds(r)

	for i, n := range u.children {
		var x = 0
		switch u.alignItems {
		case AlignCenter:
			x = (r.Dx() - u.sr[i].Dx()) / 2
		case AlignEnd:
			x = (r.Dx() - u.sr[i].Dx())
		}

		n.SetBounds(u.sr[i].Add(r.Min).Add(image.Pt(x, 0)))
	}
}

type VBoxOpt func(box *VBox)
type VBoxOptions struct{}

var VBoxOpts VBoxOptions

func (VBoxOptions) AlignItems(d AlignItems) VBoxOpt {
	return func(box *VBox) {
		box.alignItems = d
	}
}

func (VBoxOptions) BoxOpts(opts ...BoxOpt) VBoxOpt {
	return func(box *VBox) {
		for _, o := range opts {
			o(&(box.Box))
		}
	}
}

func (VBoxOptions) Contents(items ...app.Scene) VBoxOpt {
	return func(box *VBox) {
		box.children = append(box.children, items...)
	}
}

func NewVBox(opts ...VBoxOpt) *VBox {
	var box = &VBox{}
	for _, o := range opts {
		o(box)
	}

	return box
}
