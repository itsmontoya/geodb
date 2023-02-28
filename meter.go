package geodb

// Meter represents the meter (metric) as a unit of measurement
type Meter float64

// ToMeters returns itself
func (m Meter) ToMeters() Meter {
	return m
}

// ToFeet returns a conversion from Meteor to Foot
func (m Meter) ToFeet() Foot {
	return Foot(m / metersInFoot)
}
