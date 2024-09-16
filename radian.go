package geodb

import "math"

type Radian float64

// toDegrees will convert radians to degrees
func (r Radian) toDegrees() Degree {
	return Degree(float64(r) / radianMultiplier)
}

func (r Radian) cosine() float64 {
	return math.Cos(float64(r))
}
