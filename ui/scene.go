package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	size     image.Point
	bounds   image.Rectangle
	disabled bool
	hidden   bool
	focused  bool
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

func (u *Scene) HandleTouchInput() bool {
	return false
}

func (u *Scene) HasFocus() bool {
	return u.IsFocused()
}

func (u *Scene) HandleFocus(i int) bool {
	if u.IsFocusable() {
		u.Focus()
		return u.IsFocused()
	}

	return false
}

func (u *Scene) IsFocusable() bool {
	if u.IsHidden() {
		return false
	}

	if u.IsDisabled() {
		return false
	}

	return false
}

func (u *Scene) IsFocused() bool {
	return false
}

func (u *Scene) Focus() {
	u.focused = true
}

func (u *Scene) Blur() {
	u.focused = false
}

func (u *Scene) IsDisabled() bool {
	return u.disabled
}

func (u *Scene) Disable() {
	u.disabled = true
}

func (u *Scene) Enable() {
	u.disabled = false
}

func (u *Scene) SetDisabled(b bool) {
	u.disabled = b
}

func (u *Scene) IsHidden() bool {
	return false
}

func (u *Scene) Show() {
	u.hidden = false
}

func (u *Scene) Hide() {
	u.hidden = true
}

func (u *Scene) SetHidden(b bool) {
	u.hidden = b
}

func (u *Scene) FocusNext() bool {
	return false
}

func (u *Scene) FocusPrev() bool {
	return false
}

func (u *Scene) FocusUp() bool {
	return false
}

func (u *Scene) FocusDown() bool {
	return false
}

func (u *Scene) FocusLeft() bool {
	return false
}

func (u *Scene) FocusRight() bool {
	return false
}

func (u *Scene) HandleUp() bool {
	return false
}

func (u *Scene) HandleDown() bool {
	return false
}

func (u *Scene) HandleLeft() bool {
	return false
}

func (u *Scene) HandleRight() bool {
	return false
}

func (u *Scene) HandlePrev() bool {
	return false
}

func (u *Scene) HandleNext() bool {
	return false
}

func (u *Scene) HandleEnter() bool {
	return false
}

func (u *Scene) HandleExit() bool {
	return false
}

func (u *Scene) Load() error {
	return nil
}

func (u *Scene) Unload() error {
	return nil
}

func (u *Scene) Loaded() bool {
	return true
}
