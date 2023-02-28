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
)

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
