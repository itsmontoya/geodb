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

// getHaversine calculates the angular distance between two points on a sphere
// given their latitudes and the differences in their latitudes and longitudes.
// Parameters:
//
//	latA, latB: Latitudes of the two points in radians.
//	deltaLat: Difference in latitude between the two points in radians.
//	deltaLon: Difference in longitude between the two points in radians.
//
// Returns:
//
//	The angular distance in radians.
func getHaversine(latA, latB, deltaLat, deltaLon Radian) Radian {
	// Central angle component used to compute the angular distance
	centralAngleComponent := math.Pow(math.Sin(float64(deltaLat)/2), 2) +
		math.Cos(float64(latA))*math.Cos(float64(latB))*
			math.Pow(math.Sin(float64(deltaLon)/2), 2)

	// Calculate the angular distance in radians
	angularDistance := 2 * math.Atan2(math.Sqrt(centralAngleComponent), math.Sqrt(1-centralAngleComponent))
	return Radian(angularDistance)
}

func toPoints(locs []Location) (points []pip.Point) {
	points = make([]pip.Point, 0, len(locs))
	for _, a := range locs {
		point := pip.MakePoint(float64(a.lon.toDegrees()), float64(a.lat.toDegrees()))
		points = append(points, point)
	}

	return
}

func getCorners(locs []Location) (lowLat, highLat, lowLon, highLon Degree) {
	first := locs[0]
	lowLat = first.lat.toDegrees()
	highLat = lowLat
	lowLon = first.lon.toDegrees()
	highLon = lowLon

	for _, loc := range locs {
		lat := loc.lat.toDegrees()
		switch {
		case lat < lowLat:
			lowLat = lat
		case lat > highLat:
			highLat = lat
		}

		lon := loc.lon.toDegrees()
		switch {
		case lon < lowLon:
			lowLon = lon
		case lon > highLon:
			highLon = lon
		}
	}

	return
}
