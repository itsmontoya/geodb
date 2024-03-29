package geodb

import (
	"encoding/json"
	"fmt"
	"math"
)

func NewLocation(lat, lon Degree) *Location {
	l := MakeLocation(lat, lon)
	return &l
}

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

func (l *Location) MarshalJSON() (bs []byte, err error) {
	return json.Marshal(l.Coordinates())
}

func (l *Location) UnmarshalJSON(bs []byte) (err error) {
	var c Coordinates
	if err = json.Unmarshal(bs, &c); err != nil {
		return
	}

	*l = c.Location()
	return
}

func (l *Location) Coordinates() (c Coordinates) {
	c.Latitude = l.lat.toDegrees()
	c.Longitude = l.lon.toDegrees()
	return
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

func (l *Location) String() string {
	return fmt.Sprintf("Lat: %f / Lon: %f", l.lat.toDegrees(), l.lon.toDegrees())
}
