package assets

import (
	_ "embed"
)

var (
	//go:embed "wooden-fish-icon-8x8.png"
	icon_8x8_bytes []byte

	//go:embed "wooden-fish-icon-16x16.png"
	icon_16x16_bytes []byte

	//go:embed "wooden-fish-icon-32x32.png"
	icon_32x32_bytes []byte

	//go:embed "sound.ogg"
	sound_bytes []byte

	//go:embed "wooden-fish-48x48-sheet.png"
	sheet_bytes []byte

	//go:embed "wooden-fish-32x32-anim.png"
	draft_sheet_bytes []byte

	//go:embed "tap-tap-96x96-sheet.png"
	tap_tap_sheet_bytes []byte

	//go:embed "ding.ogg"
	ding_bytes []byte

	//go:embed "dong.ogg"
	dong_bytes []byte
)

var (
	DefaultSound = "sound.ogg"
	DingSound    = "ding.ogg"
	DongSound    = "dong.ogg"

	Icon8x8   = "wooden-fish-icon-8x8.png"
	Icon16x16 = "wooden-fish-icon-16x16.png"
	Icon32x32 = "wooden-fish-icon-32x32.png"

	DefaultSheet = "wooden-fish-48x48-sheet.png"
	DraftSheet   = "wooden-fish-32x32-anim.png"
	TapTapSheet  = "tap-tap-96x96-sheet.png"
)

func init() {
	Set(DefaultSound, sound_bytes)
	Set(DingSound, ding_bytes)
	Set(DongSound, dong_bytes)

	Set(Icon8x8, icon_8x8_bytes)
	Set(Icon16x16, icon_16x16_bytes)
	Set(Icon32x32, icon_32x32_bytes)

	Set(DefaultSheet, sheet_bytes)
	Set(DraftSheet, draft_sheet_bytes)
	Set(TapTapSheet, tap_tap_sheet_bytes)
}
