package game

import (
	"io"

	"github.com/bin16/wooden-fish/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimSound struct {
	FrameIndex int
	Text       string
	Stream     io.Reader
	Bytes      []byte
	Source     string
}

type Anim struct {
	ID         string
	Name       string
	Width      int
	Height     int
	img        *ebiten.Image
	Sounds     []AnimSound
	FreeMode   AnimInfo
	AutoMode   AnimInfo
	RhythmMode AnimInfo
	FPS        int // default
	Text       string
	Source     string
}

type AnimInfo struct {
	FPS int
}

func (u *Anim) Size() (w, h int) {
	return u.Width, u.Height
}

func (u *Anim) Image() *ebiten.Image {
	if u.img == nil {
		u.img = assets.NewImage(u.Source)
	}

	return u.img
}

func NewDefaultAnim() *Anim {
	return &Anim{
		ID:     "id-anim-default",
		Name:   "Default Animation",
		Text:   "Merits: %d",
		Width:  48,
		Height: 48,
		Source: assets.DefaultSheet,
		Sounds: []AnimSound{
			{
				FrameIndex: 5,
				Text:       "Merit+1",
				Source:     "sound.ogg",
			},
		},
		FreeMode: AnimInfo{
			FPS: 30,
		},
		AutoMode: AnimInfo{
			FPS: 9,
		},
		RhythmMode: AnimInfo{
			FPS: 6,
		},
	}
}

func NewDraftAnim() *Anim {
	return &Anim{
		ID:     "id-anim-draft",
		Name:   "Draft Animation",
		Width:  32,
		Height: 32,
		Text:   "Merits: %d",
		Source: assets.DraftSheet,
		Sounds: []AnimSound{
			{
				FrameIndex: 5,
				Text:       "Merit+1",
				Source:     "sound.ogg",
			},
		},
		FreeMode: AnimInfo{
			FPS: 30,
		},
		AutoMode: AnimInfo{
			FPS: 9,
		},
		RhythmMode: AnimInfo{
			FPS: 6,
		},
	}
}

func NewTapTapAnim() *Anim {
	return &Anim{
		ID:     "id-test-anim",
		Name:   "Test Animation",
		Width:  96,
		Height: 96,
		Source: assets.TapTapSheet,
		Text:   "Scores: %d",
		Sounds: []AnimSound{
			{
				FrameIndex: 6,
				Text:       "Ding~",
				Source:     "ding.ogg",
			},
			{
				FrameIndex: 16,
				Text:       "Dong~",
				Source:     "dong.ogg",
			},
		},
		FreeMode: AnimInfo{
			FPS: 15,
		},
		AutoMode: AnimInfo{
			FPS: 9,
		},
		RhythmMode: AnimInfo{
			FPS: 9,
		},
	}
}

var Animation = NewDefaultAnim()

var AnimationOptions = []*Anim{
	NewDefaultAnim(),
	NewDraftAnim(),
	NewTapTapAnim(),
}

func LoadAnim(anim *Anim) {
	for _, item := range AnimationOptions {
		if item.ID == anim.ID {
			Game.AnimID = anim.ID
			Animation = anim
			return
		}
	}

	AnimationOptions = append(AnimationOptions, anim)
	Animation = anim
	Game.AnimID = anim.ID
}

func SetAnim(id string) {
	for _, anim := range AnimationOptions {
		if anim.ID == id {
			Animation = anim
			Game.AnimID = id
			return
		}
	}

	Animation = NewDefaultAnim()
	Game.AnimID = id
}
