package geodb

func newRegion(e Entry, radius Meter) *Region {
	var r Region
	r.center = e.Center()
	r.radius = radius
	return &r
}

type Region struct {
	// Location of region
	center Location
	// Radius of target coverage in meters
	radius Meter
	// Targets within region
	ts []Entry
}

func (r *Region) insert(entry Entry) (ok bool) {
	center := entry.Center()
	if ok = r.isContainedBy(center); !ok {
		return false
	}

	r.ts = append(r.ts, entry)
	return true
}

// isWithinRadius returns whether or not a latitude and longitude are within range
func (r *Region) isWithinRadius(l Location) bool {
	if r.center.Distance(l) > r.radius {
		return false
	}

	return true
}

// appendMatches will append Matches
func (r *Region) appendMatches(s []string, l Location) []string {
	if !r.isWithinRadius(l) {
		// Return without modifying
		return s
	}

	// Iterate through all targets
	for _, entry := range r.ts {
		if !entry.IsWithin(l) {
			// This target is not within radius, continue
			continue
		}

		// Append matching key
		s = append(s, entry.Key())
	}

	return s
}

// isContainedBy returns whether or not a target is completely contained by a region
func (r *Region) isContainedBy(center Location) bool {
	if r.center.Distance(center)+r.radius > r.radius {
		return false
	}

	return true
}
