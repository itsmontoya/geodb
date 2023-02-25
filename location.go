package geodb

import "math"

// MakeLocation will return a new location
func MakeLocation(lat, lon Degree) Location {
	return Location{
		lat: lat.toRadians(),
		lon: lon.toRadians(),
	}
}

// Location represents a location
type Location struct {
	// Latitude in Radians
	lat Radian
	// Longitude in Radians
	lon Radian
}

func (l *Location) Distance(inbound Location) (m Meter) {
	lat := inbound.lat
	lon := inbound.lon
	// latH is the haversine value for the latitude
	latH := getHaversine(lat - l.lat)
	// longH is the haversine value for the longitude
	lonH := getHaversine(lon-l.lon) * Radian(math.Cos(float64(l.lat))) * Radian(math.Cos(float64(lat)))

	// Square of half the chord length between the points.
	a := float64(latH + lonH)
	// Angular distance of Radians
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return Meter(earthRadius * c)
}
