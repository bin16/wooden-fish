package ui

import (
	"image"
	"image/color"
	"time"

	"github.com/bin16/wooden-fish/animator/curve"
	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type upperStage uint8

const (
	upperInit upperStage = iota
	upperAnimIn
	upperKeep
	upperAnimOut
	upperDone
)

type upperConfig struct {
	item   app.Scene
	stage  upperStage
	curve  *curve.Curve
	offset image.Point
	alpha  float64
	color  color.Color
}

func (u *Upper) Draw(screen *ebiten.Image) {
	// util.StrokeRect(screen, u.Bounds(), hexcolor.New("#cf0"), 2)

	u.Box.Draw(screen)
}

func (g *upperConfig) Update() error {
	if g.stage == upperInit {
		g.curve = curve.New(
			curve.Duration(time.Second/2),
			curve.AutoPlay(),
		)

		return nil
	}

	if err := g.curve.Update(); err != nil {
		return err
	}

	if g.stage == upperAnimIn {
		if g.curve.IsEnd() {
			g.stage = upperAnimOut
			g.curve = curve.New(
				curve.Duration(time.Second),
				curve.AutoPlay(),
			)

			return nil
		}

		var (
			q = g.curve.Q()
			y = curve.Apply(q, 32, 0)
		)

		g.alpha = curve.Apply[float64](q, 0, 1)
		g.offset = image.Pt(0, y)
		g.color = curve.ApplyColor(q, color.Transparent, app.Theme.Color)
	}

	if g.stage == upperAnimOut {
		if g.curve.IsEnd() {
			g.stage = upperDone
			return nil
		}

		var (
			q = g.curve.Q()
			y = curve.Apply(q, 0, -16)
		)

		g.alpha = curve.Apply[float64](q, 1, 0)
		g.offset = image.Pt(0, y)
		g.color = curve.ApplyColor(q, app.Theme.Color, color.Transparent)
	}

	return nil
}

type Upper struct {
	Box

	maxWidth  int
	maxHeight int
	config    []*upperConfig
}

func (u *Upper) Layout(ow, oh int) (bw, bh int) {
	var (
		cnt = len(u.Children())
		mcw = util.NotZero(u.maxWidth, ow)
		mch = util.NotZero(u.maxHeight, oh)
		cw  = 0
		ch  = 0
	)

	if len(u.sr) != cnt {
		u.sr = make([]image.Rectangle, cnt)
	}
	for i, n := range u.Children() {
		var w, h = n.Layout(mcw, mch)
		cw = max(cw, w)
		ch = max(ch, h)
		u.sr[i] = image.Rect(0, 0, w, h)
	}

	bw = mcw
	bh = mch
	return
}

func (u *Upper) New(content app.Scene) {
	var item = NewUpperItem(
		UpperItemOpts.Content(
			Center(content),
		),
	)

	u.AddChild(item)
}

func (u *Upper) NewText(s string) {
	var text = NewText(
		TextOpts.Content(s),
	)

	u.New(text)
}

func (u *Upper) Update() error {
	if len(u.children) > 0 {
		if u.children[0].(*UpperItem).stage == upperDone {
			u.children = u.children[1:]
		}
	}

	return u.Box.Update()
}

func NewUpper() *Upper {
	var u = &Upper{
		maxWidth:  64,
		maxHeight: 16,
	}

	return u
}
