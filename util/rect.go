package util

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func pixR(d int) []image.Point {
	var pl = [][]image.Point{
		{
			{0, 0},
		},
		{
			{1, 0},
			{0, 1},
		},
		{
			{2, 0}, {2, 1},
			{0, 2}, {1, 2},
		},
		{
			{3, 0}, {3, 1}, {2, 2},
			{0, 3}, {1, 3},
		},
		{
			{4, 0}, {4, 1}, {3, 2},
			{0, 4}, {1, 4}, {2, 3},
		},
		{
			{5, 0}, {5, 1}, {4, 2}, {4, 3},
			{0, 5}, {1, 5}, {2, 4}, {3, 4},
		},
		{
			{6, 0}, {6, 1}, {6, 2}, {5, 3}, {4, 4},
			{0, 6}, {1, 6}, {2, 6}, {3, 5},
		},
		{
			{7, 0}, {7, 1}, {7, 2}, {6, 3}, {6, 4}, {5, 5},
			{0, 7}, {1, 7}, {2, 7}, {3, 6}, {4, 6},
		},
		{
			{8, 0}, {8, 1}, {8, 2}, {7, 3}, {7, 4}, {6, 5},
			{0, 8}, {1, 8}, {2, 8}, {3, 7}, {4, 7}, {5, 6},
		},
		{
			{9, 0}, {9, 1}, {9, 2}, {8, 3}, {8, 4}, {7, 5}, {6, 6},
			{0, 9}, {1, 9}, {2, 9}, {3, 8}, {4, 8}, {5, 7},
		},
		{
			{10, 0}, {10, 1}, {10, 2}, {9, 3}, {9, 4}, {8, 5}, {8, 6}, {7, 7},
			{0, 10}, {1, 10}, {2, 10}, {3, 9}, {4, 9}, {5, 8}, {6, 8},
		},
		{
			{11, 0}, {11, 1}, {11, 2}, {10, 3}, {10, 4}, {9, 5}, {9, 6}, {8, 7},
			{0, 11}, {1, 11}, {2, 11}, {3, 10}, {4, 10}, {5, 9}, {6, 9}, {7, 8},
		},
		{
			{12, 0}, {12, 1}, {12, 2}, {11, 3}, {11, 4}, {11, 5}, {10, 6}, {10, 7}, {9, 8},
			{0, 12}, {1, 12}, {2, 12}, {3, 11}, {4, 11}, {5, 11}, {6, 10}, {7, 10}, {8, 9},
		},
		{
			{13, 0}, {13, 1}, {13, 2}, {12, 3}, {12, 4}, {12, 5}, {11, 6}, {10, 7}, {10, 8}, {9, 9},
			{0, 13}, {1, 13}, {2, 13}, {3, 12}, {4, 12}, {5, 12}, {6, 11}, {7, 10}, {8, 10},
		},
		{
			{14, 0}, {14, 1}, {14, 2}, {13, 3}, {13, 4}, {13, 5}, {12, 6}, {12, 7}, {11, 8}, {10, 9},
			{0, 14}, {1, 14}, {2, 14}, {3, 13}, {4, 13}, {5, 13}, {6, 12}, {7, 12}, {8, 11}, {9, 10},
		},
		{
			{15, 0}, {15, 1}, {15, 2}, {14, 3}, {14, 4}, {14, 5}, {13, 6}, {13, 7}, {12, 8}, {11, 9}, {11, 10},
			{0, 15}, {1, 15}, {2, 15}, {3, 14}, {4, 14}, {5, 14}, {6, 13}, {7, 13}, {8, 12}, {9, 11}, {10, 11},
		},
		{
			{16, 0}, {16, 1}, {16, 2}, {16, 3}, {15, 4}, {15, 5}, {14, 6}, {14, 7}, {13, 8}, {13, 9}, {12, 10}, {11, 11},
			{0, 16}, {1, 16}, {2, 16}, {3, 16}, {4, 15}, {5, 15}, {6, 14}, {7, 14}, {8, 13}, {9, 13}, {10, 12},
		},
	}

	if m := len(pl) - 1; d > m {
		return pl[m]
	}

	return pl[d]
}

func FillRect(screen *ebiten.Image, dr image.Rectangle, clr color.Color, radius ...int) {
	var rad = 0
	if len(radius) > 0 {
		rad = radius[0]
	}
	rad = min(rad, 16)

	var (
		x0 = dr.Min.X
		y0 = dr.Min.Y
		x1 = dr.Max.X
		y1 = dr.Max.Y
	)

	for _, p := range pixR(rad) {
		var r0 = image.Rect(x0+rad-p.X, y0+rad-p.Y, x1-rad+p.X, y0+rad-p.Y+1)
		screen.SubImage(r0).(*ebiten.Image).Fill(clr)
		var r1 = image.Rect(x0+rad-p.X, y1-rad+p.Y-1, x1-rad+p.X, y1-rad+p.Y)
		screen.SubImage(r1).(*ebiten.Image).Fill(clr)
	}

	screen.SubImage(image.Rect(x0, y0+rad, x1, y1-rad)).(*ebiten.Image).Fill(clr)
}

func StrokeRect(screen *ebiten.Image, dr image.Rectangle, clr color.Color, radius ...int) {
	var rad = 0
	if len(radius) > 0 {
		rad = radius[0]
	}
	rad = min(rad, 16)

	var (
		x0 = dr.Min.X
		y0 = dr.Min.Y
		x1 = dr.Max.X - 1
		y1 = dr.Max.Y - 1
	)

	for _, r := range []image.Rectangle{
		image.Rect(x0+rad, y0, x1-rad, y0+1), // T
		image.Rect(x0+rad, y1, x1-rad, y1+1), // B
		image.Rect(x0, y0+rad, x0+1, y1-rad), // L
		image.Rect(x1, y0+rad, x1+1, y1-rad), // R
	} {
		screen.SubImage(r).(*ebiten.Image).Fill(clr)
	}

	for _, p := range pixR(rad) {
		var dl = []image.Point{
			{x0 + rad - p.X, y0 + rad - p.Y}, // L-T
			{x0 + rad - p.X, y1 - rad + p.Y}, // L-B
			{x1 - rad + p.X, y0 + rad - p.Y}, // R-T
			{x1 - rad + p.X, y1 - rad + p.Y}, // R-B
		}

		for _, d := range dl {
			screen.Set(d.X, d.Y, clr)
		}
	}
}

func StrokeCircle(screen *ebiten.Image, dr image.Rectangle, clr color.Color) {
	var (
		d  = min(dr.Dx(), dr.Dy())
		x  = dr.Min.X + (dr.Dx()-d)/2
		y  = dr.Min.Y + (dr.Dy()-d)/2
		cr = image.Rect(x, y, x+d, y+d)
	)

	StrokeRect(screen, cr, clr, d/2)
}

func FillCircle(screen *ebiten.Image, dr image.Rectangle, clr color.Color) {
	var (
		d  = min(dr.Dx(), dr.Dy())
		x  = dr.Min.X + (dr.Dx()-d)/2
		y  = dr.Min.Y + (dr.Dy()-d)/2
		cr = image.Rect(x, y, x+d, y+d)
	)

	FillRect(screen, cr, clr, d/2)
}
