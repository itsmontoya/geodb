package math

// y = mx + b
func getSlopeOffset(p1, p2 Point) (m, b float64, s slope) {
	switch {
	case p1.Y == p2.Y:
		// Line is horizontal
		b = p1.Y
		s = slopeHorizontal
	case p1.X == p2.X:
		// Line is vertical
		b = p1.X
		s = slopeVertical
	default:
		// Line is at an angle
		m = getSlope(p1, p2)
		b = p1.Y - (p1.X * m)
		s = slopeDefault
	}

	return
}

func getSlope(a, b Point) (m float64) {
	nx := b.X - a.X
	ny := b.Y - a.Y
	if nx == 0 {
		return 0
	}

	if ny == 0 {
		return 0
	}

	return ny / nx
}

func solveForY(x, m, b float64) (y float64) {
	return m*x + b
}

func solveForX(y, m, b float64) (x float64) {
	x = y - b
	x /= m
	return
}
