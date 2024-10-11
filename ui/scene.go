package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	size   image.Point
	bounds image.Rectangle
}

func (u *Scene) Update() error {
	return nil
}

func (u *Scene) Draw(screen *ebiten.Image) {}
func (u *Scene) Layout(ow, oh int) (bw, bh int) {
	return
}

func (u *Scene) SetBounds(r image.Rectangle) {
	u.bounds = r
}

func (u *Scene) Bounds() (r image.Rectangle) {
	return u.bounds
}

func (u *Scene) HandleInput() bool {
	return false
}

func (u *Scene) HandleMouseInput() bool {
	return false
}
