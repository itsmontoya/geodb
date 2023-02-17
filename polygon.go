package geodb

import "github.com/itsmontoya/geodb/math"

var _ Entry = &Polygon{}

// NewPolygon will return a new poly target
func NewPolygon(key string, points []Location, center Location) *Polygon {
	var p Polygon
	p.key = key
	p.segs = make([]*math.Segment, len(points))
	j := len(points) - 1

	for i, a := range points {
		b := points[j]
		aX := float64(a.lon.toDegrees())
		aY := float64(a.lat.toDegrees())
		bX := float64(b.lon.toDegrees())
		bY := float64(b.lat.toDegrees())
		p.segs[i] = math.NewSegment(aX, aY, bX, bY)
		j = i
	}

	return &p
}

// Polygon is a polygon target
type Polygon struct {
	key    string
	segs   []*math.Segment
	center Location
}

func (p *Polygon) Key() string {
	return p.key
}

func (p *Polygon) Locations() (locs []Location) {
	locs = make([]Location, 0, len(p.segs))
	for _, seg := range p.segs {
		point := seg.GetStart()
		loc := MakeLocation(Degree(point.X), Degree(point.Y))
		locs = append(locs, loc)
	}

	return
}

// IsWithin will return whether ot not a point at a given lat and lon are within a poly target
func (p *Polygon) IsWithin(l Location) (within bool) {
	var n int
	x := float64(l.lat.toDegrees())
	y := float64(l.lon.toDegrees())
	tgt := math.MakePoint(x, y)
	for _, seg := range p.segs {
		if seg.HasPoint(&tgt) {
			// Our target point matches one of the polygon points, return true
			return true
		}

		if seg.Intersects(tgt.X, tgt.Y) {
			// Point intersects, swap within value
			within = !within
			n++
		}
	}

	return
}

func (p *Polygon) Center() Location {
	return p.center
}

func (p *Polygon) pointIntersects(x, y float64, a, b *math.Point) (ok bool) {
	seg := math.NewSegment(a.X, a.Y, b.X, b.Y)
	return seg.Intersects(x, y)
}
