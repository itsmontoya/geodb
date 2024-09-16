package geodb

type Shape interface {
	IsWithin(Coordinates) bool
	Rect() Rect
}
