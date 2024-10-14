package assets

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"strings"

	"github.com/bin16/go-hexcolor"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

var cacheDB = map[string][]byte{}

func Get(name string) []byte {
	return cacheDB[name]
}

func Set(name string, data []byte) {
	cacheDB[name] = data
}

func Has(name string) bool {
	_, ok := cacheDB[name]
	return ok
}

func NewAudioPlayer(name string) *audio.Player {
	var data = cacheDB[name]
	if strings.HasSuffix(name, ".ogg") {
		var s, _ = vorbis.DecodeF32(bytes.NewReader(data))
		var ply, _ = audio.CurrentContext().NewPlayerF32(s)

		return ply
	}

	log.Panicf("NewAudioPlayer() unhandled file type: %s\n", name)
	return nil
}

func NewImage(name string) *ebiten.Image {
	var data = cacheDB[name]
	if strings.HasSuffix(name, ".png") {
		img, err := decodePNG(data)
		if err != nil {
			return defaultImage()
		}

		return ebiten.NewImageFromImage(img)
	}

	log.Panicf("NewImage() unhandled file type: %s\n", name)
	return defaultImage()
}

func decodePNG(data []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(data))
}

func defaultImage() *ebiten.Image {
	var img = ebiten.NewImage(1, 1)
	img.Fill(hexcolor.New("#c0c"))

	return img
}

func init() {
	audio.NewContext(48000)
}
