package util

import "cmp"

// a > 0
// a   => [0, a) 0, 1, 2, ..., a - 1
// a,b => [a, b) a, a+1, ...,  b - 1
func Range[T cmp.Ordered](a T, b ...T) (r []T) {
	var (
		zero = T(0)
		low  = zero
		high = T(a)
		step = T(1)
	)
	if len(b) == 1 {
		low = a
		high = b[0]
	}
	if len(b) > 1 {
		step = b[1]
	}

	if high > low && step > zero {
		for d := low; d < high; d += step {
			r = append(r, d)
		}

		return
	}

	if high < low && step < zero {
		for d := low; d < high; d += step {
			r = append(r, d)
		}

		return
	}

	return
}
