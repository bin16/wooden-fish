package ui

import (
	"image/color"
	"time"

	"github.com/bin16/wooden-fish/animator/curve"
	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PageStatus uint8

const (
	PageLoading PageStatus = iota
	PageLoaded
	PageUnloading
	PageUnloaded
	// TODO:
	PageAnimIn
	PageAnimInDone
	PageAnimOut
	PageAnimOutDone
)

type Page struct {
	Box

	backgroundColor color.Color

	text    *Text
	onInput []func() bool

	status   PageStatus
	curve    *curve.Curve
	cache    *ebiten.Image
	duration time.Duration
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
	if u.cache == nil || u.cache.Bounds() != screen.Bounds() {
		u.cache = ebiten.NewImage(
			screen.Bounds().Dx(),
			screen.Bounds().Dy(),
		)
	}

	u.cache.Clear()

	if clr := u.backgroundColor; clr != nil {
		u.cache.Fill(clr)
	}

	u.Box.Draw(u.cache)

	var op = &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(
	// 	float64(u.Bounds().Min.X),
	// 	float64(u.Bounds().Min.Y),
	// )

	if u.curve != nil {
		var alpha = curve.Apply(u.curve.Q(), 0.0, 1.0)
		if alpha != 1 {
			op.ColorScale.ScaleAlpha(float32(alpha))
		}
	}

	screen.DrawImage(u.cache, op)
}

func (u *Page) AnimIn() {
	u.curve = curve.NewCurve(
		curve.CurveOpts.AutoPlay(),
		curve.CurveOpts.Duration(u.duration),
		curve.CurveOpts.OnEnd(func(curve *curve.Curve) {
			u.status = PageAnimInDone
		}),
	)

	u.status = PageAnimIn
}

func (u *Page) AnimOut() {
	u.status = PageAnimOut
	u.curve = curve.NewCurve(
		curve.CurveOpts.AutoPlay(),
		curve.CurveOpts.Reversed(),
		curve.CurveOpts.Duration(u.duration),
		curve.CurveOpts.OnEnd(func(curve *curve.Curve) {
			u.status = PageAnimOutDone
		}),
	)
}

func (u *Page) Load() error {
	u.status = PageLoading
	if err := u.Box.Load(); err != nil {
		return err
	}

	u.status = PageLoaded

	u.AnimIn()

	return nil
}

func (u *Page) Unload() error {
	u.status = PageUnloading

	u.AnimOut()
	u.curve.Wait()

	if err := u.Box.Unload(); err != nil {
		return err
	}

	u.status = PageUnloaded
	return nil
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
		duration: time.Second / 4,
		text: NewText(
			TextOpts.Content("功德+1"),
		),
	}
	for _, o := range opts {
		o(page)
	}

	return page
}
