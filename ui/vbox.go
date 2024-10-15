package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/ebiten/v2"
)

type AlignItems uint8

const (
	AlignStart AlignItems = iota
	AlignCenter
	AlignEnd
	AlignStretch
)

type VBox struct {
	Box

	alignItems     AlignItems
	justifyContent JustifyContent

	maxHeight int // for scroll
	offset    image.Point
}

func (u *VBox) Layout(ow, oh int) (bw, bh int) {
	var (
		cw int
		ch int
	)

	if d := len(u.children); len(u.sr) != d {
		u.sr = make([]image.Rectangle, d)
	}
	for i, n := range u.children {
		var (
			w, h = n.Layout(ow, oh)
			x    = 0
			y    = ch
		)
		cw = max(cw, w)
		ch = ch + h

		u.sr[i] = image.Rect(0, 0, w, h).Add(image.Pt(x, y))
	}
	u.cr = image.Rect(0, 0, cw, ch)

	bw = cw
	bh = ch
	if u.justifyContent != JustifyStart {
		bh = oh
	}

	mh := u.maxHeight
	if mh == 0 {
		mh = oh
	}

	if bh > mh {
		bh = mh
	}

	return
}

func (u *VBox) SetBounds(r image.Rectangle) {
	u.Scene.SetBounds(r)

	var (
		cnt = len(u.Children())
		sp  = r.Dy() - u.cr.Dy()
		z0  = 0
		z   = 0
	)
	switch u.justifyContent {
	case JustifyCenter:
		z0 = sp / 2
	case JustifyEnd:
		z0 = sp
	case SpaceBetween:
		if cnt == 1 {
			z = sp / 2
		} else {
			z = sp / (cnt - 1)
		}
	}

	for i, n := range u.children {
		var (
			r1 = u.sr[i]
			x  = 0
			y  = r1.Min.Y + z0 + z*i
			w  = r1.Dx()
			h  = r1.Dy()
		)

		switch u.alignItems {
		case AlignCenter:
			x = (r.Dx() - u.sr[i].Dx()) / 2
		case AlignEnd:
			x = (r.Dx() - u.sr[i].Dx())
		case AlignStretch:
			w = r.Dx()
		}

		n.SetBounds(image.Rect(0, 0, w, h).Add(image.Pt(x, y)).Add(r.Min).Add(u.offset))
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

func (VBoxOptions) JustifyContent(d JustifyContent) VBoxOpt {
	return func(box *VBox) {
		box.justifyContent = d
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

func (u *VBox) Draw(screen *ebiten.Image) {
	// util.StrokeRect(screen, u.bounds, hexcolor.New("#0c0"))
	// u.Box.Draw(screen)
	for _, n := range u.children {
		if n.Bounds().Overlaps(u.bounds) {
			n.Draw(screen)
		}
	}
}

func NewVBox(opts ...VBoxOpt) *VBox {
	var box = &VBox{}
	for _, o := range opts {
		o(box)
	}

	return box
}
