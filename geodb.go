package geodb

import (
	"sync"
)

// New returns a new GeoDB
func New(regionSize Meter) *GeoDB {
	var db GeoDB
	db.regionSize = regionSize
	return &db
}

// GeoDB maanges the GeoDB service
type GeoDB struct {
	mux sync.RWMutex

	// region size
	regionSize Meter
	// Internal region store
	rs []*region
}

// Insert will insert a location to the tree
func (g *GeoDB) Insert(key string, shape Shape) (err error) {
	g.mux.Lock()
	defer g.mux.Unlock()
	e := makeEntry(key, shape)
	g.insert(e)
	return
}

// GetMatches will return the matching location keys for the provided latitude and longitude
func (g *GeoDB) GetMatches(l Location) (matches []string, err error) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return g.getMatches(l), nil
}

// RegionsLen will return the number of regions
func (g *GeoDB) RegionsLen() (n int, err error) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return len(g.rs), nil
}

// EntriesLen will return the number of targets
func (g *GeoDB) EntriesLen() (n int, err error) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	for _, r := range g.rs {
		n += len(r.ts)
	}

	return
}

// insert will insert a location to the tree
func (g *GeoDB) insert(e entry) {
	if g.tryInsert(e) {
		return
	}

	g.createRegion(e)
}

// insert will insert a location to the tree
func (g *GeoDB) tryInsert(e entry) (inserted bool) {
	for _, r := range g.rs {
		if r.insert(e) {
			inserted = true
		}
	}

	return
}

// Insert will insert a location to the tree
func (g *GeoDB) createRegion(e entry) {
	// No matches were found, so we must create a new region
	r := newRegionFromEntry(e, g.regionSize)
	g.rs = append(g.rs, r)
}

func (g *GeoDB) getMatches(l Location) (matches []string) {
	for _, r := range g.rs {
		matches = r.appendMatches(matches, l)
	}

	return
}
