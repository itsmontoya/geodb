package geodb

type Entry interface {
	Key() string
	IsWithin(Location) bool
	Center() Location
}
