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

	if clr := op.backgroundColor; clr != nil {
		FillRect(screen, dr, clr)
	}

	if clr := op.color; clr != nil {
		var rad = op.radius
		StrokeRect(screen, dr, clr, rad)
	}
}
