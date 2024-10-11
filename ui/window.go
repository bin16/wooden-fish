package ui

import (
	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	Box

	text *Text
}

func (u *Window) Draw(screen *ebiten.Image) {
	// util.DrawRect(
	// 	screen,
	// 	u.bounds,
	// 	util.DrawRectOpts.StrokeWidth(1),
	// 	util.DrawRectOpts.Color(hexcolor.New("#F6F7EB")),
	// 	util.DrawRectOpts.Radius(2),
	// 	util.DrawRectOpts.Fill(hexcolor.New("#5C415D")),
	// )

	u.Box.Draw(screen)
}

type WinOpt func(win *Window)
type WinOptions struct{}

var WinOpts WinOptions

func (WinOptions) Contents(items ...app.Scene) WinOpt {
	return func(win *Window) {
		win.children = append(win.children, items...)
	}
}

func NewWindow(opts ...WinOpt) *Window {
	var win = &Window{
		text: NewText(
			TextOpts.Content("功德+1"),
		),
	}
	for _, o := range opts {
		o(win)
	}

	return win
}
