package geodb

import "github.com/itsmontoya/pip"

func NewCoordinates(lat, lon Degree) *Coordinates {
	c := MakeCoordinates(lat, lon)
	return &c
}

func MakeCoordinates(lat, lon Degree) (c Coordinates) {
	c.Latitude = lat
	c.Longitude = lon
	return
}

type Coordinates struct {
	Latitude  Degree `json:"lat"`
	Longitude Degree `json:"lon"`
}

func (c *Coordinates) IsZero() bool {
	switch {
	case c.Latitude != 0:
		return false
	case c.Longitude != 0:
		return false
	default:
		return true
	}
}

func (c *Coordinates) Location() Location {
	return MakeLocation(c.Latitude, c.Longitude)
}

func (c *Coordinates) toPoint() pip.Point {
	return pip.MakePoint(c.Longitude.Float64(), c.Latitude.Float64())
}
