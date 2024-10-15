package app

import (
	"encoding/json"
	"image/color"

	"github.com/bin16/go-hexcolor"
)

var Theme theme = *NewDefaultTheme()
var ThemeOptions []*theme = []*theme{
	NewDefaultTheme(),
	NewHighContrastTheme(),
	NewDarkHighContrastTheme(),
	NewTheme(
		ThemeOpts.Name("Midnight Theme"),
		ThemeOpts.ID("id-theme-midnight"),
		ThemeOpts.Color(hexcolor.New("#E9A6A6")),
		ThemeOpts.BackgroundColor(hexcolor.New("#1F1D36")),
		ThemeOpts.AccentColor(hexcolor.New("#e63946")),
		ThemeOpts.SecondaryColor(hexcolor.New("#864879")),
	),
	NewTheme(
		ThemeOpts.Name("True Blue Theme"),
		ThemeOpts.ID("id-theme-true-blue"),
		ThemeOpts.Color(hexcolor.New("#F4FEC1")),
		ThemeOpts.BackgroundColor(hexcolor.New("#4464AD")),
		ThemeOpts.AccentColor(hexcolor.New("#F58F29")),
		ThemeOpts.SecondaryColor(hexcolor.New("#A4B0F5")),
	),
	NewTheme(
		ThemeOpts.Name("Dark Purple Theme"),
		ThemeOpts.ID("id-theme-dark-purple"),
		ThemeOpts.Color(hexcolor.New("#FFE1FF")),
		ThemeOpts.BackgroundColor(hexcolor.New("#433878")),
		ThemeOpts.SecondaryColor(hexcolor.New("#E4B1F0")),
		ThemeOpts.AccentColor(hexcolor.New("#E7CCCC")),
	),
}

type theme struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Color           color.Color `json:"color"`
	SecondaryColor  color.Color `json:"secondaryColor"`
	BackgroundColor color.Color `json:"backgroundColor"`
	AccentColor     color.Color `json:"accentColor"`
}

type ThemeData struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Color           string `json:"color"`
	SecondaryColor  string `json:"secondaryColor"`
	BackgroundColor string `json:"backgroundColor"`
	AccentColor     string `json:"accentColor"`
}

type ThemeOpt func(u *theme)
type themeOptions struct{}

func (themeOptions) ID(s string) ThemeOpt {
	return func(u *theme) {
		u.ID = s
	}
}

func (themeOptions) Name(s string) ThemeOpt {
	return func(u *theme) {
		u.Name = s
	}
}

func (themeOptions) Color(clr hexcolor.Color) ThemeOpt {
	return func(u *theme) {
		u.Color = clr
	}
}

func (themeOptions) BackgroundColor(clr hexcolor.Color) ThemeOpt {
	return func(u *theme) {
		u.BackgroundColor = clr
	}
}

func (themeOptions) SecondaryColor(clr hexcolor.Color) ThemeOpt {
	return func(u *theme) {
		u.SecondaryColor = clr
	}
}

func (themeOptions) AccentColor(clr hexcolor.Color) ThemeOpt {
	return func(u *theme) {
		u.AccentColor = clr
	}
}

func (themeOptions) Import(raw *ThemeData) ThemeOpt {
	return func(u *theme) {
		if raw.ID != "" {
			u.ID = raw.ID
		}

		if raw.Name != "" {
			u.Name = raw.Name
		}

		u.Color = hexcolor.New(raw.Color)
		u.SecondaryColor = hexcolor.New(raw.SecondaryColor)
		u.BackgroundColor = hexcolor.New(raw.BackgroundColor)
		u.AccentColor = hexcolor.New(raw.AccentColor)
	}
}

var ThemeOpts themeOptions

func LoadTheme(th *theme) {
	for _, item := range ThemeOptions {
		if item.ID == th.ID {
			Theme = *th
			return
		}
	}

	ThemeOptions = append(ThemeOptions, th)
	Theme = *th
}

func SetTheme(id string) {
	for _, item := range ThemeOptions {
		if item.ID == id {
			Theme = *item
			return
		}
	}

	Theme = *NewDefaultTheme()
}

func NewTheme(opts ...ThemeOpt) *theme {
	var th = NewDefaultTheme()
	for _, o := range opts {
		o(th)
	}

	return th
}

func RegisterTheme(th *theme) {
	ThemeOptions = append(ThemeOptions, th)
}

func NewThemeFromBytes(data []byte) (*theme, error) {
	var raw = &ThemeData{}
	err := json.Unmarshal(data, raw)

	var th = &theme{}

	th.ID = raw.ID
	th.Name = raw.Name
	th.Color = hexcolor.New(raw.Color)
	th.SecondaryColor = hexcolor.New(raw.SecondaryColor)
	th.BackgroundColor = hexcolor.New(raw.BackgroundColor)
	th.AccentColor = hexcolor.New(raw.AccentColor)

	return th, err
}

func NewDefaultTheme() *theme {
	return &theme{
		ID:              "id-theme-default",
		Name:            "Default Theme",
		Color:           hexcolor.New("#735751"),
		BackgroundColor: hexcolor.New("#E7D7C1"),
		SecondaryColor:  hexcolor.New("#A78A7F"),
		AccentColor:     hexcolor.New("#BF4342"),
	}
}

func NewHighContrastTheme() *theme {
	return &theme{
		ID:              "id-theme-high-contrast-theme",
		Name:            "High Contrast Theme",
		Color:           hexcolor.New("#000"),
		BackgroundColor: hexcolor.New("#fff"),
		SecondaryColor:  hexcolor.New("#666"),
		AccentColor:     hexcolor.New("#f00"),
	}
}

func NewDarkHighContrastTheme() *theme {
	return &theme{
		ID:              "id-theme-high-contrast-theme--dark",
		Name:            "High Contrast Theme (Dark)",
		Color:           hexcolor.New("#fff"),
		BackgroundColor: hexcolor.New("#000"),
		SecondaryColor:  hexcolor.New("#999"),
		AccentColor:     hexcolor.New("#ff0"),
	}
}
