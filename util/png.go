package util

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewImageFromBytes(d []byte) *ebiten.Image {
	img, err := png.Decode(bytes.NewReader(d))
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
