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

func (u *Menu) HandleTouchInput() bool {
	for i, n := range u.children {

		if app.IsTappedInBounds(n.Bounds()) {
			if n.IsDisabled() {
				return false
			}

			if u.HandleFocus(i) {
				if u.Enter() {
					return true
				}
			}
		}

	}

	return false
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

			if n.IsDisabled() {
				ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
				return false
			}

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

func (u *Menu) HandleFocus(k int) bool {
	if k < 0 {
		return false
	}

	if k > len(u.children)-1 {
		return false
	}

	if u.children[k].IsDisabled() {
		return false
	}

	u.activeIndex = k
	for i, n := range u.Children() {
		if i == k {
			n.Focus()
		} else {
			n.Blur()
		}
	}

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
	var (
		k   = u.activeIndex
		cnt = len(u.children)
	)
	if !u.HasFocus() {
		k = cnt
	}

	var queue = []int{}
	queue = append(queue, util.Range(k-1, -1)...)
	if u.loopSearch {
		queue = append(queue, util.Range(cnt-1, k)...) // loop search
	}

	for _, k1 := range queue {
		if u.HandleFocus(k1) {
			return true
		}
	}

	u.Blur()
	return false
}

func (u *Menu) FocusDown() bool {
	var (
		k   = u.activeIndex
		cnt = len(u.children)
	)
	if !u.HasFocus() {
		k = -1
	}

	var queue = []int{}
	queue = append(queue, util.Range(k+1, cnt)...)
	if u.loopSearch {
		queue = append(queue, util.Range(0, k)...) // loop search
	}

	for _, k1 := range queue {
		if u.HandleFocus(k1) {
			return true
		}
	}

	u.Blur()
	return false
}

func (u *Menu) Blur() {
	u.activeIndex = -1
	u.Box.Blur()
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
	if u.bounds.Empty() {
		return
	}

	for i, n := range u.children {
		if u.activeIndex == i {
			util.DrawRect(
				screen,
				n.Bounds(),
				util.DrawRectOpts.StrokeWidth(1),
				util.DrawRectOpts.Color(app.Theme.AccentColor),
				util.DrawRectOpts.Radius(3),
			)
		}

		n.Draw(screen)
	}
}

type MenuOpt func(menu *Menu)
type MenuOptions struct{}

var MenuOpts MenuOptions

func (MenuOptions) TextItem(s string, onEnter func(), opts ...TextOpt) MenuOpt {
	var item = NewText(
		TextOpts.Padding(1, 4),
		TextOpts.Content(s),
	)
	for _, o := range opts {
		o(item)
	}

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

func (MenuOptions) LoopSearch(b bool) MenuOpt {
	return func(menu *Menu) {
		menu.loopSearch = b
	}
}

func NewMenu(opts ...MenuOpt) *Menu {
	var menu = &Menu{}
	menu.loopSearch = true
	for _, o := range opts {
		o(menu)
	}

	return menu
}
