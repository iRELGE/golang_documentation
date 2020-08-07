package pointerreceiverRepo

import (
	"math"
)

// Vertex : struct has x and y type float
type Vertex struct {
	X, Y float64
}

// Abs : its a methode that defined on Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale : its a methode that define on vertex and can change value of X and Y that so it can midifie his recepter
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
