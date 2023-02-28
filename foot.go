package geodb

// Foot represents the foot (imperial) as a unit of measurement
type Foot float64

// ToMeters returns a conversion from Foot to Meter
func (f Foot) ToMeters() Meter {
	return Meter(f * metersInFoot)
}

// ToFeet returns itself
func (f Foot) ToFeet() Foot {
	return f
}
