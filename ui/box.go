package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/hajimehoshi/ebiten/v2"
)

// Box is basic Container
type Box struct {
	Scene
	children   []app.Scene
	cr         image.Rectangle
	sr         []image.Rectangle
	loopSearch bool
}

func (u *Box) HandleInput() bool {
	for _, n := range u.children {
		if n.HandleInput() {
			return true
		}
	}

	return false
}

func (u *Box) HandleMouseInput() bool {
	for _, n := range u.children {
		if n.HandleMouseInput() {
			return true
		}
	}

	return false
}

func (u *Box) HandleTouchInput() bool {
	for _, n := range u.children {
		if n.HandleTouchInput() {
			return true
		}
	}

	return false
}

func (u *Box) AddChild(n app.Scene) {
	u.children = append(u.children, n)
}

func (u *Box) Child(i int) app.Scene {
	if i < 0 {
		return nil
	}

	if i > len(u.children)-1 {
		return nil
	}

	return u.children[i]
}

func (u *Box) Children() []app.Scene {
	return u.children
}

func (u *Box) Update() error {
	for _, n := range u.children {
		if err := n.Update(); err != nil {
			return err
		}
	}

	return nil
}

var debug = true

func (u *Box) Draw(screen *ebiten.Image) {
	if u.bounds.Empty() {
		return
	}

	u.Scene.Draw(screen)
	for _, n := range u.children {
		n.Draw(screen)
	}
}

func (u *Box) Layout(ow, oh int) (bw, bh int) {
	if d := len(u.children); len(u.sr) != d {
		u.sr = make([]image.Rectangle, d)
	}
	for i, n := range u.children {
		var w, h = n.Layout(ow, oh)
		bw = max(bw, w)
		bh = max(bh, h)

		u.sr[i] = image.Rect(0, 0, w, h)
	}

	return
}

func (u *Box) SetBounds(r image.Rectangle) {
	u.Scene.SetBounds(r)
	for _, n := range u.children {
		// var (
		// 	x = (r.Dx() - u.sr[i].Dx()) / 2
		// 	y = (r.Dy() - u.sr[i].Dy()) / 2
		// )

		// n.SetBounds(u.sr[i].Add(image.Pt(x, y)).Add(r.Min))

		n.SetBounds(r)
	}
}

type BoxOpt func(box *Box)

type BoxOptions struct{}

var BoxOpts BoxOptions

func (BoxOptions) Contents(items ...app.Scene) BoxOpt {
	return func(box *Box) {
		box.children = append(box.children, items...)
	}
}

func (BoxOptions) LoopSearch(b bool) BoxOpt {
	return func(box *Box) {
		box.loopSearch = b
	}
}

func NewBox(opts ...BoxOpt) *Box {
	var box = &Box{}
	for _, o := range opts {
		o(box)
	}

	return box
}
