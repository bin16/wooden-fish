package ui

import (
	"image"
	"image/color"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Palette struct {
	Scene
	maxWidth, maxHeight int
	colors              []color.Color
	isVertical          bool
	cache               *ebiten.Image
	bg                  *ebiten.Image
	sr                  []image.Rectangle
	borderColor         color.Color
	gapColor            color.Color
	gap                 int
	border              int
	borderRadius        int
}

func (u *Palette) Layout(ow, oh int) (bw, bh int) {
	var (
		cnt    = len(u.colors)
		uw, uh int
		b      = u.border
		gap    = u.gap
	)
	if len(u.sr) != cnt {
		u.sr = make([]image.Rectangle, cnt)
	}

	if cnt == 0 {
		return u.maxWidth, u.maxHeight
	}

	uw = (u.maxWidth - b*2 - gap*(cnt-1)) / cnt
	uh = (u.maxHeight - b*2)

	for i := range u.colors {
		var (
			x = u.Bounds().Min.X + b + gap*i + uw*i
			y = u.Bounds().Min.Y + b
			r = image.Rect(0, 0, uw, uh).Add(image.Pt(x, y))
		)

		u.sr[i] = r
	}

	bw = uw*cnt + b*2 + gap*(cnt-1)
	bh = u.maxHeight
	return
}

func (u *Palette) Draw(screen *ebiten.Image) {
	if len(u.sr) != len(u.colors) {
		return
	}

	for i, clr := range u.colors {
		util.DrawRect(
			screen,
			u.sr[i],
			util.DrawRectOpts.Fill(clr),
			// util.DrawRectOpts.Color(hexcolor.New("#f00")),
			util.DrawRectOpts.StrokeWidth(0),
		)
	}

	if u.border != 0 {
		util.StrokeRect(screen, u.Bounds(), u.borderColor, u.borderRadius)
	}
}

type PaletteOpt func(pal *Palette)
type PaletteOptions struct{}

func (PaletteOptions) Color(clr color.Color) PaletteOpt {
	return func(pal *Palette) {
		pal.colors = append(pal.colors, clr)
	}
}

func (PaletteOptions) Colors(colors ...color.Color) PaletteOpt {
	return func(pal *Palette) {
		pal.colors = colors
	}
}

func (PaletteOptions) Border(d int) PaletteOpt {
	return func(pal *Palette) {
		pal.border = d
	}
}

func (PaletteOptions) BorderRadius(d int) PaletteOpt {
	return func(pal *Palette) {
		pal.borderRadius = d
	}
}

func (PaletteOptions) BorderColor(clr color.Color) PaletteOpt {
	return func(pal *Palette) {
		pal.borderColor = clr
	}
}

func (PaletteOptions) Gap(d int) PaletteOpt {
	return func(pal *Palette) {
		pal.gap = d
	}
}

func (PaletteOptions) GapColor(clr color.Color) PaletteOpt {
	return func(pal *Palette) {
		pal.gapColor = clr
	}
}

func (PaletteOptions) Width(d int) PaletteOpt {
	return func(pal *Palette) {
		pal.maxWidth = d
	}
}

func (PaletteOptions) Height(d int) PaletteOpt {
	return func(pal *Palette) {
		pal.maxHeight = d
	}
}

var PaletteOpts PaletteOptions

func NewPalette(opts ...PaletteOpt) *Palette {
	var pal = &Palette{
		border:      1,
		borderColor: app.Theme.Color,
		gap:         0,
		maxWidth:    32,
		maxHeight:   12,
	}
	for _, o := range opts {
		o(pal)
	}

	return pal
}
