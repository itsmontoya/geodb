package geodb

import "strconv"

type Degree float64

// toRadians will convert degrees to radians
func (d Degree) toRadians() Radian {
	return Radian(float64(d) * radianMultiplier)
}

func (d Degree) Float64() float64 {
	return float64(d)
}

func (d Degree) String() string {
	return strconv.FormatFloat(d.Float64(), 'f', -1, 64)
}
