package app

import (
	"image/color"

	"github.com/bin16/go-hexcolor"
)

type theme struct {
	Color           color.Color
	SecondaryColor  color.Color
	BackgroundColor color.Color
	AccentColor     color.Color
}

var Theme = theme{
	Color:           hexcolor.New("#223843"),
	SecondaryColor:  hexcolor.New("#4B7D95"),
	BackgroundColor: hexcolor.New("#DBD3D8"),
	AccentColor:     hexcolor.New("#D77A61"),
}
