package ui

import (
	"image/color"

	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Page struct {
	Box

	backgroundColor color.Color

	text    *Text
	onInput []func() bool
}

func (u *Page) HandleInput() bool {
	for _, fn := range u.onInput {
		if fn() {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if u.HandleEnter() {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if u.HandleExit() {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		if u.FocusUp() {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		if u.FocusDown() {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		if u.FocusLeft() {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		if u.FocusRight() {
			return true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			if u.FocusPrev() {
				return true
			}
		} else {
			if u.FocusNext() {
				return true
			}
		}
	}

	return u.Box.HandleInput()
}

func (u *Page) Draw(screen *ebiten.Image) {
	// util.DrawRect(
	// 	screen,
	// 	u.bounds,
	// 	util.DrawRectOpts.StrokeWidth(1),
	// 	util.DrawRectOpts.Color(hexcolor.New("#F6F7EB")),
	// 	util.DrawRectOpts.Radius(2),
	// 	util.DrawRectOpts.Fill(hexcolor.New("#5C415D")),
	// )

	if clr := u.backgroundColor; clr != nil {
		screen.SubImage(u.Bounds()).(*ebiten.Image).Fill(clr)
		// screen.Fill(color.Black)
	}

	u.Box.Draw(screen)
}

type PageOpt func(page *Page)
type PageOptions struct{}

var PageOpts PageOptions

func (PageOptions) Contents(items ...app.Scene) PageOpt {
	return func(page *Page) {
		page.children = append(page.children, items...)
	}
}

func (PageOptions) OnInput(fn func() bool) PageOpt {
	return func(page *Page) {
		page.onInput = append(page.onInput, fn)
	}
}

func (PageOptions) Fill(clr color.Color) PageOpt {
	return func(page *Page) {
		page.backgroundColor = clr
	}
}

func NewPage(opts ...PageOpt) *Page {
	var page = &Page{
		text: NewText(
			TextOpts.Content("功德+1"),
		),
	}
	for _, o := range opts {
		o(page)
	}

	return page
}
