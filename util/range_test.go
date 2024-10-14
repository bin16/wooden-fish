package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	var data = [][2][]int{
		{
			{5},
			{0, 1, 2, 3, 4},
		},
		{
			{7, 9},
			{7, 8},
		},
		{
			{3, 2},
			{3},
		},
	}

	for _, d := range data {
		r := Range[int](d[0][0], d[0][1:]...)
		assert.Equal(t, d[1], r)
	}
}
