package curve

import (
	"math"
	"time"
)

type EasingFunc func(x float64) (y float64)
type EventFunc func(curve *Curve)

type eventManager struct {
	handlers []EventFunc
}

func (u *eventManager) On(fn EventFunc) {
	u.handlers = append(u.handlers, fn)
}

func (u *eventManager) Emit(curve *Curve) {
	for _, fn := range u.handlers {
		fn(curve)
	}
}

type Curve struct {
	easingFunc EasingFunc
	tps        float64
	duration   time.Duration
	frameIndex int
	id         ID

	loop     bool
	swing    bool
	reversed bool
	playing  bool
	end      bool
	del      bool

	END eventManager
	DEL eventManager

	done chan bool
}

func (u *Curve) Tick() {
	if !u.playing {
		return
	}

	if u.reversed {
		u.frameIndex -= 1
	} else {
		u.frameIndex += 1
	}

	if u.frameIndex < 0 {
		if u.reversed {
			if u.swing {
				u.frameIndex = 1
				return
			}

			// TODO: end
			u.frameIndex = 0
			u.end = true
			u.playing = false
			if u.done != nil {
				u.done <- true
			}
			return
		}
	}

	var cnt = u.FrameCount()
	if u.frameIndex > cnt-1 {
		if u.swing {
			u.frameIndex = cnt - 2
			return
		}

		if u.loop {
			u.frameIndex = 0
			return
		}

		u.frameIndex = cnt - 1
		u.end = true
		u.playing = false
		if u.done != nil {
			u.done <- true
		}
		return
	}
}

func (u *Curve) curveInit() {
	u.frameIndex = 0
	if u.reversed {
		u.frameIndex = u.FrameCount() - 1
	}
}

func (u *Curve) curveUpdate() {
	if !u.playing {
		return
	}

	if u.reversed {
		u.frameIndex -= 1
	} else {
		u.frameIndex += 1
	}

	if u.frameIndex < 0 {
		if u.reversed {
			if u.swing {
				u.frameIndex = 1
				return
			}

			// TODO: end
			u.curveEnd()
			// u.frameIndex = 0
			// u.end = true
			// u.playing = false
			// if u.done != nil {
			// 	u.done <- true
			// }
			return
		}
	}

	var cnt = u.FrameCount()
	if u.frameIndex > cnt-1 {
		if u.swing {
			u.frameIndex = cnt - 2
			return
		}

		if u.loop {
			u.frameIndex = 0
			return
		}

		u.curveEnd()
		// u.frameIndex = cnt - 1
		// u.end = true
		// u.playing = false
		// if u.done != nil {
		// 	u.done <- true
		// }
		return
	}
}

func (u *Curve) curveStart() {
	u.playing = true
}

func (u *Curve) curveEnd() {
	if u.reversed {
		u.playing = false
		u.frameIndex = 0
		u.end = true
		if u.done != nil {
			u.done <- true
		}
		u.END.Emit(u)

		return
	}

	u.playing = false
	u.frameIndex = u.FrameCount() - 1
	u.end = true
	if u.done != nil {
		u.done <- true
	}
	u.END.Emit(u)
}

func (u *Curve) OnEnd(fn EventFunc) {
	u.END.On(fn)
}

func (u *Curve) OnDelete(fn EventFunc) {
	u.DEL.On(fn)
}

func (u *Curve) Q() float64 {
	var (
		x = float64(u.frameIndex) / float64(u.FrameCount())
		y = u.easingFunc(x)
	)

	return y
}

func (u *Curve) FrameCount() int {
	if u.duration == 0 {
		return 1
	}

	return int(math.Ceil(u.duration.Seconds() * u.tps))
}

func (u *Curve) IsEnd() bool {
	return u.end
}

func (u *Curve) Update() error {
	u.curveUpdate()

	return nil
}

func (u *Curve) Start() {
	u.curveStart()
}

func (u *Curve) Wait() {
	u.done = make(chan bool)

	for d := range u.done {
		if d {
			close(u.done)
			break
		}
	}
}
