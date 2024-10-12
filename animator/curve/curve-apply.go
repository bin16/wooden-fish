package curve

import "image/color"

func Apply[T int | float32 | float64](q float64, a, b T) T {
	if q == 0 {
		return a
	}

	if q == 1 {
		return b
	}

	return T(float64(a) + float64(b-a)*q)
}

func ApplyColor(q float64, m, n color.Color) color.Color {
	if q == 0 {
		return m
	}

	if q == 1 {
		return n
	}

	var (
		r0, g0, b0, a0 = m.RGBA()
		r1, g1, b1, a1 = n.RGBA()
		r              = uint8((float64(r0) + float64(r1-r0)*q) / 255)
		g              = uint8((float64(g0) + float64(g1-g0)*q) / 255)
		b              = uint8((float64(b0) + float64(b1-b0)*q) / 255)
		a              = uint8((float64(a0) + float64(a1-a0)*q) / 255)
	)

	return color.RGBA{r, g, b, a}
}
