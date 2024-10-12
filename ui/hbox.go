package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/ebiten/v2"
)

type JustifyContent uint8

const (
	JustifyStart JustifyContent = iota
	JustifyCenter
	JustifyEnd
	SpaceBetween
)

type HBox struct {
	Box

	alignItems     AlignItems
	justifyContent JustifyContent
}

func (u *HBox) Layout(ow, oh int) (bw, bh int) {
	var (
		cw = 0
		ch = 0
	)

	if d := len(u.children); len(u.sr) != d {
		u.sr = make([]image.Rectangle, d)
	}
	for i, n := range u.children {
		var (
			w, h = n.Layout(ow, oh)
			x    = cw
			y    = 0
		)
		ch = max(ch, h)
		cw += w

		u.sr[i] = image.Rect(0, 0, w, h).Add(image.Pt(x, y))
	}
	u.cr = image.Rect(0, 0, cw, ch)

	bw = ow
	bh = ch
	return
}

func (u *HBox) SetBounds(r image.Rectangle) {
	u.Scene.SetBounds(r)

	var (
		cnt = len(u.Children())
		sp  = r.Dx() - u.cr.Dx()
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
			w  = r1.Dx()
			h  = r1.Dy()
			x  = r1.Min.X + z0 + z*i
			y  = 0
		)
		switch u.alignItems {
		case AlignCenter:
			y = (r.Dy() - r1.Dy()) / 2
		case AlignEnd:
			y = (r.Dy() - r1.Dy())
		case AlignStretch:
			h = r.Dy()
		}

		n.SetBounds(image.Rect(0, 0, w, h).Add(image.Pt(x, y)).Add(r.Min))
	}
}

func (u *HBox) Draw(screen *ebiten.Image) {
	u.Box.Draw(screen)

	// // debug
	// util.StrokeRect(screen, u.Bounds().Inset(-1), hexcolor.New("#0c0"), 2)
}

type HBoxOpt func(box *HBox)
type HBoxOptions struct{}

var HBoxOpts HBoxOptions

func (HBoxOptions) AlignItems(d AlignItems) HBoxOpt {
	return func(box *HBox) {
		box.alignItems = d
	}
}

func (HBoxOptions) JustifyContent(d JustifyContent) HBoxOpt {
	return func(box *HBox) {
		box.justifyContent = d
	}
}

func (HBoxOptions) BoxOpts(opts ...BoxOpt) HBoxOpt {
	return func(box *HBox) {
		for _, o := range opts {
			o(&(box.Box))
		}
	}
}

func (HBoxOptions) Contents(items ...app.Scene) HBoxOpt {
	return func(box *HBox) {
		box.children = append(box.children, items...)
	}
}

func NewHBox(opts ...HBoxOpt) *HBox {
	var box = &HBox{}
	for _, o := range opts {
		o(box)
	}

	return box
}
