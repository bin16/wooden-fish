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
	onEnd        []func()
	playing      bool
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

	if u.Bounds().Empty() {
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
	if u.playing {
		u.tick += 1
		if u.tick > ebiten.TPS()/u.FPS {
			u.tick = 0
			u.index += 1
			u.triggerFrame(u.index)
		}

		var max = u.image.Bounds().Dy()/u.height - 1
		if u.index > max {
			u.index = 0
			for _, fn := range u.onEnd {
				fn()
			}
			u.playing = u.loop
		}
	}

	return nil
}

func (u *Anim) Play() {
	if !u.playing {
		u.playing = true
		u.index = 0
		u.tick = 0
	}
}

func (u *Anim) IsPlaying() bool {
	return u.playing
}

func (u *Anim) FrameIndex() int {
	return u.index
}

func (u *Anim) OnEnd(fn func()) {
	u.onEnd = append(u.onEnd, fn)
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

func (AnimOptions) Loop(b bool) AnimOpt {
	return func(u *Anim) {
		u.loop = b
	}
}

func (AnimOptions) AutoPlay(b bool) AnimOpt {
	return func(u *Anim) {
		u.playing = true
	}
}

func (AnchorOptions) OnEnd(fn func()) AnimOpt {
	return func(u *Anim) {
		u.OnEnd(fn)
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
