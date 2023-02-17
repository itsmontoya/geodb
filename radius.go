package geodb

var _ Entry = &Radius{}

func NewRadius(key string, radius Meter, loc Location) *Radius {
	return &Radius{
		key:      key,
		radius:   radius,
		Location: loc,
	}
}

type Radius struct {
	// Key of Radius
	key string
	// Radius in meters
	radius Meter
	// Location of Radius
	Location
}

func (r *Radius) Key() string {
	return r.key
}

// IsWithin returns whether or not a latitude and longitude are within range
func (r *Radius) IsWithin(l Location) bool {
	if r.Distance(l) > r.radius {
		return false
	}

	return true
}

func (r *Radius) Center() Location {
	return r.Location
}
