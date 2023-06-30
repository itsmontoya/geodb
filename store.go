package geodb

// store maanges the store service
type store struct {
	// region size
	regionSize Meter
	// Internal region store
	rs []*region
}

// Insert will insert a location to the tree
func (s *store) Insert(key string, shape Shape) (err error) {
	e := makeEntry(key, shape)
	s.insert(e)
	return
}

// GetMatches will return the matching location keys for the provided latitude and longitude
func (s *store) GetMatches(l Location) (matches []string, err error) {
	for _, r := range s.rs {
		matches = r.appendMatches(matches, l)
	}

	return
}

// RegionsLen will return the number of regions
func (s *store) RegionsLen() (n int, err error) {
	return len(s.rs), nil
}

// EntriesLen will return the number of targets
func (s *store) EntriesLen() (n int, err error) {
	for _, r := range s.rs {
		n += len(r.ts)
	}

	return
}

// insert will insert a location to the tree
func (s *store) insert(e entry) {
	if s.tryInsert(e) {
		return
	}

	s.createRegion(e)
}

// insert will insert a location to the tree
func (s *store) tryInsert(e entry) (inserted bool) {
	for _, r := range s.rs {
		if r.insert(e) {
			inserted = true
		}
	}

	return
}

// Insert will insert a location to the tree
func (s *store) createRegion(e entry) {
	// No matches were found, so we must create a new region
	r := newRegionFromEntry(e, s.regionSize)
	s.rs = append(s.rs, r)
}
