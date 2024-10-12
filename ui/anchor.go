package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
)

type Anchor struct {
	Box

	hAlign AlignItems
	vAlign AlignItems
}

func (u *Anchor) Layout(ow, oh int) (bw, bh int) {
	var (
		cw, ch int
	)

	for _, n := range u.Children() {
		var w, h = n.Layout(ow, oh)
		cw = max(cw, w)
		ch = max(ch, h)
	}
	u.cr = image.Rect(0, 0, cw, ch)

	return ow, oh
}

func (u *Anchor) SetBounds(r image.Rectangle) {
	u.Scene.SetBounds(r)

	for _, n := range u.Children() {
		var (
			w = u.cr.Dx()
			h = u.cr.Dy()
			x = 0
			y = 0
		)

		switch u.hAlign {
		case AlignCenter:
			x = (r.Dx() - u.cr.Dx()) / 2
		case AlignEnd:
			x = (r.Dx() - u.cr.Dx())
		case AlignStretch:
			w = r.Dx()
		}

		switch u.vAlign {
		case AlignCenter:
			y = (r.Dy() - u.cr.Dy()) / 2
		case AlignEnd:
			y = r.Dy() - u.cr.Dy()
		case AlignStretch:
			h = r.Dy()
		}

		n.SetBounds(image.Rect(0, 0, w, h).Add(image.Pt(x, y)).Add(r.Min))
	}
}

type AnchorOpt func(u *Anchor)
type AnchorOptions struct{}

var AnchorOpts AnchorOptions

func (AnchorOptions) HAlign(d AlignItems) AnchorOpt {
	return func(u *Anchor) {
		u.hAlign = d
	}
}

func (AnchorOptions) VAlign(d AlignItems) AnchorOpt {
	return func(u *Anchor) {
		u.vAlign = d
	}
}

func (AnchorOptions) Content(n app.Scene) AnchorOpt {
	return func(u *Anchor) {
		u.children = []app.Scene{n}
	}
}

func NewAnchor(opts ...AnchorOpt) func(content app.Scene) *Anchor {
	return func(content app.Scene) *Anchor {
		var anchor = &Anchor{}
		for _, o := range opts {
			o(anchor)
		}
		AnchorOpts.Content(content)(anchor)

		return anchor
	}
}

func Top(content app.Scene) *Anchor {
	return NewAnchor(
		AnchorOpts.HAlign(AlignCenter),
		AnchorOpts.VAlign(AlignStart),
	)(content)
}

func Center(content app.Scene) *Anchor {
	return NewAnchor(
		AnchorOpts.HAlign(AlignCenter),
		AnchorOpts.VAlign(AlignCenter),
	)(content)
}

func Bottom(content app.Scene) *Anchor {
	return NewAnchor(
		AnchorOpts.HAlign(AlignStretch),
		AnchorOpts.VAlign(AlignEnd),
	)(content)
}

func BottomLeft(content app.Scene) *Anchor {
	return NewAnchor(
		AnchorOpts.HAlign(AlignStart),
		AnchorOpts.VAlign(AlignEnd),
	)(content)
}

func BottomRight(content app.Scene) *Anchor {
	return NewAnchor(
		AnchorOpts.HAlign(AlignEnd),
		AnchorOpts.VAlign(AlignEnd),
	)(content)
}
