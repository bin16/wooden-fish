package ui

import (
	"image"
	"time"

	"github.com/bin16/wooden-fish/animator/curve"
	"github.com/bin16/wooden-fish/animator/easing"
	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/ebiten/v2"
)

type UpperItem struct {
	Box

	curve *curve.Curve
	stage upperStage

	y0, y1 int
	a0, a1 float32

	alpha  float32
	offset image.Point
	cache  *ebiten.Image
}

func (u *UpperItem) AnimIn() {
	u.a0 = 0.0
	u.a1 = 1.0
	u.y0 = 4
	u.y1 = 0
	u.curve = curve.New(
		curve.Easing(easing.EaseOutCirc),
		curve.Duration(time.Second/2),
		curve.AutoPlay(),
	)
	u.stage = upperAnimIn
}

func (u *UpperItem) AnimOut() {
	u.a0 = 1.0
	u.a1 = 0.1
	u.y0 = 0
	u.y1 = -8
	u.curve = curve.New(
		curve.Easing(easing.EaseInCirc),
		curve.Duration(time.Second/2),
		curve.AutoPlay(),
	)
	u.stage = upperAnimOut
}

func (u *UpperItem) Draw(screen *ebiten.Image) {
	if u.cache == nil || u.cache.Bounds() != screen.Bounds() {
		u.cache = ebiten.NewImage(
			screen.Bounds().Dx(),
			screen.Bounds().Dy(),
		)
	}
	u.cache.Clear()

	var op = &ebiten.DrawImageOptions{}

	if u.stage == upperAnimIn || u.stage == upperAnimOut {
		for _, n := range u.Children() {
			n.Draw(u.cache)
		}
	}

	op.ColorScale.ScaleAlpha(u.alpha)
	op.GeoM.Translate(
		float64(u.offset.X),
		float64(u.offset.Y),
	)

	screen.DrawImage(u.cache, op)
}

func (u *UpperItem) Update() error {
	if u.stage == upperInit {
		u.AnimIn()
		return nil
	}

	if err := u.curve.Update(); err != nil {
		return err
	}

	var (
		q = u.curve.Q()
		y = curve.Apply(q, u.y0, u.y1)
	)

	u.alpha = curve.Apply[float32](q, u.a0, u.a1)
	u.offset = image.Pt(0, y)

	if u.stage == upperAnimIn {
		if u.curve.IsEnd() {
			u.AnimOut()
			return nil
		}

	}

	if u.stage == upperAnimOut {
		if u.curve.IsEnd() {
			u.stage = upperDone
			return nil
		}
	}

	return nil
}

type UpperItemOpt func(u *UpperItem)
type UpperItemOptions struct{}

var UpperItemOpts UpperItemOptions

func (UpperItemOptions) Content(n app.Scene) UpperItemOpt {
	return func(u *UpperItem) {
		u.children = []app.Scene{n}
	}
}

func NewUpperItem(opts ...UpperItemOpt) *UpperItem {
	var item = &UpperItem{}
	for _, o := range opts {
		o(item)
	}

	return item
}
