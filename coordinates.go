package geodb

type Coordinates struct {
	Latitude  Degree `json:"lat"`
	Longitude Degree `json:"lon"`
}

func (c *Coordinates) IsZero() bool {
	switch {
	case c.Latitude != 0:
		return false
	case c.Longitude != 0:
		return false
	default:
		return true
	}
}

func (c *Coordinates) Location() Location {
	return MakeLocation(c.Latitude, c.Longitude)
}
