package geodb

import "github.com/itsmontoya/pip"

var _ Shape = &Polygon{}

// NewPolygon will return a new poly target
func NewPolygon(locs []Location, center Location) *Polygon {
	var p Polygon
	p.polygon = pip.New(toPoints(locs))
	p.center = center
	return &p
}

// Polygon is a polygon target
type Polygon struct {
	polygon *pip.Polygon
	center  Location
}

// IsWithin will return whether ot not a point at a given lat and lon are within a poly target
func (p *Polygon) IsWithin(l Location) (within bool) {
	x := float64(l.lon.toDegrees())
	y := float64(l.lat.toDegrees())
	point := pip.MakePoint(x, y)
	return p.polygon.IsWithin(point)
}

func (p *Polygon) Center() Location {
	return p.center
}
