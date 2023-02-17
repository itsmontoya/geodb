package math

import "math"

// NewSegment will return a new Segment from the provided coordinate values
func NewSegment(x1, y1, x2, y2 float64) *Segment {
	var s Segment
	s.p1 = Point{x1, y1}
	s.p2 = Point{x2, y2}
	s.m, s.b, s.st = getSlopeOffset(s.p1, s.p2)
	return &s
}

// Segment is a line segment
type Segment struct {
	p1 Point
	p2 Point

	m  float64
	b  float64
	st slope
}

func (s *Segment) GetStart() Point {
	return s.p1
}

func (s *Segment) GetEnd() Point {
	return s.p2
}

// GetX will get the x value at the given y coordinate
func (s *Segment) GetX(y float64) (x float64) {
	switch s.st {
	case slopeDefault:
		return solveForX(y, s.m, s.b)

	case slopeHorizontal:
		return math.NaN()

	case slopeVertical:
		return s.p1.X
	}

	return
}

// GetY will get the y value at the given x coordinate
func (s *Segment) GetY(x float64) (y float64) {
	switch s.st {
	case slopeDefault:
		return solveForY(x, s.m, s.b)

	case slopeHorizontal:
		return s.p1.Y

	case slopeVertical:
		return math.NaN()
	}

	return
}

// Intersects tests if a line extending along the x axis (y = b)
func (s *Segment) Intersects(x, y float64) (ok bool) {
	switch s.st {
	case slopeDefault:
		var nx float64
		if nx = s.GetX(y); math.IsNaN(nx) || nx < x {
			return
		}

		return s.isWithinSegment(nx, y)
	case slopeHorizontal:
		if y != s.p1.Y {
			return
		}

		_, max := s.getXVals()
		return x <= max
	case slopeVertical:
		return x <= s.p1.X && s.isWithinY(y)
	}

	return
}

// HasPoint will return whether or not a matching point exists within the segment
func (s *Segment) HasPoint(pt *Point) bool {
	return s.p1 == *pt || s.p2 == *pt
}

func (s *Segment) isWithinX(x float64) bool {
	min, max := s.getXVals()
	return x >= min && x <= max
}

func (s *Segment) isWithinY(y float64) bool {
	min, max := s.getYVals()
	return y >= min && y <= max
}

// isWithinSegment will return whether a provided point (x,y) is potentially within the range of the segment
// Note: This does NOT test whether or not the provided coordinate exists within the segment
func (s *Segment) isWithinSegment(x, y float64) bool {
	return s.isWithinX(x) && s.isWithinY(y)
}

func (s *Segment) getXVals() (min, max float64) {
	if s.p1.X <= s.p2.X {
		return s.p1.X, s.p2.X
	}

	return s.p2.X, s.p1.X
}

func (s *Segment) getYVals() (min, max float64) {
	if s.p1.Y <= s.p2.Y {
		return s.p1.Y, s.p2.Y
	}

	return s.p2.Y, s.p1.Y
}
