package tools

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (v1 Vector) To(v2 Vector) Vector {
	return Vector{v2.X - v1.X, v2.Y - v1.Y, v2.Z - v1.Z}
}

func (v1 Vector) Dist(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v2.X-v1.X, 2) + math.Pow(v2.Y-v1.Y, 2))
}

func (v1 *Vector) Add(v2 Vector) {
	v1.X += v2.X
	v1.Y += v2.Y
}

func (v1 *Vector) Mul(m float64) {
	v1.X *= m
	v1.Y *= m
}

func Add(v1 Vector, v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func Mul(v1 Vector, m float64) Vector {
	return Vector{v1.X * m, v1.Y * m, v1.Z * m}
}
