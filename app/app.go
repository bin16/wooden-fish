package app

import (
	"errors"
	"image"
	"math"

	"github.com/bin16/wooden-fish/animator/curve"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type App struct {
	stack Stack

	uiScale float64
	onInput []func() bool

	drag0 image.Point
	drag  bool

	errors []error
}

func (app *App) Layout(ow, oh int) (bw, bh int) {
	var w, h = app.CanvasSize(ow, oh)

	return app.stack.Layout(w, h)
}

func (app *App) CanvasSize(ow, oh int) (w, h int) {
	var scale = app.uiScale
	w = int(math.Round(float64(ow) / scale))
	h = int(math.Round(float64(oh) / scale))
	return
}

func (u *App) Draw(screen *ebiten.Image) {
	if clr := Theme.BackgroundColor; clr != nil {
		screen.Fill(clr)
	}

	u.stack.Draw(screen)
}

func (u *App) HandleDrag() bool {
	var p = image.Pt(ebiten.CursorPosition())

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		u.drag = true
		u.drag0 = p
		return true
	}

	if u.drag {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			u.drag = false
			return true
		}

		var np = p.Sub(u.drag0).Add(image.Pt(ebiten.WindowPosition()))
		ebiten.SetWindowPosition(np.X, np.Y)

		return true
	}

	return false
}

func (u *App) Update() error {
	for _, err := range app.errors {
		if errors.Is(err, ebiten.Termination) {
			return ebiten.Termination
		}
	}

	if err := curve.Update(); err != nil {
		return err
	}

	u.stack.HandleInput()
	u.stack.HandleMouseInput()
	u.HandleDrag()
	if err := app.stack.Update(); err != nil {
		return err
	}

	return nil
}

func (app *App) Push(p Scene) {
	app.stack.Push(p)
}

func (app *App) Pop() {
	app.stack.Pop()
}

func (app *App) Load(p Scene) {
	app.stack.Load(p)
}

func (app *App) Preload(p Scene) {
	go app.loadScene(p)
}

func (app *App) loadScene(p Scene) {
	if err := p.Load(); err != nil {
		app.errors = append(app.errors, err)
	}
}

func (u *App) Quit() {
	u.errors = append(u.errors, ebiten.Termination)
}
