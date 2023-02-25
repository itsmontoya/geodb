package geodb

func makeEntry(key string, shape Shape) (e entry) {
	e.key = key
	e.shape = shape
	return
}

type entry struct {
	key   string
	shape Shape
}
