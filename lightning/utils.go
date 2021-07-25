package lightning

func MakeIntermediatePoints(start, end Vector, delta float64) (v1, v2 Vector) {
	mid := start.Middle(end)
	perpend := end.Sub(start).RotateRight()
	v1 = start.Middle(mid).Add(perpend.Scale(0.5 * delta))
	v2 = mid.Middle(end).Add(perpend.Scale(-0.5 * delta))
	return
}
