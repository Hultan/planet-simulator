package data

import (
	"math"
)

type Vector2 struct {
	X, Y float64
}

func (v Vector2) Add(o Vector2) Vector2 {
	return Vector2{v.X + o.X, v.Y + o.Y}
}

func (v Vector2) Sub(o Vector2) Vector2 {
	return Vector2{v.X - o.X, v.Y - o.Y}
}

func (v Vector2) Mul(m float64) Vector2 {
	return Vector2{v.X * m, v.Y * m}
}

func (v Vector2) Div(d float64) Vector2 {
	return Vector2{v.X / d, v.Y / d}
}

func (v Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
