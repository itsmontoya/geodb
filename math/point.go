package math

func MakePoint(x, y float64) (p Point) {
	p.X = x
	p.Y = y
	return
}

// Point is a coordinate point
type Point struct {
	X float64
	Y float64
}
