package geodb

type Shape interface {
	IsWithin(Location) bool
	Center() Location
	Radius() Meter
}
