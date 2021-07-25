package lightning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicGenerator(t *testing.T) {
	testCases := []struct {
		desc string
		gen  DeltaGenerator
		res  []float64
	}{
		{
			desc: "default",
			gen:  NewCyclicGenerator([]float64{0, 1, 2, 3}),
			res:  []float64{3, 2, 1, 0, 3},
		},
		{
			desc: "empty",
			gen:  NewCyclicGenerator(nil),
			res:  []float64{0, 0, 0},
		},
		{
			desc: "single",
			gen:  NewCyclicGenerator([]float64{1}),
			res:  []float64{1, 1, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			total := len(tC.res)
			res := make([]float64, total)
			for i := 0; i < total; i++ {
				res[i] = tC.gen.Next()
			}
			assert.Equal(t, tC.res, res)
		})
	}
}
