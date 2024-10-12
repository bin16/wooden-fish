package app

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Stack struct {
	stack    []Scene // current page & overlays
	next     []Scene // for load
	queue    []Scene // for push
	popCount int

	errors []error

	unloading bool
	unloaded  bool

	blank Scene
}

func (u *Stack) Load(p Scene) {
	u.next = append(u.next, p)
	if cnt := len(u.stack); cnt > 0 {
		go u.unloadScene(u.stack[cnt-1])
	}
	go u.loadScene(p)
}

func (u *Stack) Push(p Scene) {
	u.queue = append(u.queue, p)
	go u.loadScene(p)
}

func (u *Stack) Pop() {
	if len(u.stack) <= 1 {
		return
	}

	if u.popCount < 0 {
		u.popCount = 0
	}

	u.popCount += 1
	fmt.Println(u.unloading, u.unloaded)
}

func (u *Stack) List() []Scene {
	return u.stack
}

func (u *Stack) Update() error {

	if len(u.stack) < 1 {
		if p := u.blank; p != nil {
			if err := p.Update(); err != nil {
				return err
			}
		}
	}

	// Update
	for i, n := range u.stack {
		if i > len(u.stack)-2 {
			if err := n.Update(); err != nil {
				return err
			}
		}
	}

	if u.popCount > 0 {
		u.handleUnload()
	} else if len(u.next) > 0 {
		var p = u.next[0]
		if p.Loaded() {
			u.stack = []Scene{p}
			u.next = u.next[1:]
		}
	} else if len(u.queue) > 0 {
		var p = u.queue[0]
		if p.Loaded() {
			u.stack = append(u.stack, p)
			u.queue = u.queue[1:]
		}
	}

	if len(u.errors) > 0 {
		return u.errors[0]
	}

	return nil
}

func (u *Stack) Draw(screen *ebiten.Image) {
	if len(u.stack) < 1 {
		if p := u.blank; p != nil {
			p.Draw(screen)
		}
	}

	for _, p := range u.stack {
		p.Draw(screen)
	}
}

func (u *Stack) Layout(w, h int) (bw, bh int) {
	if p := u.blank; p != nil {
		p.Layout(w, h)
	}

	for _, p := range u.stack {
		p.Layout(w, h)
		p.SetBounds(image.Rect(0, 0, w, h))
	}

	return w, h
}

func (u *Stack) HandleInput() bool {
	if cnt := len(u.stack); cnt > 0 {
		return u.stack[cnt-1].HandleInput()
	}

	return false
}

func (u *Stack) HandleMouseInput() bool {
	if cnt := len(u.stack); cnt > 0 {
		return u.stack[cnt-1].HandleMouseInput()
	}

	return false
}

func (u *Stack) loadScene(p Scene) {
	if err := p.Load(); err != nil {
		u.errors = append(u.errors, err)
	}
}

func (u *Stack) unloadScene(p Scene) {
	u.unloading = true
	u.unloaded = false
	if err := p.Unload(); err != nil {
		u.errors = append(u.errors, err)
	}
	u.unloading = false
	u.unloaded = true
}

func (u *Stack) handleUnload() {
	if u.unloaded {
		u.stack = u.stack[:len(u.stack)-1]
		u.popCount -= 1
		u.unloading = false
		u.unloaded = false
		return
	}

	if u.unloading {
		return
	}

	if u.popCount < 0 {
		u.popCount = 0
		u.unloaded = false
		u.unloading = false
	}

	if u.popCount >= 1 {
		go u.unloadScene(u.stack[len(u.stack)-1])
		return
	}
}
