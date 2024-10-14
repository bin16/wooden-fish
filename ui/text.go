package ui

import (
	"image/color"
	"math"

	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

type Text struct {
	Scene

	textContent string
	face        text.Face
	color       color.Color
	padding     app.Space
	alpha       float64

	getFn func() string
}

func (u *Text) SetColor(clr color.Color) {
	u.color = clr
}

func (u *Text) SetContent(s string) {
	u.textContent = s
}

func (u *Text) SetFace(face text.Face) {
	u.face = face
}

func (u *Text) SetFontFace(face font.Face) {
	u.face = text.NewGoXFace(face)
}

func (u *Text) Layout(ow, oh int) (bw, bh int) {
	var (
		lh = u.face.Metrics().HAscent + u.face.Metrics().HDescent + u.face.Metrics().HLineGap
	)

	if u.textContent == "" {
		return 1, int(lh)
	}

	var tw, th = text.Measure(u.textContent, u.face, lh)
	bw = int(math.Ceil(tw)) + u.padding.X()
	bh = int(math.Ceil(th)) + u.padding.Y()
	return
}

func (u *Text) Draw(screen *ebiten.Image) {
	var op = &text.DrawOptions{}
	if clr := u.color; clr != nil {
		op.ColorScale.ScaleWithColor(clr)
	} else if clr := app.Theme.Color; clr != nil {
		op.ColorScale.ScaleWithColor(clr)
	}

	if a := u.alpha; a != 1 {
		op.ColorScale.ScaleAlpha(float32(a))
	}

	var p = u.Bounds().Min.Add(u.padding.TopLeft())
	op.GeoM.Translate(
		float64(p.X),
		float64(p.Y),
	)

	if u.IsDisabled() {
		op.ColorScale.ScaleAlpha(.5)
	}

	text.Draw(screen, u.textContent, u.face, op)
}

func (u *Text) Update() error {
	if fn := u.getFn; fn != nil {
		var s = u.getFn()
		u.SetContent(s)
	}

	return u.Scene.Update()
}

type TextOpt func(text *Text)
type TextOptions struct{}

var TextOpts TextOptions

func (TextOptions) Padding(num ...int) TextOpt {
	return func(text *Text) {
		text.padding = app.NewSpace(num...)
	}
}

func (TextOptions) Content(s string) TextOpt {
	return func(text *Text) {
		text.SetContent(s)
	}
}

func (TextOptions) Color(clr color.Color) TextOpt {
	return func(text *Text) {
		text.SetColor(clr)
	}
}

func (TextOptions) Face(face text.Face) TextOpt {
	return func(text *Text) {
		text.SetFace(face)
	}
}

func (TextOptions) FontFace(face font.Face) TextOpt {
	return func(text *Text) {
		text.SetFontFace(face)
	}
}

func (TextOptions) Pull(fn func() string) TextOpt {
	return func(text *Text) {
		text.getFn = fn
	}
}

func (TextOptions) SetDisabled(b bool) TextOpt {
	return func(text *Text) {
		text.SetDisabled(b)
	}
}

func NewText(opts ...TextOpt) *Text {
	var text = &Text{
		face:  text.NewGoXFace(bitmapfont.FaceSC),
		alpha: 1.0,
	}
	for _, o := range opts {
		o(text)
	}

	return text
}
