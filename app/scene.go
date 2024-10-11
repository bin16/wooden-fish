package app

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Layout(ow, oh int) (bw, bh int)
	Draw(screen *ebiten.Image)
	SetBounds(r image.Rectangle)
	Bounds() (r image.Rectangle)
	HandleInput() bool
	HandleMouseInput() bool
}
