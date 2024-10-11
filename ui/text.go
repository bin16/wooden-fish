package ui

import (
	"image/color"
	"math"

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
	bw = int(math.Ceil(tw))
	bh = int(math.Ceil(th))
	return
}

func (u *Text) Draw(screen *ebiten.Image) {
	var op = &text.DrawOptions{}
	if clr := u.color; clr != nil {
		op.ColorScale.ScaleWithColor(clr)
	}

	op.GeoM.Translate(
		float64(u.Bounds().Min.X),
		float64(u.Bounds().Min.Y),
	)

	text.Draw(screen, u.textContent, u.face, op)
}

type TextOpt func(text *Text)
type TextOptions struct{}

var TextOpts TextOptions

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

func NewText(opts ...TextOpt) *Text {
	var text = &Text{
		face:  text.NewGoXFace(bitmapfont.FaceSC),
		color: color.Black,
	}
	for _, o := range opts {
		o(text)
	}

	return text
}
