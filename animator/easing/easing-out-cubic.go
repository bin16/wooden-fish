package easing

import "math"

// https://easings.net/#easeOutCubic
func EaseOutCubic(x float64) float64 {
	return 1 - math.Pow(1-x, 3)
}

// https://easings.net/#easeInCubic
func EaseInCubic(x float64) float64 {
	return x * x * x
}

// https://easings.net/#easeInOutCubic
func EaseInOutCubic(x float64) float64 {
	if x < .5 {
		return 4 * x * x * x
	}

	return 1 - math.Pow(-2*x+2, 3)/2
}

// https://easings.net/#easeInCirc
func EaseInCirc(x float64) float64 {
	return 1 - math.Sqrt(1-math.Pow(x, 2))
}

// https://easings.net/#easeOutCirc
func EaseOutCirc(x float64) float64 {
	return math.Sqrt(1 - math.Pow(x-1, 2))
}
