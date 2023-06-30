package geodb

import (
	"sync"
)

// New returns a new GeoDB
func New(regionSize Meter) *GeoDB {
	var db GeoDB
	db.s.regionSize = regionSize
	return &db
}

// GeoDB maanges the GeoDB service
type GeoDB struct {
	mux sync.RWMutex
	s   store
}

// Insert will insert a location to the tree
func (g *GeoDB) Insert(key string, shape Shape) (err error) {
	g.mux.Lock()
	defer g.mux.Unlock()
	return g.s.Insert(key, shape)
}

// GetMatches will return the matching location keys for the provided latitude and longitude
func (g *GeoDB) GetMatches(l Location) (matches []string, err error) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return g.s.GetMatches(l)
}

// RegionsLen will return the number of regions
func (g *GeoDB) RegionSize() (sz Meter) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	sz = g.s.regionSize
	return
}

// RegionsLen will return the number of regions
func (g *GeoDB) RegionsLen() (n int, err error) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return g.s.RegionsLen()
}

// EntriesLen will return the number of targets
func (g *GeoDB) EntriesLen() (n int, err error) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	return g.s.EntriesLen()
}

// Transaction will create a fresh database transaction to insert and populate.
// If the called function returns an error, the database will not be updated.
// If the called function returns no error, the database will be updated with
// newly populated Transaction store
func (g *GeoDB) Transaction(fn func(txn Transaction) error) (err error) {
	var s store
	s.regionSize = g.RegionSize()
	if err = fn(&s); err != nil {
		return
	}

	g.mux.Lock()
	defer g.mux.Unlock()
	g.s = s
	return
}
