package geodb

type Radian float64

// toDegrees will convert radians to degrees
func (r Radian) toDegrees() Degree {
	return Degree(float64(r) / radianMultiplier)
}
