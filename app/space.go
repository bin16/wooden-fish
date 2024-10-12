package app

import "image"

type Space struct {
	Top, Right, Bottom, Left int
}

func (sp Space) X() int {
	return sp.Left + sp.Right
}

func (sp Space) Y() int {
	return sp.Top + sp.Bottom
}

func (sp Space) TopLeft() image.Point {
	return image.Pt(sp.Left, sp.Top)
}

func NewSpace(num ...int) Space {
	if len(num) == 0 {
		return Space{}
	}

	if len(num) == 1 {
		return Space{
			num[0],
			num[0],
			num[0],
			num[0],
		}
	}

	if len(num) == 2 {
		return Space{
			num[0],
			num[1],
			num[0],
			num[1],
		}
	}

	if len(num) == 3 {
		return Space{
			num[0],
			num[1],
			num[2],
			num[1],
		}
	}

	return Space{
		num[0],
		num[1],
		num[2],
		num[3],
	}
}
