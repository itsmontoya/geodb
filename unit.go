package geodb

// Unit represents any unit of measurement which can be converted to both meters and feet
type Unit interface {
	ToMeters() Meter
	ToFeet() Foot
}
