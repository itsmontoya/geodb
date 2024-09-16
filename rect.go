package geodb

type Rect struct {
	Min Coordinates
	Max Coordinates
}

func (r *Rect) Center() (l Location) {
	l.lat = r.centerLat()
	l.lon = r.centerLon()
	return
}

func (r *Rect) IsZero() bool {
	switch {
	case !r.Min.IsZero():
		return false
	case !r.Max.IsZero():
		return false
	default:
		return true
	}
}

func (r *Rect) SetRect(in Rect) {
	if r.IsZero() {
		*r = in
		return
	}

	r.SetLat(in.Max.Latitude)
	r.SetLat(in.Min.Latitude)
	r.SetLon(in.Max.Longitude)
	r.SetLon(in.Min.Longitude)
}

func (r *Rect) SetCoords(in Coordinates) {
	if r.IsZero() {
		r.Max = in
		r.Min = in
		return
	}

	r.SetLat(in.Latitude)
	r.SetLon(in.Longitude)
}

func (r *Rect) SetLat(lat Degree) {
	switch {
	case lat > r.Max.Latitude:
		r.Max.Latitude = lat
	case lat < r.Min.Latitude:
		r.Min.Latitude = lat
	}
}

func (r *Rect) SetLon(lon Degree) {
	switch {
	case lon > r.Max.Longitude:
		r.Max.Longitude = lon
	case lon < r.Min.Longitude:
		r.Min.Longitude = lon
	}
}

func (r *Rect) EnsureValidRanges() {
	// Ensure minLat and maxLat are within -90° and +90°
	if r.Min.Latitude < -90 {
		r.Min.Latitude = -90
	}

	if r.Max.Latitude > 90 {
		r.Max.Latitude = 90
	}

	// Normalize longitudes to be within -180° and +180°
	if r.Min.Longitude < -180 {
		r.Min.Longitude += 360
	}

	if r.Max.Longitude > 180 {
		r.Max.Longitude -= 360
	}
}

func (r *Rect) HasGreaterLatitudeDelta() bool {
	latDelta := r.Max.Latitude - r.Min.Latitude
	if latDelta < 0 {
		latDelta *= -1
	}

	lonDelta := r.Max.Longitude - r.Min.Longitude
	if lonDelta < 0 {
		lonDelta *= -1
	}

	return latDelta > lonDelta
}

func (r *Rect) Split() (r1, r2 Rect) {
	if r.HasGreaterLatitudeDelta() {
		return r.splitByLatitude()
	}

	return r.splitByLongitude()
}

func (r *Rect) GetMatchingNode(ns ...node) (match node, ok bool) {
	for _, n := range ns {
		if n.IsFullyContained(r) {
			return n, true
		}
	}

	for _, n := range ns {
		if !n.DoesNotOverlap(r) {
			return n, true
		}
	}

	return
}

func (r *Rect) ContainsCoordinates(c Coordinates) (contains bool) {
	switch {
	case c.Latitude > r.Max.Latitude:
		return false
	case c.Latitude < r.Min.Latitude:
		return false
	case c.Longitude > r.Max.Longitude:
		return false
	case c.Longitude < r.Min.Longitude:
		return false
	default:
		return true
	}
}

func (r *Rect) IsFullyContained(in *Rect) (contained bool) {
	switch {
	case in.Max.Latitude > r.Max.Latitude:
		return false
	case in.Min.Latitude < r.Min.Latitude:
		return false
	case in.Max.Longitude > r.Max.Longitude:
		return false
	case in.Min.Longitude < r.Min.Longitude:
		return false
	default:
		return true
	}
}

func (r *Rect) DoesNotOverlap(in *Rect) (notContained bool) {
	switch {
	case in.Min.Latitude > r.Max.Latitude:
		return true
	case in.Max.Latitude < r.Min.Latitude:
		return true
	case in.Min.Longitude > r.Max.Longitude:
		return true
	case in.Max.Longitude < r.Min.Longitude:
		return true
	default:
		return false
	}
}

func (r *Rect) centerLat() (lat Radian) {
	center := ((r.Min.Latitude + r.Max.Latitude) / 2.0)
	return center.toRadians()
}

func (r *Rect) centerLon() (lon Radian) {
	center := r.centerLonDegree()
	return center.toRadians()
}

func (r *Rect) centerLonDegree() (lon Degree) {
	if r.Max.Longitude >= r.Min.Longitude {
		// Standard case
		return (r.Min.Longitude + r.Max.Longitude) / 2.0
	}

	// Rectangle crosses the International Date Line
	lon = (r.Min.Longitude + r.Max.Longitude + 360.0) / 2.0
	if lon > 180.0 {
		lon -= 360.0
	}

	return
}

func (r *Rect) splitByLatitude() (r1, r2 Rect) {
	middle := (r.Max.Latitude + r.Min.Latitude) / 2
	r1.Max.Latitude = middle
	r1.Min.Latitude = r.Min.Latitude
	r2.Max.Latitude = r.Max.Latitude
	r2.Min.Latitude = middle

	r1.Max.Longitude = r.Max.Longitude
	r1.Min.Longitude = r.Min.Longitude
	r2.Max.Longitude = r.Max.Longitude
	r2.Min.Longitude = r.Min.Longitude
	return
}

func (r *Rect) splitByLongitude() (r1, r2 Rect) {
	middle := (r.Max.Longitude + r.Min.Longitude) / 2
	r1.Max.Longitude = middle
	r1.Min.Longitude = r.Min.Longitude
	r2.Max.Longitude = r.Max.Longitude
	r2.Min.Longitude = middle

	r1.Max.Latitude = r.Max.Latitude
	r1.Min.Latitude = r.Min.Latitude
	r2.Max.Latitude = r.Max.Latitude
	r2.Min.Latitude = r.Min.Latitude
	return
}
