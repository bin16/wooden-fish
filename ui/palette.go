package ui

import (
	"fmt"
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
	maxColumns          int
	sr                  []image.Rectangle
	borderColor         color.Color
	gapColor            color.Color
	gap                 int
	border              int
	borderRadius        int
	r0                  image.Rectangle
}

func (u *Palette) Layout(ow, oh int) (bw, bh int) {
	var (
		cnt     = len(u.colors)
		uw, uh  int
		b       = u.border
		gap     = u.gap
		columns = util.OR(u.maxColumns > 0, u.maxColumns, cnt)
		rows    = cnt/columns + util.OR(cnt%columns > 0, 1, 0)
	)
	if len(u.sr) != cnt {
		u.sr = make([]image.Rectangle, cnt)
	}

	if cnt == 0 {
		return u.maxWidth, u.maxHeight
	}

	uw = (u.maxWidth - b*2 - gap*(columns-1)) / columns
	uh = (u.maxHeight - b*2 - gap*(rows-1)) / rows

	for i := range u.colors {
		var (
			col = i % columns
			row = i / columns
			x   = b + gap*col + uw*col
			y   = b + gap*row + uh*row
			r   = image.Rect(0, 0, uw, uh).Add(image.Pt(x, y))
		)

		u.sr[i] = r
	}

	bw = uw*columns + b*2 + gap*(columns-1)
	bh = uh*rows + b*2 + gap*(rows-1)
	return
}

func (u *Palette) Draw(screen *ebiten.Image) {
	if u.bounds.Empty() || len(u.sr) != len(u.colors) {
		return
	}

	for i, clr := range u.colors {
		util.DrawRect(
			screen,
			u.sr[i].Add(u.bounds.Min),
			util.DrawRectOpts.Fill(clr),
			util.DrawRectOpts.StrokeWidth(0),
		)
	}

	if u.border != 0 {
		util.StrokeRect(screen, u.Bounds(), u.borderColor, u.borderRadius)
	}
}

func (u *Palette) SetBounds(r image.Rectangle) {
	if r0 := u.bounds; r != r0 {
		fmt.Println(r, "<--", r0)
	}

	u.Scene.SetBounds(r)
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

func (PaletteOptions) Columns(d int) PaletteOpt {
	return func(pal *Palette) {
		pal.maxColumns = d
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
		maxColumns:  2,
	}
	for _, o := range opts {
		o(pal)
	}

	return pal
}
