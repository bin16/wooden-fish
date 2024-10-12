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
	HasFocus() bool
	IsFocusable() bool
	IsFocused() bool
	Focus()
	Blur()
	HandleFocus(i int) bool
	FocusUp() bool
	FocusDown() bool
	FocusLeft() bool
	FocusRight() bool
	FocusPrev() bool
	FocusNext() bool
	IsHidden() bool
	Hide()
	Show()
	IsDisabled() bool
	Disable()
	Enable()

	HandleEnter() bool
	HandleExit() bool
}

type Box interface {
	Scene

	Child(i int) Scene
	Children() []Scene
	AddChild(n Scene)
}
