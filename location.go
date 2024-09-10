package geodb

import (
	"encoding/json"
	"fmt"
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

// Distance calculates the great-circle distance between two locations.
func (l *Location) Distance(inbound Location) (m Meter) {
	// Calculate differences in latitude and longitude
	deltaLat := inbound.lat - l.lat
	deltaLon := inbound.lon - l.lon

	// Use the getHaversine function to calculate the angular distance
	angularDistance := getHaversine(l.lat, inbound.lat, deltaLat, deltaLon)

	// Calculate the distance in meters
	return Meter(earthRadius * angularDistance)
}

func (l *Location) String() string {
	return fmt.Sprintf("Lat: %f / Lon: %f", l.lat.toDegrees(), l.lon.toDegrees())
}
