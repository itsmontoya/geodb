package geodb

import (
	"sync"
)

// New returns a new GeoDB
func New() *GeoDB {
	var db GeoDB
	return &db
}

// GeoDB maanges the GeoDB service
type GeoDB struct {
	mux   sync.RWMutex
	trunk trunk
	count int
}

// Insert will insert a location to the tree
func (g *GeoDB) Insert(key string, shape Shape) {
	g.mux.Lock()
	defer g.mux.Unlock()
	g.trunk.Insert(key, shape)
	g.count++

	if g.trunk.Len() <= 8 {
		return
	}

	g.trunk.nodes = g.trunk.Split()
}

// GetMatches will return the matching location keys for the provided latitude and longitude
func (g *GeoDB) GetMatches(c Coordinates) (matches []string) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	es := g.trunk.AppendMatches(nil, c)
	matches = make([]string, 0, len(es))
	for _, e := range es {
		matches = append(matches, e.key)
	}

	return
}

// EntriesLen will return the number of targets
func (g *GeoDB) Len() (n int) {
	g.mux.RLock()
	defer g.mux.RUnlock()
	n = g.count
	return
}

// Transaction will create a fresh database transaction to insert and populate.
// If the called function returns an error, the database will not be updated.
// If the called function returns no error, the database will be updated with
// newly populated Transaction store
func (g *GeoDB) Transaction(fn func(txn Transaction) error) (err error) {
	var t trunk

	if err = fn(&t); err != nil {
		return
	}

	g.mux.Lock()
	defer g.mux.Unlock()
	g.trunk = t
	return
}
