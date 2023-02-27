package geodb

func newRegion(s Shape, radius Meter) *region {
	var r region
	r.center = s.Center()
	r.radius = radius
	return &r
}

func newRegionFromEntry(e entry, radius Meter) *region {
	r := newRegion(e.shape, radius)
	r.ts = append(r.ts, e)
	return r
}

type region struct {
	// Location of region
	center Location
	// Radius of target coverage in meters
	radius Meter
	// Targets within region
	ts []entry
}

func (r *region) insert(e entry) (ok bool) {
	center := e.shape.Center()
	if ok = r.isContainedBy(center); !ok {
		return false
	}

	r.ts = append(r.ts, e)
	return true
}

// isContainedBy returns whether or not a target is completely contained by a region
func (r *region) isContainedBy(center Location) bool {
	return r.center.Distance(center)+r.radius <= r.radius
}

// isWithinRadius returns whether or not a latitude and longitude are within range
func (r *region) isWithinRadius(l Location) bool {
	return r.center.Distance(l) <= r.radius
}

// appendMatches will append Matches
func (r *region) appendMatches(s []string, l Location) []string {
	if !r.isWithinRadius(l) {
		// Return without modifying
		return s
	}

	// Iterate through all targets
	for _, entry := range r.ts {
		if !entry.shape.IsWithin(l) {
			// This target is not within radius, continue
			continue
		}

		// Append matching key
		s = append(s, entry.key)
	}

	return s
}
