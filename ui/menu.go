package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Menu struct {
	VBox

	activeIndex int

	onEnter map[int][]func()
	onExit  []func()
}

func (u *Menu) HandleMouseInput() bool {
	var cursor = image.Pt(ebiten.CursorPosition())
	if !cursor.In(u.Bounds()) {
		return false
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		if u.Exit() {
			return true
		}
	}

	for i, n := range u.Children() {

		if cursor.In(n.Bounds()) {
			if u.HandleFocus(i) {
				ebiten.SetCursorShape(ebiten.CursorShapePointer)

				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
					if u.Enter() {
						return true
					}
				}

				return true
			}
		}

	}

	return false
}

func (u *Menu) HandleFocus(i int) bool {
	if i < 0 {
		return false
	}

	if i > len(u.children)-1 {
		return false
	}

	u.activeIndex = i
	return true
}

func (u *Menu) HasFocus() bool {
	if u.activeIndex < 0 {
		return false
	}

	if u.activeIndex > len(u.children)-1 {
		return false
	}

	return true
}

func (u *Menu) FocusUp() bool {
	return u.HandleFocus(u.activeIndex - 1)
}

func (u *Menu) FocusDown() bool {
	return u.HandleFocus(u.activeIndex + 1)
}

func (u *Menu) FocusPrev() bool {
	return u.FocusUp()
}

func (u *Menu) FocusNext() bool {
	return u.FocusDown()
}

func (u *Menu) FocusLeft() bool {
	return false
}

func (u *Menu) FocusRight() bool {
	return u.Enter()
}

func (u *Menu) HandleEnter() bool {
	return u.Enter()
}

func (u *Menu) HandleExit() bool {
	return u.Exit()
}

func (u *Menu) Exit() bool {
	if len(u.onExit) > 0 {
		for _, fn := range u.onExit {
			fn()
		}

		return true
	}

	return false
}

func (u *Menu) Enter() bool {
	var k = u.activeIndex
	if k < 0 {
		return false
	}

	if k > len(u.children)-1 {
		return false
	}

	if onEnter, ok := u.onEnter[k]; ok {
		if len(onEnter) > 0 {
			for _, fn := range onEnter {
				fn()
			}

			return true
		}
	}

	return true
}

func (u *Menu) OnEnter(i int, fn func()) {
	if u.onEnter == nil {
		u.onEnter = make(map[int][]func())
	}

	u.onEnter[i] = append(u.onEnter[i], fn)
}

func (u *Menu) OnExit(fn func()) {
	u.onExit = append(u.onExit, fn)
}

func (u *Menu) Draw(screen *ebiten.Image) {
	for i, n := range u.children {
		if u.activeIndex == i {
			util.DrawRect(
				screen,
				n.Bounds(),
				util.DrawRectOpts.StrokeWidth(1),
				util.DrawRectOpts.Color(app.Theme.SecondaryColor),
				util.DrawRectOpts.Radius(2),
			)
		}

		n.Draw(screen)
	}
}

type MenuOpt func(menu *Menu)
type MenuOptions struct{}

var MenuOpts MenuOptions

func (MenuOptions) TextItem(s string, onEnter func()) MenuOpt {
	var item = NewText(
		TextOpts.Padding(1, 4),
		TextOpts.Content(s),
	)

	return func(menu *Menu) {
		var k = len(menu.children)
		menu.AddChild(item)
		menu.OnEnter(k, onEnter)
	}
}

func (MenuOptions) Item(item app.Scene, onEnter func()) MenuOpt {
	return func(menu *Menu) {
		var k = len(menu.children)
		menu.AddChild(item)
		menu.OnEnter(k, onEnter)
	}
}

func (MenuOptions) OnExit(onExit func()) MenuOpt {
	return func(menu *Menu) {
		menu.OnExit(onExit)
	}
}

func NewMenu(opts ...MenuOpt) *Menu {
	var menu = &Menu{}
	for _, o := range opts {
		o(menu)
	}

	return menu
}
