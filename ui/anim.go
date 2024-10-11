package ui

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"

	"github.com/bin16/go-hexcolor"
	"github.com/hajimehoshi/ebiten/v2"
)

type Anim struct {
	Scene
	index        int
	loop         bool
	width        int
	height       int
	image        *ebiten.Image
	maxIndex     int
	FPS          int
	tick         int
	maxTick      int
	frameHandler map[int][]func()
}

func (u *Anim) Layout(ow, oh int) (bw, bh int) {
	return u.width, u.height
}

func (u *Anim) Draw(screen *ebiten.Image) {
	if u.image == nil {
		screen.SubImage(u.Bounds()).(*ebiten.Image).Fill(
			hexcolor.Parse("#c00"),
		)

		return
	}

	var op = &ebiten.DrawImageOptions{}

	op.GeoM.Translate(
		float64(u.Bounds().Min.X),
		float64(u.Bounds().Min.Y),
	)

	screen.DrawImage(
		u.image.SubImage(
			image.Rect(0, 0, u.width, u.height).Add(image.Pt(0, u.height*u.index)),
		).(*ebiten.Image),
		op,
	)
}

func (u *Anim) triggerFrame(index int) {
	if u.frameHandler == nil {
		return
	}

	if handlers, ok := u.frameHandler[index]; ok {
		for _, fn := range handlers {
			fn()
		}
	}
}

func (u *Anim) Update() error {
	u.tick += 1
	if u.tick > ebiten.TPS()/u.FPS {
		u.tick = 0
		u.index += 1
		u.triggerFrame(u.index)
	}

	var max = u.image.Bounds().Dy()/u.height - 1
	if u.index > max {
		u.index = 0
	}

	return nil
}

type AnimOpt func(u *Anim)
type AnimOptions struct{}

var AnimOpts AnimOptions

func (AnimOptions) NewImageFromBytes(d []byte) AnimOpt {
	var img, _ = png.Decode(bytes.NewReader(d))
	return func(u *Anim) {
		u.image = ebiten.NewImageFromImage(img)
	}
}

func (AnimOptions) Image(img *ebiten.Image) AnimOpt {
	return func(u *Anim) {
		u.image = img
	}
}

func (AnimOptions) Size(w, h int) AnimOpt {
	return func(u *Anim) {
		u.width = w
		u.height = h
	}
}

func (AnimOptions) FPS(d int) AnimOpt {
	return func(u *Anim) {
		u.FPS = d
	}
}

func (AnimOptions) OnFrame(index int, fn func()) AnimOpt {
	return func(u *Anim) {
		if u.frameHandler == nil {
			u.frameHandler = make(map[int][]func())
		}

		u.frameHandler[index] = append(u.frameHandler[index], fn)
	}
}

func NewAnim(opts ...AnimOpt) *Anim {
	// TODO: debug
	var img = ebiten.NewImage(32, 32)
	img.Fill(hexcolor.New("#0c0"))

	var anim = &Anim{
		width:  32,
		height: 32,
		image:  img,
		FPS:    9,
	}
	for _, o := range opts {
		o(anim)
	}

	return anim
}
