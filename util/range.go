package util

import (
	"cmp"
)

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
	if len(b) > 0 {
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

func Map[A, B any](a []A, fn func(A) B) (r []B) {
	for _, n := range a {
		r = append(r, fn(n))
	}

	return
}

func OR[T any](flag bool, a, b T) T {
	if flag {
		return a
	}

	return b
}

func NotZero[T ~int | ~float64](items ...T) T {
	for _, n := range items {
		if n != T(0) {
			return n
		}
	}

	return T(0)
}
