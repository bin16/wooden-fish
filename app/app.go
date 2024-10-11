package app

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type App struct {
	layers  []Scene
	uiScale float64
	quit    bool
	onInput []func() bool

	drag0 image.Point
	drag  bool
}

func (u *App) Layout(ow, oh int) (bw, bh int) {
	var (
		uiScale = 2
		cw      = ow / uiScale
		ch      = oh / uiScale
	)

	for _, scene := range u.layers {
		scene.Layout(cw, ch)
		scene.SetBounds(image.Rect(0, 0, cw, ch))
	}

	return cw, ch
}

func (u *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for _, n := range u.layers {
		n.Draw(screen)
	}
}

func (u *App) HandleInput() bool {
	for _, fn := range u.onInput {
		if fn() {
			return true
		}
	}

	var cnt = len(u.layers)
	if cnt < 1 {
		return false
	}

	return u.layers[cnt-1].HandleInput()
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
	u.HandleInput()
	u.HandleDrag()

	if u.quit {
		return ebiten.Termination
	}

	var cnt = len(u.layers)
	if cnt > 0 {
		return u.layers[cnt-1].Update()
	}

	return nil
}

func (u *App) Load(scene Scene) {
	if len(u.layers) == 0 {
		u.layers = append(u.layers, scene)
	} else {
		u.layers[len(u.layers)-1] = scene
	}
}

func (u *App) Quit() {
	u.quit = true
}
