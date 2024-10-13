package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/ebiten/v2"
)

// Tap area
type TapBox struct {
	Box

	events app.Events

	mouseover bool
	mousedown bool
}

func (u *TapBox) HandleTouchInput() bool {
	if app.IsTappedInBounds(u.bounds) {
		u.events.Emit("touch")
		u.events.Emit("tap")
		return true
	}

	return app.IsTappedInBounds(u.bounds)
}

func (u *TapBox) HandleMouseInput() bool {
	var (
		p         = image.Pt(ebiten.CursorPosition())
		mouseover = p.In(u.Bounds())
		mousedown = mouseover && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	)

	if mousedown && !u.mousedown {
		u.events.Emit("mousedown")
	}

	if !mousedown && mousedown {
		u.events.Emit("mouseup")
		u.events.Emit("press") // or click?
		u.events.Emit("tap")
	}

	if mouseover && !u.mouseover {
		u.events.Emit("mouseenter")
	}

	if !mouseover && u.mouseover {
		u.events.Emit("mouseleave")
	}

	return false
}

func (u *TapBox) OnTap(fn func(data ...any)) {
	u.events.On("tap", fn)
}

type TapBoxOpt func(u *TapBox)
type TapBoxOptions struct{}

var TapBoxOpts TapBoxOptions

func (TapBoxOptions) Content(n app.Scene) TapBoxOpt {
	return func(u *TapBox) {
		u.children = []app.Scene{n}
	}
}

func (TapBoxOptions) OnTap(fn func(data ...any)) TapBoxOpt {
	return func(u *TapBox) {
		u.OnTap(fn)
	}
}

func NewTapBox(opts ...TapBoxOpt) *TapBox {
	var box = &TapBox{}
	for _, o := range opts {
		o(box)
	}

	return box
}

func OnTap(content app.Scene, fn func(data ...any)) *TapBox {
	return NewTapBox(
		TapBoxOpts.Content(content),
		TapBoxOpts.OnTap(fn),
	)
}
