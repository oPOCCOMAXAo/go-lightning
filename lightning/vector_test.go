package lightning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	testCases := []struct {
		desc     string
		initial  Vector
		expected []Vector
	}{
		{
			desc:    "default",
			initial: Vector{X: 2, Y: 1},
			expected: []Vector{
				{X: 1, Y: -2},
				{X: -2, Y: -1},
				{X: -1, Y: 2},
				{X: 2, Y: 1},
			},
		},
		{
			desc:    "vertical",
			initial: Vector{X: 0, Y: 1},
			expected: []Vector{
				{X: 1, Y: 0},
				{X: 0, Y: -1},
				{X: -1, Y: 0},
				{X: 0, Y: 1},
			},
		},
		{
			desc:    "diagonal",
			initial: Vector{X: 1, Y: 1},
			expected: []Vector{
				{X: 1, Y: -1},
				{X: -1, Y: -1},
				{X: -1, Y: 1},
				{X: 1, Y: 1},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := []Vector{}
			iter := tC.initial
			for i := 0; i < 4; i++ {
				iter = iter.RotateRight()
				actual = append(actual, iter)
			}
			assert.Equal(t, tC.expected, actual)
		})
	}
}
