package curve

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func New(opts ...CurveOpt) *Curve {
	return NewCurve(opts...)
}

func AutoPlay() CurveOpt {
	return Options.AutoPlay()
}

func Duration(d time.Duration) CurveOpt {
	return Options.Duration(d)
}

func OnEnd(fn EventFunc) CurveOpt {
	return Options.OnEnd(fn)
}

func Easing(fn EasingFunc) CurveOpt {
	return Options.EasingFunc(fn)
}

func TPS(d int) CurveOpt {
	return Options.TPS(d)
}

func Loop() CurveOpt {
	return Options.Loop()
}

func Swing() CurveOpt {
	return Options.Swing()
}

func NewCurve(opts ...CurveOpt) *Curve {
	var id = newID()
	var curve = &Curve{
		easingFunc: easeOutCirc,
		tps:        float64(ebiten.TPS()),
		duration:   time.Second,
		frameIndex: 0,
		end:        false,
		id:         id,
	}
	for _, fn := range opts {
		fn(curve)
	}
	curve.curveInit()

	curveDB.Store(id, curve)

	// fmt.Println("new curve", id)

	return curve
}

func easeOutCirc(x float64) float64 {
	return math.Sqrt(1 - math.Pow(x-1, 2))
}
