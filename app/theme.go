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
	// Color:           hexcolor.New("#223843"),
	// SecondaryColor:  hexcolor.New("#4B7D95"),
	// BackgroundColor: hexcolor.New("#DBD3D8"),
	// AccentColor:     hexcolor.New("#D77A61"),

	// Color:           hexcolor.New("#0081A7"),
	// SecondaryColor:  hexcolor.New("#00AFB9"),
	// BackgroundColor: hexcolor.New("#FED9B7"),
	// AccentColor:     hexcolor.New("#F07167"),

	Color:           hexcolor.New("#E7D7C1"),
	BackgroundColor: hexcolor.New("#A78A7F"),
	SecondaryColor:  hexcolor.New("#735751"),
	AccentColor:     hexcolor.New("#BF4342"),
}
