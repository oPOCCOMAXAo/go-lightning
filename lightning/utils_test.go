package lightning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeIntermediatePoints(t *testing.T) {
	testCases := []struct {
		desc   string
		start  Vector
		end    Vector
		delta  float64
		result [2]Vector
	}{
		{
			desc: "zero",
		},
		{
			desc:  "vertical",
			start: Vector{X: 0, Y: 0},
			end:   Vector{X: 0, Y: 1},
			delta: 1,
			result: [2]Vector{
				{X: 0.5, Y: 0.25},
				{X: -0.5, Y: 0.75},
			},
		},
		{
			desc:  "horizontal",
			start: Vector{X: 0, Y: 0},
			end:   Vector{X: 1, Y: 0},
			delta: 1,
			result: [2]Vector{
				{X: 0.25, Y: -0.5},
				{X: 0.75, Y: 0.5},
			},
		},
		{
			desc:  "egyptian",
			start: Vector{X: 0, Y: 0},
			end:   Vector{X: 3, Y: 4},
			delta: 0.5,
			result: [2]Vector{
				{X: 1.75, Y: 0.25},
				{X: 1.25, Y: 3.75},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v1, v2 := MakeIntermediatePoints(tC.start, tC.end, tC.delta)
			assert.Equal(t, tC.result, [2]Vector{v1, v2})
		})
	}
}
