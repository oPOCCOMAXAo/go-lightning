package lightning

import "math"

var (
	vector0 = Vector{X: 0, Y: 0}
	vector1 = Vector{X: 1, Y: 1}
)

type Vector struct {
	X, Y float64
}

func (l Vector) Add(r Vector) Vector {
	return Vector{
		X: l.X + r.X,
		Y: l.Y + r.Y,
	}
}

func (l Vector) Sub(r Vector) Vector {
	return Vector{
		X: l.X - r.X,
		Y: l.Y - r.Y,
	}
}

func (l Vector) Scale(scale float64) Vector {
	return Vector{
		X: l.X * scale,
		Y: l.Y * scale,
	}
}

func (l Vector) LengthManhattan() float64 {
	return math.Abs(l.X) + math.Abs(l.Y)
}

func (l Vector) LengthSqr() float64 {
	return l.X*l.X + l.Y*l.Y
}

func (l Vector) RotateRight() Vector {
	return Vector{
		X: l.Y,
		Y: -l.X,
	}
}

func (l Vector) Middle(r Vector) Vector {
	return l.Add(r).Scale(0.5)
}

func (l Vector) Min(r Vector) Vector {
	return Vector{
		X: math.Min(l.X, r.X),
		Y: math.Min(l.Y, r.Y),
	}
}

func (l Vector) Max(r Vector) Vector {
	return Vector{
		X: math.Max(l.X, r.X),
		Y: math.Max(l.Y, r.Y),
	}
}
