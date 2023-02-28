package geodb

var _ Shape = &Radius{}

func NewRadius(radius Meter, loc Location) *Radius {
	return &Radius{
		radius:   radius,
		Location: loc,
	}
}

type Radius struct {
	// Radius in meters
	radius Meter
	// Location of Radius
	Location
}

// IsWithin returns whether or not a latitude and longitude are within range
func (r *Radius) IsWithin(l Location) bool {
	return r.Distance(l) <= r.radius
}

func (r *Radius) Center() Location {
	return r.Location
}

func (r *Radius) Radius() Meter {
	return r.radius
}
