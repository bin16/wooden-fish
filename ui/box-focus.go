package ui

import (
	"github.com/bin16/wooden-fish/util"
)

func (u *Box) HasFocus() bool {
	for _, n := range u.Children() {
		if n.HasFocus() {
			return true
		}
	}

	return false
}

func (u *Box) Focus() {
	if u.HasFocus() {
		return
	}

	for _, n := range u.Children() {
		if n.FocusNext() {
			return
		}
	}
}

func (u *Box) Blur() {
	for _, n := range u.Children() {
		n.Blur()
	}
}

func (u *Box) FocusNext() bool {
	var (
		k   = -1
		cnt = len(u.children)
	)
	for i, n := range u.children {
		if n.HasFocus() {
			k = i
			if n.FocusNext() {
				return true
			}

			break
		}
	}

	var queue = []int{}
	queue = append(queue, util.Range(k+1, cnt)...)
	if u.loopSearch {
		queue = append(queue, util.Range(0, k)...) // loop search
	}

	for _, i := range queue {
		var n = u.Child(i)
		if n.FocusNext() {
			if p := u.Child(k); p != nil {
				p.Blur()
			}
			return true
		}
	}

	return false
}

func (u *Box) FocusPrev() bool {
	var (
		cnt = len(u.children)
		k   = cnt
	)
	for i := range u.children {
		var n = u.Child(cnt - 1 - i)
		if n.HasFocus() {
			k = i
			if n.FocusPrev() {
				return true
			}

			// TODO: blur?
			break
		}
	}

	var queue = []int{}
	queue = append(queue, util.Range(k-1, -1, -1)...)
	queue = append(queue, util.Range(cnt-1, k)...) // loop search

	for _, i := range queue {
		var n = u.Child(i)
		if n.FocusPrev() {
			if p := u.Child(k); p != nil {
				p.Blur()
			}
			return true
		}
	}

	return false
}

func (u *Box) FocusLeft() bool {
	var (
		cnt = len(u.children)
		k   = cnt
	)
	for i := range u.children {
		var n = u.Child(cnt - 1 - i)
		if n.HasFocus() {
			k = i
			if n.FocusLeft() {
				return true
			}

			// TODO: blur?
			break
		}
	}

	var queue = []int{}
	queue = append(queue, util.Range(k-1, -1, -1)...)
	queue = append(queue, util.Range(cnt-1, k)...) // loop search

	for _, i := range queue {
		var n = u.Child(i)
		if n.FocusLeft() {
			if p := u.Child(k); p != nil {
				p.Blur()
			}
			return true
		}
	}

	return false
}

func (u *Box) FocusUp() bool {
	var (
		cnt = len(u.children)
		k   = cnt
	)
	for i := range u.children {
		var n = u.Child(cnt - 1 - i)
		if n.HasFocus() {
			k = i
			if n.FocusUp() {
				return true
			}

			// TODO: blur?
			break
		}
	}

	var queue = []int{}
	queue = append(queue, util.Range(k-1, -1, -1)...)
	queue = append(queue, util.Range(cnt-1, k)...) // loop search

	for _, i := range queue {
		var n = u.Child(i)
		if n.FocusUp() {
			if p := u.Child(k); p != nil {
				p.Blur()
			}
			return true
		}
	}

	return false
}

func (u *Box) FocusRight() bool {
	var (
		k   = -1
		cnt = len(u.children)
	)
	for i, n := range u.children {
		if n.HasFocus() {
			k = i
			if n.FocusRight() {
				return true
			}

			break
		}
	}

	var queue = []int{}
	queue = append(queue, util.Range(k+1, cnt)...)
	if u.loopSearch {
		queue = append(queue, util.Range(0, k)...) // loop search
	}

	for _, i := range queue {
		var n = u.Child(i)
		if n.FocusRight() {
			if p := u.Child(k); p != nil {
				p.Blur()
			}
			return true
		}
	}

	return false
}

func (u *Box) FocusDown() bool {
	var (
		k   = -1
		cnt = len(u.children)
	)
	for i, n := range u.children {
		if n.HasFocus() {
			k = i
			if n.FocusDown() {
				return true
			}

			break
		}
	}

	var queue = []int{}
	queue = append(queue, util.Range(k+1, cnt)...)
	if u.loopSearch {
		queue = append(queue, util.Range(0, k)...) // loop search
	}

	for _, i := range queue {
		var n = u.Child(i)
		if n.FocusDown() {
			if p := u.Child(k); p != nil {
				p.Blur()
			}
			return true
		}
	}

	return false
}

func (u *Box) HandleEnter() bool {
	for _, n := range u.Children() {
		if n.HasFocus() {
			if n.HandleEnter() {
				return true
			}
		}
	}

	return false
}

func (u *Box) HandleExit() bool {
	for _, n := range u.Children() {
		if n.HasFocus() {
			if n.HandleExit() {
				return true
			}
		}
	}

	return false
}
