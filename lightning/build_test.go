package lightning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		desc  string
		start Vector
		end   Vector
		dist  float64
		gen   DeltaGenerator
		res   []Vector
	}{
		{
			desc:  "straight",
			start: Vector{X: 0, Y: 0},
			end:   Vector{X: 4, Y: 0},
			dist:  1,
			gen:   NewCyclicGenerator([]float64{0}),
			res: []Vector{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1.5, Y: 0},
				{X: 2.5, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
			},
		},
		{
			desc:  "curve",
			start: Vector{X: 0, Y: 0},
			end:   Vector{X: 3, Y: 0},
			dist:  1,
			gen:   NewCyclicGenerator([]float64{0.5}),
			res: []Vector{
				{X: 0, Y: 0},
				{X: 0, Y: -0.375},
				{X: 0.75, Y: -0.375},
				{X: 0.75, Y: -0.75},
				{X: 1.5, Y: -0.75},
				{X: 1.875, Y: -0.375},
				{X: 1.875, Y: 0},
				{X: 1.125, Y: 0},
				{X: 1.125, Y: 0.375},
				{X: 1.5, Y: 0.75},
				{X: 2.25, Y: 0.75},
				{X: 2.25, Y: 0.375},
				{X: 3, Y: 0.375},
				{X: 3, Y: 0},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := Build(tC.start, tC.end, tC.dist, tC.gen)
			assert.Equal(t, tC.res, res)
		})
	}
}
