package curve

import "time"

type CurveOpt func(u *Curve)
type CurveOptions struct{}

var Options CurveOptions
var CurveOpts CurveOptions

func (CurveOptions) Easing(fn EasingFunc) CurveOpt {
	return func(u *Curve) {
		u.easingFunc = fn
	}
}

func (CurveOptions) EasingFunc(fn EasingFunc) CurveOpt {
	return func(u *Curve) {
		u.easingFunc = fn
	}
}

func (CurveOptions) TPS(d int) CurveOpt {
	return func(u *Curve) {
		u.tps = float64(d)
	}
}

func (CurveOptions) Duration(d time.Duration) CurveOpt {
	return func(u *Curve) {
		u.duration = d
	}
}

func (CurveOptions) Loop() CurveOpt {
	return func(u *Curve) {
		u.loop = true
	}
}

func (CurveOptions) Swing() CurveOpt {
	return func(u *Curve) {
		u.swing = true
	}
}

func (CurveOptions) Reversed() CurveOpt {
	return func(u *Curve) {
		u.reversed = true
	}
}

func (CurveOptions) AutoPlay() CurveOpt {
	return func(u *Curve) {
		u.playing = true
	}
}

func (CurveOptions) OnEnd(fn EventFunc) CurveOpt {
	return func(u *Curve) {
		u.END.On(fn)
	}
}
