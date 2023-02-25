package geodb

import (
	"math"

	"github.com/itsmontoya/pip"
)

const (
	// Radius of the earth in meters
	earthRadius = 6371000
	// Degrees to radians multiplier
	radianMultiplier = (math.Pi / 180.0)
	// Meters to foot conversion value
	metersInFoot = 0.3048
	// Feet to miles conversion value
	feetInMile = 5280
)

// ToMeters will convert a foot value to meters
func ToMeters(feet float64) (meters float64) {
	return feet * metersInFoot
}

// MilesToMeters will return the number of meters in a provided number of miles
func MilesToMeters(miles float64) (meters float64) {
	return ToMeters(miles * feetInMile)
}

// GetDistance will return distance
func GetDistance(a, b Location) Unit {
	return a.Distance(b)
}

// getHaversine will return a haversine value from a provided delta
func getHaversine(delta Radian) Radian {
	return Radian(math.Pow(math.Sin((float64(delta))/2), 2))
}

func toPoints(locs []Location) (points []pip.Point) {
	points = make([]pip.Point, 0, len(locs))
	for _, a := range locs {
		point := pip.MakePoint(float64(a.lon.toDegrees()), float64(a.lat.toDegrees()))
		points = append(points, point)
	}

	return
}
