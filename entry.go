package geodb

import "fmt"

func makeEntry(key string, shape Shape) (e entry) {
	e.key = key
	e.shape = shape
	return
}

type entry struct {
	key   string
	shape Shape
}

func (e *entry) Rect() Rect {
	return e.shape.Rect()
}

func (e *entry) String() string {
	return fmt.Sprintf("<%s/%v>", e.key, e.shape)
}
