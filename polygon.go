package geodb

import (
	"fmt"

	"github.com/itsmontoya/pip"
)

var _ Shape = &Polygon{}

// NewPolygon will return a new poly target
func NewPolygon(locs []Location) (pp *Polygon, err error) {
	if len(locs) < 3 {
		err = fmt.Errorf("invalid number of locations, have <%d> and need a minimum of <3>", len(locs))
		return
	}

	var p Polygon
	p.polygon = pip.New(toPoints(locs))
	p.setReferencePoints(locs)
	return &p, nil
}

// Polygon is a polygon target
type Polygon struct {
	polygon *pip.Polygon
	center  Location
	radius  Meter
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

func (p *Polygon) Radius() Meter {
	return p.radius
}

func (p *Polygon) setReferencePoints(locs []Location) {
	lowLat, highLat, lowLon, highLon := getCorners(locs)
	centerLat := (highLat + lowLat) / 2
	centerLon := (highLon + lowLon) / 2
	p.center = MakeLocation(centerLat, centerLon)
	furthestPoint := MakeLocation(lowLat, lowLon)
	p.radius = p.center.Distance(furthestPoint)
}
