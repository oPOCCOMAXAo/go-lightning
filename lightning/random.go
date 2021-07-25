package lightning

import "math/rand"

type DeltaGenerator interface {
	Next() float64
}

type randGenerator struct {
	rand  *rand.Rand
	delta float64
}

func (s *randGenerator) Next() float64 {
	return (s.rand.Float64()*2 - 1) * s.delta
}

func NewRandGenerator(deltaModulus float64, seed int64) DeltaGenerator {
	return &randGenerator{
		rand:  rand.New(rand.NewSource(seed)),
		delta: deltaModulus,
	}
}

type cyclicGenerator struct {
	cycle []float64
	ptr   int
}

func (c *cyclicGenerator) Next() float64 {
	c.ptr--
	if c.ptr < 0 {
		c.ptr = len(c.cycle) - 1
	}
	return c.cycle[c.ptr]
}

func NewCyclicGenerator(cycle []float64) DeltaGenerator {
	if len(cycle) == 0 {
		cycle = []float64{0}
	}
	return &cyclicGenerator{
		cycle: cycle,
	}
}
