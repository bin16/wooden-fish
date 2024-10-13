package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/util"
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
		if util.AnyOf(
			u.events.Emit("touch"),
			u.events.Emit("tap"),
		) {
			return true
		}
	}

	return false
}

func (u *TapBox) HandleMouseInput() bool {
	var (
		p         = image.Pt(ebiten.CursorPosition())
		mouseover = p.In(u.Bounds())
		mousedown = mouseover && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
		flag      = false
	)

	if mouseover {
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
	}

	if mousedown && !u.mousedown {
		flag = flag || u.events.Emit("mousedown")
		flag = flag || u.events.Emit("press") // or click?
		flag = flag || u.events.Emit("tap")
	}

	if !mousedown && mousedown {
		flag = flag || u.events.Emit("mouseup")
	}

	if mouseover && !u.mouseover {
		flag = flag || u.events.Emit("mouseenter")
	}

	if !mouseover && u.mouseover {
		flag = flag || u.events.Emit("mouseleave")
	}

	u.mousedown = mousedown
	u.mouseover = mouseover

	return flag
}

func (u *TapBox) OnTap(fn app.EventFunc) {
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

func (TapBoxOptions) OnTap(fn app.EventFunc) TapBoxOpt {
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

func OnTap(content app.Scene, fn app.EventFunc) *TapBox {
	return NewTapBox(
		TapBoxOpts.Content(content),
		TapBoxOpts.OnTap(fn),
	)
}
