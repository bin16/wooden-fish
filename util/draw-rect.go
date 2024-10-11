package util

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawRectOptions struct {
	strokeWidth     int
	color           color.Color
	backgroundColor color.Color
	radius          int
}

type DrawRectOpt func(op *DrawRectOptions)

var DrawRectOpts DrawRectOptions

func NewDrawRectOpts() *DrawRectOptions {
	return &DrawRectOptions{
		strokeWidth: 1,
	}
}

func (DrawRectOptions) StrokeWidth(d int) DrawRectOpt {
	return func(op *DrawRectOptions) {
		op.strokeWidth = d
	}
}

func (DrawRectOptions) Color(clr color.Color) DrawRectOpt {
	return func(op *DrawRectOptions) {
		op.color = clr
	}
}

func (DrawRectOptions) Fill(clr color.Color) DrawRectOpt {
	return func(op *DrawRectOptions) {
		op.backgroundColor = clr
	}
}

func (DrawRectOptions) Radius(d int) DrawRectOpt {
	return func(op *DrawRectOptions) {
		op.radius = d
	}
}

func DrawRect(screen *ebiten.Image, dr image.Rectangle, opts ...DrawRectOpt) {
	var op = NewDrawRectOpts()
	for _, o := range opts {
		o(op)
	}

	var (
		d   = op.strokeWidth
		clr = op.color
		rad = op.radius
	)
	rad = max(0, rad)
	rad = min(rad, 5)

	if d == 0 || clr == nil {
		return
	}

	var dl = [][]int{
		{},
		{1},
		{2, 1},
		{3, 2, 1},
		{4, 3, 2, 1},
		{5, 3, 2, 1, 1},
	}

	if clr := op.backgroundColor; clr != nil {
		screen.SubImage(
			image.Rect(0, rad, dr.Dx(), dr.Dy()-rad),
		).(*ebiten.Image).Fill(clr)
		for i := 0; i < rad; i++ {
			var d = dl[rad][i]
			screen.SubImage(
				image.Rect(d, i, dr.Dx()-d, i+1),
			).(*ebiten.Image).Fill(clr)
			screen.SubImage(
				image.Rect(d, dr.Dy()-i-1, dr.Dx()-d, dr.Dy()-i),
			).(*ebiten.Image).Fill(clr)
		}
	}

	var (
		x0 = 0
		x1 = d
		x2 = rad
		x3 = dr.Dx() - rad
		x4 = dr.Dx() - d
		x5 = dr.Dx()

		y0 = 0
		y1 = d
		y2 = rad
		y3 = dr.Dy() - rad
		y4 = dr.Dy() - d
		y5 = dr.Dy()
	)

	var rl = []image.Rectangle{
		image.Rect(x2, y0, x3, y1),
		image.Rect(x2, y4, x3, y5),
		image.Rect(x0, y2, x1, y3),
		image.Rect(x4, y2, x5, y3),
	}

	for _, r := range rl {
		screen.SubImage(r.Add(dr.Min)).(*ebiten.Image).Fill(clr)
	}

	var pl = [][]image.Point{
		{
			image.Pt(0, 0),
		},
		{},
		{
			image.Pt(1, 1),
		},
		{
			image.Pt(1, 2),
			image.Pt(2, 1),
		},
		{
			image.Pt(1, 3),
			image.Pt(2, 2),
			image.Pt(3, 1),
		},
		{
			image.Pt(1, 4),
			image.Pt(1, 3),
			image.Pt(2, 2),
			image.Pt(3, 1),
			image.Pt(4, 1),
		},
	}

	for _, p := range pl[rad] {
		var (
			p0 = dr.Min.Add(image.Pt(p.X, p.Y))
			p1 = dr.Min.Add(image.Pt(dr.Dx()-1-p.X, p.Y))
			p2 = dr.Min.Add(image.Pt(dr.Dx()-1-p.X, dr.Dy()-1-p.Y))
			p3 = dr.Min.Add(image.Pt(p.X, dr.Dy()-1-p.Y))
		)

		screen.Set(p0.X, p0.Y, clr)
		screen.Set(p1.X, p1.Y, clr)
		screen.Set(p2.X, p2.Y, clr)
		screen.Set(p3.X, p3.Y, clr)
	}
}
