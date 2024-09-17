package geodb

import (
	"fmt"

	"github.com/itsmontoya/pip"
)

var _ Shape = &Polygon{}

// NewPolygon will return a new poly target
func NewPolygon(coords []Coordinates) (pp *Polygon, err error) {
	if len(coords) < 3 {
		err = fmt.Errorf("invalid number of coordinates, have <%d> and need a minimum of <3>", len(coords))
		return
	}

	var p Polygon
	p.border = coords
	p.polygon = pip.New(toPoints(coords))
	p.setRect()
	return &p, nil
}

// Polygon is a polygon target
type Polygon struct {
	border  []Coordinates
	polygon *pip.Polygon
	rect    Rect
}

// IsWithin will return whether ot not a point at a given lat and lon are within a poly target
func (p *Polygon) IsWithin(c Coordinates) (within bool) {
	x := float64(c.Longitude)
	y := float64(c.Latitude)
	point := pip.MakePoint(x, y)
	return p.polygon.IsWithin(point)
}

func (p *Polygon) Rect() Rect {
	return p.rect
}

func (p *Polygon) Radius() (radius Meter) {
	r := p.rect
	c := r.Center()
	for _, point := range p.border {
		loc := point.Location()
		distance := loc.Distance(c)
		if distance > radius {
			radius = distance
		}
	}

	return
}

func (p *Polygon) setRect() {
	for _, coords := range p.border {
		p.rect.SetCoords(coords)
	}
}
