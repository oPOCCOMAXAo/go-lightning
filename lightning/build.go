package lightning

func Build(start, end Vector, stopDistance float64, generator DeltaGenerator) []Vector {
	res := NewQueue()
	toProcess := NewStack()
	current := start
	next := end
	hasNext := true
	res.Push(start)
	for hasNext {
		if current.Sub(next).LengthManhattan() <= stopDistance {
			res.Push(next)
			current = next
			next, hasNext = toProcess.Pop()
		} else {
			m1, m2 := MakeIntermediatePoints(current, next, generator.Next())
			toProcess.Push(next)
			toProcess.Push(m2)
			next = m1
		}
	}
	return res.Array()
}
