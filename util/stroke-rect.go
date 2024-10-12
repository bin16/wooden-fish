package util

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func StrokeRect(screen *ebiten.Image, dr image.Rectangle, clr color.Color, radius ...int) {
	var dl = [][]int{
		{},
		{1},
		{2, 1},
		{3, 2, 1},
		{4, 3, 2, 1},
		{5, 3, 2, 1, 1},
	}

	var rad = 0
	if len(radius) > 0 {
		rad = radius[0]
	}
	if max := len(dl) - 1; rad > max {
		rad = max
	}

	var rl = []image.Rectangle{
		image.Rect(0, rad, 1, dr.Dy()-rad),               // left
		image.Rect(dr.Dx()-1, rad, dr.Dx(), dr.Dy()-rad), // right
		image.Rect(rad, 0, dr.Dx()-rad, 1),               // top
		image.Rect(rad, dr.Dy()-1, dr.Dx()-rad, dr.Dy()), // bottom
	}
	for _, r := range rl {
		screen.SubImage(r.Add(dr.Min)).(*ebiten.Image).Fill(clr)
	}

	var p0 = 0
	for i, p1 := range dl[rad] {
		var offset = []int{p1}
		if p1 < p0 {
			offset = Range(p1, p0)
		}

		for _, d := range offset {
			var (
				x0 = dr.Min.X + d
				x1 = dr.Max.X - d - 1
				y0 = dr.Min.Y + i
				y1 = dr.Max.Y - i - 1
			)
			screen.Set(x0, y0, clr) // top left
			screen.Set(x1, y0, clr) // top right
			screen.Set(x0, y1, clr) // bottom left
			screen.Set(x1, y1, clr) // bottom right
		}

		p0 = p1
	}
}
