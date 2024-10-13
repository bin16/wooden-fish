package assets

import (
	"bytes"
	_ "embed"

	"github.com/bin16/wooden-fish/util"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
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
)

var (
	Icon_8x8              = util.NewImageFromBytes(icon_8x8_bytes)
	Icon_16x16            = util.NewImageFromBytes(icon_16x16_bytes)
	Icon_32x32            = util.NewImageFromBytes(icon_32x32_bytes)
	DefaultSoundBytes     = sound_bytes
	DefaultSoundStream, _ = vorbis.DecodeF32(bytes.NewReader(sound_bytes))
	DefaultAnimSheetBytes = sheet_bytes
	DefaultAnimSheet      = util.NewImageFromBytes(sheet_bytes)
)
