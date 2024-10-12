package ui

import (
	"image/color"

	"github.com/bin16/go-hexcolor"
	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Dummy struct {
	Scene

	Width       int
	Height      int
	Color       color.Color
	StrokeWidth int
	StrokeColor color.Color
}

func (u *Dummy) Layout(ow, oh int) (bw, bh int) {
	bw = u.Width
	if bw == 0 {
		bw = 16
	}

	bh = u.Height
	if bh == 0 {
		bh = 16
	}

	return
}

func (u *Dummy) Draw(screen *ebiten.Image) {
	// fmt.Println(u.Bounds())
	util.DrawRect(
		screen,
		u.Bounds(),
		util.DrawRectOpts.Fill(u.Color),
		util.DrawRectOpts.Radius(2),
		util.DrawRectOpts.StrokeWidth(u.StrokeWidth),
		util.DrawRectOpts.Color(u.StrokeColor),
	)

	// screen.SubImage(u.Bounds()).(*ebiten.Image).Fill(color.White)
}

type DummyOpt func(u *Dummy)
type DummyOptions struct{}

var DummyOpts DummyOptions

func (DummyOptions) Width(d int) DummyOpt {
	return func(u *Dummy) {
		u.Width = d
	}
}

func (DummyOptions) Height(d int) DummyOpt {
	return func(u *Dummy) {
		u.Height = d
	}
}

func (DummyOptions) Size(w, h int) DummyOpt {
	return func(u *Dummy) {
		u.Width = w
		u.Height = h
	}
}

func (DummyOptions) Color(clr color.Color) DummyOpt {
	return func(u *Dummy) {
		u.Color = clr
	}
}

func NewDummy(opts ...DummyOpt) *Dummy {
	var dummy = &Dummy{
		Width:       16,
		Height:      16,
		Color:       hexcolor.New("#FFBC42"),
		StrokeWidth: 1,
		StrokeColor: hexcolor.New("#73D2DE"),
	}
	for _, o := range opts {
		o(dummy)
	}

	return dummy
}
