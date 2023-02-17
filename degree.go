package geodb

type Degree float64

// toRadians will convert degrees to radians
func (d Degree) toRadians() Radian {
	return Radian(float64(d) * radianMultiplier)
}
