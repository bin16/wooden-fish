package app

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	prevTouchIDs         = []ebiten.TouchID{}
	touchIDs             = []ebiten.TouchID{} // active touch id
	prevTouchInfo        = map[ebiten.TouchID]image.Point{}
	touchInfo            = map[ebiten.TouchID]image.Point{}
	justPressedTouchIDs  = []ebiten.TouchID{} // only new touch id
	justReleasedTouchIDs = []ebiten.TouchID{}
)

func updateTouchState() {
	prevTouchIDs = touchIDs
	prevTouchInfo = touchInfo

	touchIDs = ebiten.AppendTouchIDs(touchIDs[:0])
	touchInfo = make(map[ebiten.TouchID]image.Point)

	justPressedTouchIDs = inpututil.AppendJustPressedTouchIDs(justPressedTouchIDs[:0])
	justReleasedTouchIDs = inpututil.AppendJustReleasedTouchIDs(justReleasedTouchIDs[:0])

	for _, id := range touchIDs {
		touchInfo[id] = image.Pt(ebiten.TouchPosition(id))
	}
}

// tap done and release soon
// ebiten.TouchPosition() return (0, 0) after release,
// so need to keep prev position in prevTouchInfo
func IsTappedInBounds(b image.Rectangle) bool {
	for id, p := range prevTouchInfo {
		if p.In(b) {
			if _, ok := touchInfo[id]; !ok {
				return true
			}
		}
	}

	return false
}
