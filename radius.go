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
func (r *Radius) IsWithin(c Coordinates) bool {
	return r.Distance(c.Location()) <= r.radius
}

func (r *Radius) Center() Location {
	return r.Location
}

func (r *Radius) Radius() Meter {
	return r.radius
}

func (r *Radius) Rect() (rect Rect) {
	// Delta latitude in Radians
	deltaLat := Radian(r.radius / earthRadius)
	// Delta longitude in Radians

	deltaLon := Radian(float64(r.radius) / (earthRadius * r.lat.cosine()))

	// Compute min and max latitudes
	rect.Min.Latitude = (r.lat - deltaLat).toDegrees()
	rect.Max.Latitude = (r.lat + deltaLat).toDegrees()

	// Compute min and max longitudes
	rect.Min.Longitude = (r.lon - deltaLon).toDegrees()
	rect.Max.Longitude = (r.lon + deltaLon).toDegrees()
	rect.EnsureValidRanges()
	return
}
